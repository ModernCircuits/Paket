package paket

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Project struct {
	Name       string      `hcl:"name" json:"name"`
	Vendor     string      `hcl:"vendor" json:"vendor"`
	Identifier string      `hcl:"identifier" json:"identifier"`
	Version    string      `hcl:"version" json:"version"`
	License    string      `hcl:"license,optional" json:"license,omitempty"`
	WorkDir    string      `hcl:"work_dir,optional" json:"work_dir,omitempty"`
	Installers []Installer `hcl:"installer,block" json:"installers,omitempty"`
}

func ReadProjectFile(path string) (*Project, error) {
	var project Project
	if err := hclsimple.DecodeFile(path, nil, &project); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}
	return &project, nil
}

type Installer struct {
	OS         string     `hcl:"os,label" json:"os"`
	UUID       string     `hcl:"uuid,optional" json:"uuid,omitempty"`
	Welcome    string     `hcl:"welcome,optional" json:"welcome,omitempty"`
	Conclusion string     `hcl:"conclusion,optional" json:"conclusion,omitempty"`
	Artifacts  []artifact `hcl:"artifact,block" json:"artifacts"`
}

type artifact struct {
	Tag         string `hcl:"tag,label" json:"tag"`
	Name        string `hcl:"name,optional" json:"name,omitempty"`
	Version     string `hcl:"version,optional" json:"version,omitempty"`
	Payload     string `hcl:"payload" json:"payload"`
	Destination string `hcl:"destination" json:"destination"`
}
