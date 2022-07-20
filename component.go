package paket

type Component struct {
	Tag         string `hcl:"tag,label"`
	Name        string `hcl:"name,optional"`
	Version     string `hcl:"version,optional"`
	PayloadPath string `hcl:"payload_path"`
	InstallPath string `hcl:"install_path"`
}
