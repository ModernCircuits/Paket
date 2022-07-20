package paket

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/moderncircuits/paket/innosetup"
	"github.com/zclconf/go-cty/cty"
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

func NewProject(path string) (*Project, error) {
	windowsConstants := innosetup.DirectoryConstants()
	windowsVars := map[string]cty.Value{}
	for _, constant := range windowsConstants {
		windowsVars[constant] = cty.StringVal(fmt.Sprintf("{%s}", constant))
	}
	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"windows": cty.ObjectVal(windowsVars),
		},
	}

	project, err := ReadProjectFile(path, ctx)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func ReadProjectFile(path string, ctx *hcl.EvalContext) (*Project, error) {
	var project Project
	if err := hclsimple.DecodeFile(path, ctx, &project); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}
	return &project, nil
}

func (p Project) Run(platform string) error {
	handlers := map[string]func(Project) error{
		"aix":       runNotImplemented,
		"android":   runNotImplemented,
		"darwin":    runMacOS,
		"dragonfly": runNotImplemented,
		"freebsd":   runNotImplemented,
		"illumos":   runNotImplemented,
		"ios":       runNotImplemented,
		"js":        runNotImplemented,
		"linux":     runNotImplemented,
		"netbsd":    runNotImplemented,
		"openbsd":   runNotImplemented,
		"plan9":     runNotImplemented,
		"solaris":   runNotImplemented,
		"windows":   runWindowsInnoSetup,
	}

	handler, ok := handlers[platform]
	if !ok {
		return runNotImplemented(p)
	}

	return handler(p)
}

func runNotImplemented(project Project) error {
	return fmt.Errorf("not implemented for this platform")
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
