package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/moderncircuits/paket/pkgbuild"
	"github.com/moderncircuits/paket/productbuild"
)

type Component struct {
	Tag         string `json:"tag,omitempty"`
	PayloadPath string `json:"payloadPath,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
}

type MacOS struct {
	Components []Component `json:"components,omitempty"`
}

type Windows struct {
	Components []Component `json:"components,omitempty"`
}

type Installer struct {
	ProductName    string  `json:"productName,omitempty"`
	ProductVersion string  `json:"productVersion,omitempty"`
	Identifier     string  `json:"identifier,omitempty"`
	ResourcePath   string  `json:"resourcePath,omitempty"`
	MacOS          MacOS   `json:"macOS,omitempty"`
	Windows        Windows `json:"windows,omitempty"`
}

func (i Installer) Create() error {

	installerScript := productbuild.InstallerGuiScript{
		AuthoringTool:        "Paket",
		AuthoringToolVersion: "0.1.0",
		AuthoringToolBuild:   "git",

		Title:      i.ProductName,
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

	for _, component := range i.MacOS.Components {
		id := fmt.Sprintf("%s.%s", i.Identifier, strings.ToLower(component.Tag))
		pkgBuild := pkgbuild.Command{
			Identifier:      id,
			Version:         i.ProductVersion,
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

		ref := productbuild.PkgRef{ID: id, Version: i.ProductVersion}
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
