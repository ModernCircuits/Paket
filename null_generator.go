package paket

import "io"

type NullGenerator struct {
}

func (ng NullGenerator) Info() GeneratorInfo                            { return GeneratorInfo{"null", []string{}} }
func (ng NullGenerator) Configure(ProjectConfig, InstallerConfig) error { return nil }
func (ng NullGenerator) Build(io.Writer) error                          { return nil }
func (ng NullGenerator) Run(io.Writer) error                            { return nil }
