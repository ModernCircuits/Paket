package paket

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

type Project struct {
	Name       string      `hcl:"name" json:"name"`
	Vendor     string      `hcl:"vendor" json:"vendor"`
	Identifier string      `hcl:"identifier" json:"identifier"`
	Version    string      `hcl:"version" json:"version"`
	License    string      `hcl:"license,optional" json:"license,omitempty"`
	WorkDir    string      `hcl:"work_dir,optional" json:"work_dir,omitempty"`
	Installers []Installer `hcl:"installer,block" json:"installers,omitempty"`
	generators []Generator
}

func ReadProjectFile(path string) (*Project, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("in ReadProjectFile failed to load configuration file: %v", err)
	}

	parser := hclparse.NewParser()
	hclFile, parseDiag := parser.ParseHCL(buf, path)
	if parseDiag.HasErrors() {
		return nil, fmt.Errorf("in ReadProjectFile failed to parse configuration: %s", parseDiag.Error())
	}

	ctx := createParseContext()

	var project Project
	decodeDiag := gohcl.DecodeBody(hclFile.Body, ctx, &project)
	if decodeDiag.HasErrors() {
		return nil, fmt.Errorf("in ReadProjectFile failed to decode configuration: %s", decodeDiag.Error())
	}
	return &project, nil
}

func createParseContext() *hcl.EvalContext {
	projectName := ""
	if str := os.Getenv("PAKET_VERSION"); str != "" {
		projectName = str
	}
	variables := map[string]cty.Value{
		"env": cty.ObjectVal(map[string]cty.Value{
			"project": cty.StringVal(projectName),
		}),
	}
	return &hcl.EvalContext{
		Variables: variables,
	}
}
