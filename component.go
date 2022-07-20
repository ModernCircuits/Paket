package paket

type Component struct {
	Tag         string `hcl:"tag,label" json:"tag"`
	Name        string `hcl:"name,optional" json:"name,omitempty"`
	Version     string `hcl:"version,optional" json:"version,omitempty"`
	Payload     string `hcl:"payload" json:"payload"`
	Destination string `hcl:"destination" json:"destination"`
}
