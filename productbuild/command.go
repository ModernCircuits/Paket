package productbuild

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/pkgbuild"
)

func CreateDistributionXML(project paket.Project) error {
	installerScript := InstallerGuiScript{
		AuthoringTool:        "Paket",
		AuthoringToolVersion: "0.1.0",
		AuthoringToolBuild:   "git",

		Title:      project.Name,
		License:    &License{File: project.License},
		Welcome:    nil,
		Conclusion: nil,

		Options: &Options{
			Customize: "always",
		},

		ChoicesOutline: ChoicesOutline{
			Lines: []Line{},
		},

		Choices: []Choice{},
	}

	var macOS *paket.InstallerConfig
	for _, installer := range project.Installer {
		if installer.OS == "macOS" {
			macOS = &installer
			break
		}
	}

	for _, component := range macOS.Components {
		id := fmt.Sprintf("%s.%s", project.Identifier, strings.ToLower(component.Tag))

		version := project.Version
		if component.Version != "" {
			version = component.Version
		}

		pkgBuild := pkgbuild.Command{
			Identifier:      id,
			Version:         version,
			Component:       component.Payload,
			InstallLocation: component.Destination,
			Output:          fmt.Sprintf("%s.pkg", component.Tag),
		}
		err := pkgBuild.Run()
		if err != nil {
			return err
		}

		line := Line{Choice: id}
		installerScript.ChoicesOutline.Lines = append(installerScript.ChoicesOutline.Lines, line)

		ref := PkgRef{ID: id, Version: project.Version}
		choice := Choice{
			ID:      id,
			Title:   component.Tag,
			PkgRefs: []PkgRef{ref},
		}
		installerScript.Choices = append(installerScript.Choices, choice)

		installerScript.PkgRefs = append(installerScript.PkgRefs, ref)

	}

	file, err := xml.MarshalIndent(installerScript, "  ", "    ")
	if err != nil {
		return err
	}
	file = []byte(xml.Header + string(file))

	err = ioutil.WriteFile("test.xml", file, 0644)
	if err != nil {
		return err
	}

	return nil
}
