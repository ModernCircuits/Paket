package paket

type Project struct {
	Name       string
	Vendor     string
	Identifier string
	Version    string
	License    string
	WorkDir    string
	Installers []Generator
}
