package paket

type Component struct {
	Tag         string `hcl:"tag,label"`
	Name        string `hcl:"name,optional"`
	Version     string `hcl:"version,optional"`
	Payload     string `hcl:"payload"`
	Destination string `hcl:"destination"`
}
