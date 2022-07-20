package pkgbuild

import (
	"errors"
	"fmt"
	"os/exec"
)

type Command struct {
	Identifier      string
	Version         string
	Component       string
	InstallLocation string
	Output          string
}

func (c Command) Run() error {
	if c.Identifier == "" {
		return errors.New("pkgbuild: identifier is required")
	}
	if c.Version == "" {
		return errors.New("pkgbuild: version is required")
	}
	if c.Component == "" {
		return errors.New("pkgbuild: component is required")
	}
	if c.InstallLocation == "" {
		return errors.New("pkgbuild: install-location is required")
	}
	if c.Output == "" {
		return errors.New("pkgbuild: output path is required")
	}

	cmd := exec.Command("pkgbuild")
	cmd.Args = append(cmd.Args, "--identifier", c.Identifier)
	cmd.Args = append(cmd.Args, "--version", c.Version)
	cmd.Args = append(cmd.Args, "--component", c.Component)
	cmd.Args = append(cmd.Args, "--install-location", c.InstallLocation)
	cmd.Args = append(cmd.Args, c.Output)

	fmt.Printf("%v\n", cmd)

	// if err := cmd.Wait(); err != nil {
	// 	return fmt.Errorf("pkgbuild: %v", err)
	// }

	return nil
}
