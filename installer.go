package paket

type Installer struct {
	Name       string     `hcl:"name,label" json:"name"`
	Generator  string     `hcl:"generator,label" json:"generator"`
	UUID       string     `hcl:"uuid,optional" json:"uuid,omitempty"`
	Welcome    string     `hcl:"welcome,optional" json:"welcome,omitempty"`
	Conclusion string     `hcl:"conclusion,optional" json:"conclusion,omitempty"`
	Artifacts  []Artifact `hcl:"artifact,block" json:"artifacts"`
}
