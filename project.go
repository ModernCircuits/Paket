package paket

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Project struct {
	Name       string            `hcl:"name" json:"name"`
	Vendor     string            `hcl:"vendor" json:"vendor"`
	Identifier string            `hcl:"identifier" json:"identifier"`
	Version    string            `hcl:"version" json:"version"`
	License    string            `hcl:"license,optional" json:"license,omitempty"`
	WorkDir    string            `hcl:"work_dir,optional" json:"work_dir,omitempty"`
	Installers []InstallerConfig `hcl:"installer,block" json:"installers,omitempty"`

	generators []Generator
}

func NewProject(path string) (*Project, error) {
	return ReadProjectFile(path, nil)
}

func ReadProjectFile(path string, ctx *hcl.EvalContext) (*Project, error) {
	var project Project
	if err := hclsimple.DecodeFile(path, ctx, &project); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}
	return &project, nil
}

func (p Project) RunTag(tag string) error {
	generator := getGeneratorForTag(p.generators, tag)
	if generator == nil {
		return fmt.Errorf("no generator for tag %s", tag)
	}

	return generator.Configure(p, p.Installers[0])
}

func (p *Project) RegisterGenerator(g Generator) error {
	if hasGeneratorForTag(p.generators, g.Tag()) {
		return fmt.Errorf("generator for tag %s already registered", g.Tag())
	}
	p.generators = append(p.generators, g)
	return nil
}

func getGeneratorForTag(generators []Generator, tag string) Generator {
	for _, g := range generators {
		if g.Tag() == tag {
			return g
		}
	}
	return nil
}

func hasGeneratorForTag(generators []Generator, tag string) bool {
	return getGeneratorForTag(generators, tag) != nil
}
