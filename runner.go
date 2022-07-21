package paket

import (
	"fmt"
)

type Runner struct {
	generators []Generator
}

func NewRunner() *Runner {
	return &Runner{
		generators: []Generator{},
	}
}

func (r Runner) RunTag(config ProjectConfig, tag string) error {
	generator := getGeneratorForTag(r.generators, tag)
	if generator == nil {
		return fmt.Errorf("no generator for tag %s", tag)
	}

	return generator.Configure(config, config.Installers[0])
}

func (r *Runner) RegisterGenerator(g Generator) error {
	if hasGeneratorForTag(r.generators, g.Info().Tag) {
		return fmt.Errorf("generator for tag %s already registered", g.Info().Tag)
	}
	r.generators = append(r.generators, g)
	return nil
}

func getGeneratorForTag(generators []Generator, tag string) Generator {
	for _, g := range generators {
		if g.Info().Tag == tag {
			return g
		}
	}
	return nil
}

func hasGeneratorForTag(generators []Generator, tag string) bool {
	return getGeneratorForTag(generators, tag) != nil
}
