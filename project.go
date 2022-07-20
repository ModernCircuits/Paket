package paket

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/zclconf/go-cty/cty"
)

type Project struct {
	Name       string            `hcl:"name"`
	Identifier string            `hcl:"identifier"`
	Version    string            `hcl:"version"`
	License    string            `hcl:"license,optional"`
	Installer  []InstallerConfig `hcl:"installer,block"`
}

func ReadFile(path string) (*Project, error) {
	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"windows": cty.ObjectVal(map[string]cty.Value{
				"commoncf64": cty.StringVal("{commoncf64}"),
			}),
		},
	}

	var project Project
	if err := hclsimple.DecodeFile(path, ctx, &project); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}
	return &project, nil
}

func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
