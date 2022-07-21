package innosetup

import (
	"io"

	"github.com/moderncircuits/paket"
)

type Generator struct {
}

func (g Generator) Tag() string                                          { return "InnoSetup" }
func (g Generator) Configure(paket.Project, paket.InstallerConfig) error { return nil }
func (g Generator) Build(io.Writer) error                                { return nil }
func (g Generator) Run(io.Writer) error                                  { return nil }
