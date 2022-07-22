package paket

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

type Runner struct {
	WorkDir string

	generators map[string]Generator
}

func (r *Runner) ReadProjectFile(path string) (*Project, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error ReadProjectFile reading file: %v", err)
	}

	parser := hclparse.NewParser()
	src, parseDiag := parser.ParseHCL(buf, path)
	if parseDiag.HasErrors() {
		return nil, fmt.Errorf("error ReadProjectFile parsing HCL: %s", parseDiag.Error())
	}

	var projectHCL projectHCL
	decodeDiag := gohcl.DecodeBody(src.Body, nil, &projectHCL)
	if decodeDiag.HasErrors() {
		return nil, fmt.Errorf("error ReadProjectFile decoding HCL: %s", decodeDiag.Error())
	}

	project := Project{
		Name:       projectHCL.Name,
		Vendor:     projectHCL.Vendor,
		Identifier: projectHCL.Identifier,
		Version:    projectHCL.Version,
		License:    projectHCL.License,
	}

	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"project": cty.ObjectVal(map[string]cty.Value{
				"name":       cty.StringVal(project.Name),
				"vendor":     cty.StringVal(project.Vendor),
				"identifier": cty.StringVal(project.Identifier),
				"version":    cty.StringVal(project.Version),
				"license":    cty.StringVal(project.License),
			}),
		},
	}

	project.Installers = make([]Generator, 0)
	for _, installer := range projectHCL.InstallerHCL {
		g, ok := r.generators[installer.Generator]
		if !ok {
			return nil, fmt.Errorf("no generator registered for tag: %s", installer.Generator)
		}
		installerCtx := ctx.NewChild()
		if err := g.Configure(project, *installerCtx, installer.HCL); err != nil {
			return nil, err
		}

		project.Installers = append(project.Installers, g)
	}

	return &project, nil
}

func NewRunner() *Runner {
	return &Runner{
		generators: map[string]Generator{},
	}
}

func (r *Runner) RegisterGenerator(g Generator) error {
	tag := g.Info().Tag
	if _, found := r.generators[tag]; found {
		return fmt.Errorf("generator for tag %s already registered", tag)
	}
	r.generators[tag] = g
	return nil
}

func (r *Runner) RegisterGenerators(generators []Generator) error {
	for _, generator := range generators {
		if err := r.RegisterGenerator(generator); err != nil {
			return err
		}
	}
	return nil
}

type projectHCL struct {
	Name         string          `hcl:"name"`
	Vendor       string          `hcl:"vendor"`
	Identifier   string          `hcl:"identifier"`
	Version      string          `hcl:"version"`
	License      string          `hcl:"license,optional"`
	InstallerHCL []*installerHCL `hcl:"installer,block"`
}

type installerHCL struct {
	Name      string   `hcl:"name,label"`
	Generator string   `hcl:"generator,label"`
	HCL       hcl.Body `hcl:",remain"`
}
