package paket

import (
	"os"

	"github.com/hashicorp/hcl/v2"
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
