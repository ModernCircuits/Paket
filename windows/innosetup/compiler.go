package innosetup

import (
	"io"

	"github.com/moderncircuits/paket"
)

type Compiler struct {
}

func (c *Compiler) Info() paket.GeneratorInfo {
	return paket.GeneratorInfo{
		Tag:        "InnoSetup",
		RunnableOn: []string{"windows"},
	}
}
func (c *Compiler) Configure(paket.ProjectConfig, paket.InstallerConfig) error { return nil }
func (c *Compiler) Build(io.Writer) error                                      { return nil }
func (c *Compiler) Run(io.Writer) error                                        { return nil }
