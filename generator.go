package paket

import "io"

// An installer generator backend. e.g. InnoSetup on windows
// or pkgbuild/productbuild on macOS.
type Generator interface {
	// Tag matches the generator selected in the config file.
	Tag() string

	// Coverts the global project configuration into a more
	// specfic form understood by this generator.
	ConfigureInstaller(ProjectConfig, InstallerConfig) error

	// Creates the build  environment including folders and
	// configuration files needed by the generator.
	BuildInstaller(io.Writer) error

	// Runs the generator. This may be a no-op for some generators.
	RunInstaller(io.Writer) error
}
