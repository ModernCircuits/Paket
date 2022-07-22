package paket

import (
	"io"

	"github.com/hashicorp/hcl/v2"
)

type NullGenerator struct {
}

func (ng NullGenerator) Info() GeneratorInfo                    { return GeneratorInfo{"null", []string{}} }
func (ng NullGenerator) Build(io.Writer) error                  { return nil }
func (ng NullGenerator) Run(io.Writer) error                    { return nil }
func (ng NullGenerator) Import(io.Reader) (*Project, error)     { return nil, nil }
func (ng NullGenerator) Export(Project, io.Writer) error        { return nil }
func (ng NullGenerator) ParseInstaller(Project, hcl.Body) error { return nil }
