package paket

import "io"

type GeneratorInfo struct {
	Tag        string
	RunnableOn []string
}

// An installer generator backend. e.g. InnoSetup on windows
// or pkgbuild/productbuild on macOS.
type Generator interface {
	// Info describes the generator.
	Info() GeneratorInfo

	// Coverts the global project configuration into a more
	// specfic form understood by this generator.
	Configure(ProjectConfig, InstallerConfig) error

	// Creates the build  environment including folders and
	// configuration files needed by the generator.
	Build(io.Writer) error

	// Runs the generator. This may be a no-op for some generators.
	Run(io.Writer) error
}
