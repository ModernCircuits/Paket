package paket

import "io"

type Generator interface {
	Tag() string
	Configure(Project, InstallerConfig) error
	Build(io.Writer) error
	Run(io.Writer) error
}
