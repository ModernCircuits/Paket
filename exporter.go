package paket

type Exporter interface {
	Export(ProjectConfig, string) error
}
