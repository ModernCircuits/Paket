package paket

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type ProjectHCL struct {
	Name       string `hcl:"name"`
	Vendor     string `hcl:"vendor"`
	Identifier string `hcl:"identifier"`
	Version    string `hcl:"version"`
	License    string `hcl:"license,optional"`

	InstallerHCL []*InstallerHCL `hcl:"installer,block"`
}

type InstallerHCL struct {
	Name      string   `hcl:"name,label"`
	Generator string   `hcl:"generator,label"`
	HCL       hcl.Body `hcl:",remain"`
}

func (r *Runner) ReadProjectHCL(path string) (*ProjectHCL, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error ReadProjectFile reading file: %v", err)
	}

	parser := hclparse.NewParser()
	src, parseDiag := parser.ParseHCL(buf, path)
	if parseDiag.HasErrors() {
		return nil, fmt.Errorf("error ReadProjectFile parsing HCL: %s", parseDiag.Error())
	}

	ctx := createParseContext()

	var project ProjectHCL
	decodeDiag := gohcl.DecodeBody(src.Body, ctx, &project)
	if decodeDiag.HasErrors() {
		return nil, fmt.Errorf("error ReadProjectFile decoding HCL: %s", decodeDiag.Error())
	}

	for _, installer := range project.InstallerHCL {
		switch installer.Generator {
		case "macos-pkg":
		case "innosetup":
		}
	}

	return &project, nil
}
