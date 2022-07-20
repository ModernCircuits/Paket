package paket

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Project struct {
	Name       string            `hcl:"name" json:"name"`
	Vendor     string            `hcl:"vendor" json:"vendor"`
	Identifier string            `hcl:"identifier" json:"identifier"`
	Version    string            `hcl:"version" json:"version"`
	License    string            `hcl:"license,optional" json:"license,omitempty"`
	WorkDir    string            `hcl:"work_dir,optional" json:"work_dir,omitempty"`
	Installer  []InstallerConfig `hcl:"installer,block" json:"installers,omitempty"`
}

func ReadFile(path string, ctx *hcl.EvalContext) (*Project, error) {
	var project Project
	if err := hclsimple.DecodeFile(path, ctx, &project); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}
	return &project, nil
}

func (p Project) Run() error {
	if err := runMacOS(p); err != nil {
		return err
	}

	if err := runWindowsInnoSetup(p); err != nil {
		return err
	}

	return nil
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
