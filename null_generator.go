package paket

import "io"

type NullGenerator struct {
}

func (ng NullGenerator) Tag() string                                             { return "null" }
func (ng NullGenerator) ConfigureInstaller(ProjectConfig, InstallerConfig) error { return nil }
func (ng NullGenerator) BuildInstaller(io.Writer) error                          { return nil }
func (ng NullGenerator) RunInstaller(io.Writer) error                            { return nil }
