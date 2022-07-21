package paket

import "io"

type NullGenerator struct {
}

func (ng NullGenerator) Tag() string                              { return "null" }
func (ng NullGenerator) Configure(Project, InstallerConfig) error { return nil }
func (ng NullGenerator) Build(io.Writer) error                    { return nil }
func (ng NullGenerator) Run(io.Writer) error                      { return nil }
