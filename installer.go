package paket

type InstallerConfig struct {
	OS         string      `hcl:"os,label"`
	UUID       string      `hcl:"uuid,optional"`
	Welcome    string      `hcl:"welcome,optional"`
	Conclusion string      `hcl:"conclusion,optional"`
	Components []Component `hcl:"component,block"`
}
