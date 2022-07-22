package paket

import "io"

type Exporter interface {
	Export(ProjectConfig, io.Writer) error
}
