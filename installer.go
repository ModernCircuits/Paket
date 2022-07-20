package paket

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/moderncircuits/paket/pkgbuild"
	"github.com/moderncircuits/paket/productbuild"
)

type Project struct {
	Name        string            `hcl:"name"`
	Version     string            `hcl:"version"`
	Identifier  string            `hcl:"identifier"`
	WindowsUUID string            `hcl:"windows_uuid,optional"`
	Installer   []InstallerConfig `hcl:"installer,block"`
}

type InstallerConfig struct {
	OS         string            `hcl:"os,label"`
	Components []ComponentConfig `hcl:"component,block"`
}

type ComponentConfig struct {
	Tag         string `hcl:"tag,label"`
	Name        string `hcl:"name,optional"`
	Version     string `hcl:"version"`
	PayloadPath string `hcl:"payload_path"`
	InstallPath string `hcl:"install_path"`
}

func (p Project) Create() error {

	installerScript := productbuild.InstallerGuiScript{
		AuthoringTool:        "Paket",
		AuthoringToolVersion: "0.1.0",
		AuthoringToolBuild:   "git",

		Title:      p.Name,
		License:    &productbuild.License{File: "LICENSE.txt"},
		Welcome:    nil,
		Conclusion: nil,

		Options: &productbuild.Options{
			Customize: "always",
		},

		ChoicesOutline: productbuild.ChoicesOutline{
			Lines: []productbuild.Line{},
		},

		Choices: []productbuild.Choice{},
	}

	var macOS *InstallerConfig
	for _, installer := range p.Installer {
		if installer.OS == "macOS" {
			macOS = &installer
			break
		}
	}

	for _, component := range macOS.Components {
		id := fmt.Sprintf("%s.%s", p.Identifier, strings.ToLower(component.Tag))
		pkgBuild := pkgbuild.Command{
			Identifier:      id,
			Version:         p.Version,
			Component:       component.PayloadPath,
			InstallLocation: component.InstallPath,
			Output:          fmt.Sprintf("%s.pkg", component.Tag),
		}
		err := pkgBuild.Run()
		if err != nil {
			return err
		}

		line := productbuild.Line{Choice: id}
		installerScript.ChoicesOutline.Lines = append(installerScript.ChoicesOutline.Lines, line)

		ref := productbuild.PkgRef{ID: id, Version: p.Version}
		choice := productbuild.Choice{
			ID:      id,
			Title:   component.Tag,
			PkgRefs: []productbuild.PkgRef{ref},
		}
		installerScript.Choices = append(installerScript.Choices, choice)

		installerScript.PkgRefs = append(installerScript.PkgRefs, ref)

	}

	file, err := xml.MarshalIndent(installerScript, "  ", "    ")
	if err != nil {
		return err
	}
	file = []byte(xml.Header + string(file))

	err = ioutil.WriteFile("notes1.xml", file, 0644)
	if err != nil {
		return err
	}

	return nil

}
