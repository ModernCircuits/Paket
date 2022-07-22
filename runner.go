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

func (r *Runner) ReadProjectFile(path string) (*Project, error) {
	return r.ReadProjectHCL(path)
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
