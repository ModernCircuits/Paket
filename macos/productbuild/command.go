package productbuild

type Command struct {
	Distribution InstallerGuiScript `json:"distribution"`
	ResourcePath string             `json:"resourcePath"`
	PackagePath  string             `json:"packagePath"`
	OutputFile   string             `json:"outputFile"`
}

func (c Command) Run() error {
	return nil
}
