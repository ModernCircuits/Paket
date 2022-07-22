package innosetup

import (
	"errors"
	"fmt"
	"io"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/moderncircuits/paket"
)

type Compiler struct {
	installerConfig *InstallerConfig
}

func (c *Compiler) Info() paket.GeneratorInfo {
	return paket.GeneratorInfo{
		Tag:        "innosetup",
		RunnableOn: []string{"windows"},
	}
}

func (c *Compiler) Parse(project paket.Project, body hcl.Body) error {
	var installerConfig InstallerConfig
	diag := gohcl.DecodeBody(body, nil, &installerConfig)
	if diag.HasErrors() {
		return fmt.Errorf("in macos.Native.Parse failed to decode configuration: %s", diag.Error())
	}
	c.installerConfig = &installerConfig
	return nil
}

func (c *Compiler) Configure(paket.Project, paket.Installer) error { return nil }
func (c *Compiler) Build(io.Writer) error                          { return nil }
func (c *Compiler) Run(io.Writer) error                            { return nil }

func (c *Compiler) Import(io.Reader) (*paket.Project, error) {
	return nil, fmt.Errorf("unimplemented import for generator: %s", c.Info().Tag)
}

func (c *Compiler) Export(project paket.Project, w io.Writer) error {
	var installerConfig *paket.Installer
	for _, config := range project.Installers {
		if config.Generator == c.Info().Tag {
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
