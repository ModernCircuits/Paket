package productbuild

// productbuild                                        \
//     --distribution "distribution-fixed.xml"         \
//     --resources .                                   \
//     --package-path "$out_dir"                       \
//     "$out_dir/$plugin_name.pkg"                     \

type Command struct {
	Distribution InstallerGuiScript `json:"distribution"`
	ResourcePath string             `json:"resourcePath"`
	PackagePath  string             `json:"packagePath"`
	OutputFile   string             `json:"outputFile"`
}

func (c Command) Run() error {
	return nil
}
