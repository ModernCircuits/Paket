package macos

import (
	"errors"
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

func (n *Native) Configure(project paket.Project, installer paket.Installer) error {
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

func (n *Native) Import(io.Reader) (*paket.Project, error) {
	return nil, fmt.Errorf("unimplemented import for tag: %s", n.Info().Tag)
}

func (n *Native) Export(project paket.Project, w io.Writer) error {
	var installerConfig *paket.Installer
	for _, config := range project.Installers {
		if config.OS == n.Info().Tag {
			installerConfig = &config
			break
		}
	}

	if installerConfig == nil {
		return errors.New("macOS installer config not found")
	}

	script, _, err := n.createMacInstaller(project, *installerConfig)
	if err != nil {
		return err
	}

	return script.WriteFile(w)
}

func (n *Native) createMacInstaller(project paket.Project, installer paket.Installer) (*productbuild.InstallerGuiScript, []Task, error) {
	if installer.OS != n.Info().Tag {
		return nil, nil, fmt.Errorf("tag %q does not match generator tag %q", installer.OS, n.Info().Tag)
	}
	script := productbuild.NewInstallerGuiScript(project.Name)
	tasks := []Task{}

	if project.License != "" {
		script.License = productbuild.License{File: project.License}
	}

	if installer.Welcome != "" {
		script.Welcome = productbuild.Welcome{File: installer.Welcome}
	}

	if installer.Conclusion != "" {
		script.Conclusion = productbuild.Conclusion{File: installer.Conclusion}
	}

	for _, artifact := range installer.Artifacts {
		id := fmt.Sprintf("%s.%s", project.Identifier, strings.ToLower(artifact.Tag))

		version := project.Version
		if artifact.Version != "" {
			version = artifact.Version
		}

		pkgBuild := pkgbuild.Command{
			Identifier:      id,
			Version:         version,
			Component:       artifact.Payload,
			InstallLocation: artifact.Destination,
			Output:          fmt.Sprintf("%s.pkg", artifact.Tag),
		}

		tasks = append(tasks, func() error { return pkgBuild.Run() })

		line := productbuild.Line{Choice: id}
		script.ChoicesOutline.Lines = append(script.ChoicesOutline.Lines, line)

		ref := productbuild.PkgRef{ID: id, Version: project.Version}
		choice := productbuild.Choice{
			ID:      id,
			Title:   artifact.Tag,
			PkgRefs: []productbuild.PkgRef{ref},
		}
		script.Choices = append(script.Choices, choice)

		script.PkgRefs = append(script.PkgRefs, ref)

	}

	return &script, tasks, nil
}
