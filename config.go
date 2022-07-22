package paket

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
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
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("in ReadProjectFile failed to load configuration file: %v", err)
	}

	parser := hclparse.NewParser()
	hclFile, parseDiag := parser.ParseHCL(buf, path)
	if parseDiag.HasErrors() {
		return nil, fmt.Errorf("in ReadProjectFile failed to parse configuration: %s", parseDiag.Error())
	}

	var project Project
	decodeDiag := gohcl.DecodeBody(hclFile.Body, nil, &project)
	if decodeDiag.HasErrors() {
		return nil, fmt.Errorf("in ReadProjectFile failed to decode configuration: %s", decodeDiag.Error())
	}
	return &project, nil
}

type Installer struct {
	OS         string     `hcl:"os,label" json:"os"`
	UUID       string     `hcl:"uuid,optional" json:"uuid,omitempty"`
	Welcome    string     `hcl:"welcome,optional" json:"welcome,omitempty"`
	Conclusion string     `hcl:"conclusion,optional" json:"conclusion,omitempty"`
	Artifacts  []Artifact `hcl:"artifact,block" json:"artifacts"`
}

type Artifact struct {
	Tag         string `hcl:"tag,label" json:"tag"`
	Name        string `hcl:"name,optional" json:"name,omitempty"`
	Version     string `hcl:"version,optional" json:"version,omitempty"`
	Payload     string `hcl:"payload" json:"payload"`
	Destination string `hcl:"destination" json:"destination"`
}
