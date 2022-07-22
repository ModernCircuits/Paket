package paket

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type ProjectHCL struct {
	Name       string `hcl:"name"`
	Vendor     string `hcl:"vendor"`
	Identifier string `hcl:"identifier"`
	Version    string `hcl:"version"`
	License    string `hcl:"license,optional"`

	InstallerHCL []*InstallerHCL `hcl:"installer,block"`
}

type InstallerHCL struct {
	Name      string   `hcl:"name,label"`
	Generator string   `hcl:"generator,label"`
	HCL       hcl.Body `hcl:",remain"`
}

func (r *Runner) ReadProjectHCL(path string) (*Project, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error ReadProjectFile reading file: %v", err)
	}

	parser := hclparse.NewParser()
	src, parseDiag := parser.ParseHCL(buf, path)
	if parseDiag.HasErrors() {
		return nil, fmt.Errorf("error ReadProjectFile parsing HCL: %s", parseDiag.Error())
	}

	ctx := createParseContext()

	var projectHCL ProjectHCL
	decodeDiag := gohcl.DecodeBody(src.Body, ctx, &projectHCL)
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

	project.generators = make([]Generator, 0)
	for _, installer := range projectHCL.InstallerHCL {
		g, ok := r.generators[installer.Generator]
		if !ok {
			return nil, fmt.Errorf("no generator registered for tag: %s", installer.Generator)
		}
		if err := g.Parse(project, installer.HCL); err != nil {
			return nil, err
		}

		project.generators = append(project.generators, g)
	}

	return &project, nil
}
