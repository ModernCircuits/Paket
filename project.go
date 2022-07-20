package paket

type Project struct {
	Name        string            `hcl:"name"`
	Identifier  string            `hcl:"identifier"`
	Version     string            `hcl:"version"`
	License     string            `hcl:"license,optional"`
	WindowsUUID string            `hcl:"windows_uuid,optional"`
	Installer   []InstallerConfig `hcl:"installer,block"`
}
