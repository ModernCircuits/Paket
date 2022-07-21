package macos

import (
	"fmt"
	"io"
	"strings"

	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/macos/productbuild"
)

type NativeGenerator struct {
}

func (ng NativeGenerator) Tag() string                                          { return "macOS" }
func (ng NativeGenerator) Configure(paket.Project, paket.InstallerConfig) error { return nil }
func (ng NativeGenerator) Build(io.Writer) error                                { return nil }
func (ng NativeGenerator) Run(io.Writer) error                                  { return nil }

func createMacInstaller(project paket.Project) (*productbuild.InstallerGuiScript, error) {
	script := productbuild.NewInstallerGuiScript(project.Name)

	if project.License != "" {
		script.License = &productbuild.License{File: project.License}
	}

	var macOS *paket.InstallerConfig
	for _, installer := range project.Installer {
		if installer.OS == "macOS" {
			macOS = &installer
			break
		}
	}

	if macOS == nil {
		return nil, fmt.Errorf("no macOS installer config found")
	}

	if macOS.Welcome != "" {
		script.Welcome = &productbuild.Welcome{File: macOS.Welcome}
	}

	if macOS.Conclusion != "" {
		script.Conclusion = &productbuild.Conclusion{File: macOS.Conclusion}
	}

	for _, component := range macOS.Components {
		id := fmt.Sprintf("%s.%s", project.Identifier, strings.ToLower(component.Tag))

		// version := project.Version
		// if component.Version != "" {
		// 	version = component.Version
		// }

		// pkgBuild := pkgbuild.Command{
		// 	Identifier:      id,
		// 	Version:         version,
		// 	Component:       component.Payload,
		// 	InstallLocation: component.Destination,
		// 	Output:          fmt.Sprintf("%s.pkg", component.Tag),
		// }
		// err := pkgBuild.Run()
		// if err != nil {
		// 	return nil, err
		// }

		line := productbuild.Line{Choice: id}
		script.ChoicesOutline.Lines = append(script.ChoicesOutline.Lines, line)

		ref := productbuild.PkgRef{ID: id, Version: project.Version}
		choice := productbuild.Choice{
			ID:      id,
			Title:   component.Tag,
			PkgRefs: []productbuild.PkgRef{ref},
		}
		script.Choices = append(script.Choices, choice)

		script.PkgRefs = append(script.PkgRefs, ref)

	}

	return &script, nil
}
