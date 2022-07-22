package paket

import "io"

type Importer interface {
	Import(io.Reader) (ProjectConfig, error)
}
