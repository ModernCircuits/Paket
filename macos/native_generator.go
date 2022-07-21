package macos

import (
	"fmt"
	"io"
	"strings"

	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/macos/pkgbuild"
	"github.com/moderncircuits/paket/macos/productbuild"
)

type Task func() error

type NativeGenerator struct {
	installerScript productbuild.InstallerGuiScript
	tasks           []Task
}

func (g NativeGenerator) Info() paket.GeneratorInfo {
	return paket.GeneratorInfo{
		Tag:        "macOS",
		RunnableOn: []string{"darwin"},
	}
}

func (g *NativeGenerator) ConfigureInstaller(project paket.ProjectConfig, installer paket.InstallerConfig) error {
	script, tasks, err := createMacInstaller(project, installer)
	if err != nil {
		return err
	}

	g.installerScript = *script
	g.tasks = tasks
	return nil
}

func (g NativeGenerator) BuildInstaller(io.Writer) error { return nil }

func (g NativeGenerator) RunInstaller(io.Writer) error { return nil }

func createMacInstaller(project paket.ProjectConfig, installer paket.InstallerConfig) (*productbuild.InstallerGuiScript, []Task, error) {
	script := productbuild.NewInstallerGuiScript(project.Name)
	tasks := []Task{}

	if project.License != "" {
		script.License = &productbuild.License{File: project.License}
	}

	if installer.Welcome != "" {
		script.Welcome = &productbuild.Welcome{File: installer.Welcome}
	}

	if installer.Conclusion != "" {
		script.Conclusion = &productbuild.Conclusion{File: installer.Conclusion}
	}

	for _, component := range installer.Components {
		id := fmt.Sprintf("%s.%s", project.Identifier, strings.ToLower(component.Tag))

		version := project.Version
		if component.Version != "" {
			version = component.Version
		}

		tasks = append(tasks, func() error {
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

			return nil
		})

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

	return &script, tasks, nil
}
