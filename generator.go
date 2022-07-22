package paket

import (
	"io"

	"github.com/hashicorp/hcl/v2"
)

type GeneratorInfo struct {
	Tag        string
	RunnableOn []string
}

// An installer generator backend. e.g. InnoSetup on windows
// or pkgbuild/productbuild on macOS.
type Generator interface {
	// Info describes the generator.
	Info() GeneratorInfo

	// Converts the hcl installer block configuration into a more specfic form
	// understood by this generator.
	Configure(Project, hcl.EvalContext, hcl.Body) error

	// Creates the build  environment including folders and configuration files
	// needed by the generator.
	Build(io.Writer) error

	// Runs the generator. This may be a no-op for some generators.
	Run(io.Writer) error

	// Import a platform specific configuration from a reader. The reader will
	// probably come from a configuration file like an InnoSetup *.iss file or
	// a pkgbuild/productbuild *.xml distribution file. Roundtrip import/export
	// is most likely lossy.
	Import(io.Reader) (*Project, error)

	// Export a platform specific configuration to a writer. Roundtrip
	// import/export is most likely lossy.
	Export(Project, io.Writer) error
}
