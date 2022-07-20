package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	ProjectName    string            `hcl:"project_name"`
	ProjectVersion string            `hcl:"project_version"`
	BundleID       string            `hcl:"bundle_id"`
	WindowsUUID    string            `hcl:"windows_uuid,optional"`
	Installer      []InstallerConfig `hcl:"installer,block"`
}

type InstallerConfig struct {
	OS         string            `hcl:"os,label"`
	Components []ComponentConfig `hcl:"component,block"`
}

type ComponentConfig struct {
	Tag         string `hcl:"tag,label"`
	Name        string `hcl:"name,optional"`
	Version     string `hcl:"version"`
	Payload     string `hcl:"payload"`
	InstallPath string `hcl:"install_path"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var config Config
	err := hclsimple.DecodeFile("test_data/config.hcl", nil, &config)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %v", err)
	}

	jsonFile, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("test.json", jsonFile, 0644)
	if err != nil {
		return err
	}

	return nil
}
