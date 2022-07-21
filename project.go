package paket

import (
	"fmt"
)

type Project struct {
	config     ProjectConfig
	generators []Generator
}

func NewProject(path string) (*Project, error) {
	config, err := ReadProjectConfigFile(path)
	if err != nil {
		return nil, err
	}
	return &Project{
		config:     *config,
		generators: []Generator{},
	}, nil
}

func (p Project) RunTag(tag string) error {
	generator := getGeneratorForTag(p.generators, tag)
	if generator == nil {
		return fmt.Errorf("no generator for tag %s", tag)
	}

	return generator.ConfigureInstaller(p.config, p.config.Installers[0])
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
