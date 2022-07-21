package paket

type Importer interface {
	Import(string) (ProjectConfig, error)
}
