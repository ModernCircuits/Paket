package macos

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/macos/pkgbuild"
	"github.com/moderncircuits/paket/macos/productbuild"
)

type Native struct {
	installerConfig *InstallerConfig
	installerScript *productbuild.InstallerGuiScript
	tasks           []func() error
}

func (n *Native) Info() paket.GeneratorInfo {
	return paket.GeneratorInfo{
		Tag:        "macos-pkg",
		RunnableOn: []string{"darwin"},
	}
}

func (n *Native) ParseInstaller(project paket.Project, body hcl.Body) error {
	var installerConfig InstallerConfig
	diag := gohcl.DecodeBody(body, nil, &installerConfig)
	if diag.HasErrors() {
		return fmt.Errorf("in macos.Native.ParseInstaller failed to decode configuration: %s", diag.Error())
	}

	script, tasks, err := n.createMacInstaller(project, installerConfig)
	if err != nil {
		return err
	}

	n.installerConfig = &installerConfig
	n.installerScript = script
	n.tasks = tasks

	return nil
}

func (n *Native) Build(io.Writer) error { return nil }

func (n *Native) Run(io.Writer) error { return nil }

func (n *Native) Import(io.Reader) (*paket.Project, error) {
	return nil, fmt.Errorf("unimplemented import for tag: %s", n.Info().Tag)
}

func (n *Native) Export(project paket.Project, w io.Writer) error {
	if n.installerScript == nil {
		return errors.New("in macos.Native.Export no config set")
	}

	return n.installerScript.WriteFile(w)
}

func (n *Native) createMacInstaller(project paket.Project, installer InstallerConfig) (*productbuild.InstallerGuiScript, []func() error, error) {
	script := productbuild.NewInstallerGuiScript(project.Name)
	tasks := []func() error{}

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
