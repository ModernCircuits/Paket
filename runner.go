package paket

import (
	"fmt"
)

type Runner struct {
	WorkDir string

	generators map[string]Generator
}

func NewRunner() *Runner {
	return &Runner{
		generators: map[string]Generator{},
	}
}

func (r Runner) RunTag(config Project, tag string) error {
	generator, found := r.generators[tag]
	if !found {
		return fmt.Errorf("no generator for tag %s", tag)
	}

	return generator.Configure(config, config.Installers[0])
}

func (r *Runner) RegisterGenerator(g Generator) error {
	tag := g.Info().Tag
	if _, found := r.generators[tag]; found {
		return fmt.Errorf("generator for tag %s already registered", tag)
	}
	r.generators[tag] = g
	return nil
}
