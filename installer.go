package paket

type InstallerConfig struct {
	OS         string      `hcl:"os,label"`
	Components []Component `hcl:"component,block"`
}
