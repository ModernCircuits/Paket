// SPDX-License-Identifier: BSL-1.0

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

// Info implements paket.Generator
func (c *Compiler) Info() paket.GeneratorInfo {
	return paket.GeneratorInfo{
		Tag:        "innosetup",
		RunnableOn: []string{"windows"},
	}
}

// Configure implements paket.Generator
func (c *Compiler) Configure(project paket.Project, ctx *hcl.EvalContext, body hcl.Body) error {
	var installerConfig InstallerConfig
	diag := gohcl.DecodeBody(body, ctx, &installerConfig)
	if diag.HasErrors() {
		return fmt.Errorf("in innosetup.Compiler.Configure failed to decode configuration: %s", diag.Error())
	}
	c.installerConfig = &installerConfig
	return nil
}

// Build implements paket.Generator
func (c *Compiler) Build(io.Writer) error { return nil }

// Run implements paket.Generator
func (c *Compiler) Run(io.Writer) error { return nil }

// Import implements paket.Generator
func (c *Compiler) Import(io.Reader) (*paket.Project, error) {
	return nil, fmt.Errorf("unimplemented import for generator: %s", c.Info().Tag)
}

// Export implements paket.Generator
func (c *Compiler) Export(project paket.Project, w io.Writer) error {
	if c.installerConfig == nil {
		return errors.New("innosetup installer config not found")
	}

	iss := NewISS(project)
	return iss.WriteFile(w)
}
