package innosetup

import (
	"errors"
	"fmt"
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
func (c *Compiler) Configure(paket.Project, paket.Installer) error { return nil }
func (c *Compiler) Build(io.Writer) error                          { return nil }
func (c *Compiler) Run(io.Writer) error                            { return nil }

func (c *Compiler) Import(io.Reader) (*paket.Project, error) {
	return nil, fmt.Errorf("unimplemented import for tag: %s", c.Info().Tag)
}

func (c *Compiler) Export(project paket.Project, w io.Writer) error {
	var installerConfig *paket.Installer
	for _, config := range project.Installers {
		if config.OS == c.Info().Tag {
			installerConfig = &config
			break
		}
	}

	if installerConfig == nil {
		return errors.New("innosetup installer config not found")
	}

	iss := NewISS(project)
	return iss.WriteFile(w)
}
