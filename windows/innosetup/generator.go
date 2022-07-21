package innosetup

import (
	"io"

	"github.com/moderncircuits/paket"
)

type Generator struct {
}

func (g *Generator) Info() paket.GeneratorInfo {
	return paket.GeneratorInfo{
		Tag:        "InnoSetup",
		RunnableOn: []string{"windows"},
	}
}
func (g *Generator) Configure(paket.ProjectConfig, paket.InstallerConfig) error { return nil }
func (g *Generator) Build(io.Writer) error                                      { return nil }
func (g *Generator) Run(io.Writer) error                                        { return nil }
