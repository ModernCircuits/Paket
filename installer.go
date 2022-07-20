package paket

type InstallerConfig struct {
	OS         string      `hcl:"os,label" json:"os"`
	UUID       string      `hcl:"uuid,optional" json:"uuid,omitempty"`
	Welcome    string      `hcl:"welcome,optional" json:"welcome,omitempty"`
	Conclusion string      `hcl:"conclusion,optional" json:"conclusion,omitempty"`
	Components []Component `hcl:"component,block" json:"components"`
}
