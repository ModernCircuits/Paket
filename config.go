package paket

type InstallerConfig struct {
	OS         string            `hcl:"os,label" json:"os"`
	UUID       string            `hcl:"uuid,optional" json:"uuid,omitempty"`
	Welcome    string            `hcl:"welcome,optional" json:"welcome,omitempty"`
	Conclusion string            `hcl:"conclusion,optional" json:"conclusion,omitempty"`
	Components []ComponentConfig `hcl:"component,block" json:"components"`
}

type ComponentConfig struct {
	Tag         string `hcl:"tag,label" json:"tag"`
	Name        string `hcl:"name,optional" json:"name,omitempty"`
	Version     string `hcl:"version,optional" json:"version,omitempty"`
	Payload     string `hcl:"payload" json:"payload"`
	Destination string `hcl:"destination" json:"destination"`
}
