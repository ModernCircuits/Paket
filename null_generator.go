package paket

import "io"

type NullGenerator struct {
}

func (ng NullGenerator) Info() GeneratorInfo                { return GeneratorInfo{"null", []string{}} }
func (ng NullGenerator) Configure(Project, Installer) error { return nil }
func (ng NullGenerator) Build(io.Writer) error              { return nil }
func (ng NullGenerator) Run(io.Writer) error                { return nil }
func (ng NullGenerator) Import(io.Reader) (*Project, error) { return nil, nil }
func (ng NullGenerator) Export(Project, io.Writer) error    { return nil }
func (ng NullGenerator) Parse(Project, InstallerHCL) error  { return nil }
