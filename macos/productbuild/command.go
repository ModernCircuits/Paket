package productbuild

import (
	"errors"
	"fmt"
	"os/exec"
)

type Command struct {
	Distribution InstallerGuiScript `json:"distribution"`
	ResourcePath string             `json:"resourcePath"`
	PackagePath  string             `json:"packagePath"`
	OutputFile   string             `json:"outputFile"`

	executable string
	args       []string
	cmdOutput  *string
}

func (c *Command) Run() error {
	if c.ResourcePath == "" {
		return errors.New("productbuild: resources is required")
	}
	if c.PackagePath == "" {
		return errors.New("productbuild: package-path is required")
	}
	if c.OutputFile == "" {
		return errors.New("productbuild: output path is required")
	}

	cmd := exec.Command("productbuild")
	if c.executable != "" {
		cmd = exec.Command(c.executable, c.args...)
	}

	distFile := "dist.xml"
	cmd.Args = append(cmd.Args, "--distribution", distFile)
	cmd.Args = append(cmd.Args, "--resources", c.ResourcePath)
	cmd.Args = append(cmd.Args, "--package-path", c.PackagePath)
	cmd.Args = append(cmd.Args, c.OutputFile)

	outBytes, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("productbuild: %v", err)
	}

	c.cmdOutput = new(string)
	*c.cmdOutput = string(outBytes)

	return nil
}

func (c *Command) SetExecutable(cmd string, args []string) {
	c.executable = cmd
	if args != nil {
		c.args = args
	}
}

func (c Command) GetCombinedOutput() (string, error) {
	if c.cmdOutput == nil {
		return "", errors.New("productbuild: output is not set, run command first")
	}

	return *c.cmdOutput, nil
}
