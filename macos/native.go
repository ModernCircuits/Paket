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

type Native struct {
	installerScript productbuild.InstallerGuiScript
	tasks           []Task
}

func (n *Native) Info() paket.GeneratorInfo {
	return paket.GeneratorInfo{
		Tag:        "macOS",
		RunnableOn: []string{"darwin"},
	}
}

func (n *Native) Configure(project paket.ProjectConfig, installer paket.InstallerConfig) error {
	script, tasks, err := n.createMacInstaller(project, installer)
	if err != nil {
		return err
	}

	n.installerScript = *script
	n.tasks = tasks
	return nil
}

func (n *Native) Build(io.Writer) error { return nil }

func (n *Native) Run(io.Writer) error { return nil }

func (n *Native) createMacInstaller(project paket.ProjectConfig, installer paket.InstallerConfig) (*productbuild.InstallerGuiScript, []Task, error) {
	if installer.OS != n.Info().Tag {
		return nil, nil, fmt.Errorf("tag %q does not match generator tag %q", installer.OS, n.Info().Tag)
	}
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

		pkgBuild := pkgbuild.Command{
			Identifier:      id,
			Version:         version,
			Component:       component.Payload,
			InstallLocation: component.Destination,
			Output:          fmt.Sprintf("%s.pkg", component.Tag),
		}

		tasks = append(tasks, func() error { return pkgBuild.Run() })

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
