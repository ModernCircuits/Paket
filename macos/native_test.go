// SPDX-License-Identifier: BSL-1.0

package macos

import (
	"bytes"
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/stretchr/testify/assert"
)

func TestNative(t *testing.T) {
	native := Native{}
	out := &bytes.Buffer{}
	assert.Implements(t, (*paket.Generator)(nil), &native)
	assert.Equal(t, "macos-pkg", native.Info().Tag)

	assert.NoError(t, native.Build(out))
	assert.NoError(t, native.Run(out))
	assert.Empty(t, out.String())

	_, err := native.Import(out)
	assert.Error(t, err)
	assert.Error(t, native.Export(paket.Project{}, nil))
	assert.Empty(t, out.String())
}

func Test_NativeReadProjectFile(t *testing.T) {

	{
		generator := &Native{}
		runner := paket.NewRunner()
		assert.NotNil(t, runner)

		err := runner.RegisterGenerator(generator)
		assert.NoError(t, err)

		project, err := runner.ReadProjectFile("testdata/mac_only.hcl")
		assert.NoError(t, err)
		assert.Len(t, project.Installers, 1)

		native, ok := project.Installers[0].(*Native)
		assert.True(t, ok)
		assert.NotNil(t, native)
		assert.NotNil(t, native.installerScript)
		assert.Equal(t, "Plugin Template", native.installerScript.Title)
		assert.NotEmpty(t, native.installerScript.License)
		assert.NotEmpty(t, native.installerScript.Welcome)
		assert.NotEmpty(t, native.installerScript.Conclusion)
		assert.Len(t, native.installerScript.Choices, 3)
		assert.Len(t, native.tasks, 3)
	}

}

func TestNativeExport(t *testing.T) {
	{
		native := Native{}
		out := &bytes.Buffer{}
		err := native.Export(paket.Project{}, out)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "in macos.Native.Export no config set")
	}

	{
		project := paket.Project{Name: "Foo Bar"}
		native := Native{}
		script, _, err := native.createMacInstaller(project, InstallerConfig{})
		assert.NoError(t, err)
		assert.NotNil(t, script)

		native.installerConfig = &InstallerConfig{}
		native.installerScript = script

		out := &bytes.Buffer{}
		err = native.Export(project, out)
		assert.NoError(t, err)
		assert.Contains(t, out.String(), `installer-gui-script authoringTool="Paket" authoringToolVersion="0.1.0"`)
		assert.Contains(t, out.String(), `Foo Bar`)
	}
}
