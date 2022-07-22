// SPDX-License-Identifier: BSL-1.0

package paket

import (
	"io"

	"github.com/hashicorp/hcl/v2"
)

// NullGenerator implements paket.Generator. Does nothing. Fails never.
type NullGenerator struct {
}

// Info implements paket.Generator
func (ng NullGenerator) Info() GeneratorInfo { return GeneratorInfo{"null", []string{}} }

// Configure implements paket.Generator
func (ng NullGenerator) Build(io.Writer) error { return nil }

// Build implements paket.Generator
func (ng NullGenerator) Run(io.Writer) error { return nil }

// Run implements paket.Generator
func (ng NullGenerator) Import(io.Reader) (*Project, error) { return nil, nil }

// Import implements paket.Generator
func (ng NullGenerator) Export(Project, io.Writer) error { return nil }

// Export implements paket.Generator
func (ng NullGenerator) Configure(Project, hcl.EvalContext, hcl.Body) error { return nil }
