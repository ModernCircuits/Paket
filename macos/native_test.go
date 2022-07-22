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
	assert.Equal(t, "macOS", native.Info().Tag)

	assert.Error(t, native.Configure(paket.ProjectConfig{}, paket.InstallerConfig{}))
	assert.NoError(t, native.Build(out))
	assert.NoError(t, native.Run(out))
	assert.Empty(t, out.String())

	_, err := native.Import(out)
	assert.Error(t, err)
	assert.Error(t, native.Export(paket.ProjectConfig{}, nil))
	assert.Empty(t, out.String())
}

func Test_NativeConfigure(t *testing.T) {

	{
		config := paket.ProjectConfig{}
		native := Native{}
		err := native.Configure(config, paket.InstallerConfig{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), `does not match generator tag "macOS"`)
	}

	{
		config, err := paket.ReadProjectConfigFile("../testdata/minimal.hcl")
		assert.NoError(t, err)

		native := Native{}
		err = native.Configure(*config, config.Installers[0])
		assert.NoError(t, err)
		assert.NotNil(t, native.installerScript)
		assert.Equal(t, "Plugin Template", native.installerScript.Title)
		assert.Empty(t, native.installerScript.License)
		assert.Empty(t, native.installerScript.Welcome)
		assert.Empty(t, native.installerScript.Conclusion)
		assert.Len(t, native.installerScript.Choices, 1)
		assert.Len(t, native.tasks, 1)
	}

	{
		config, err := paket.ReadProjectConfigFile("../testdata/full.hcl")
		assert.NoError(t, err)

		native := Native{}
		err = native.Configure(*config, config.Installers[0])
		assert.NoError(t, err)
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
		err := native.Export(paket.ProjectConfig{Installers: []paket.InstallerConfig{}}, out)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "macOS installer config not found")
	}

	{
		project := paket.ProjectConfig{Name: "Foo Bar", Installers: []paket.InstallerConfig{{OS: "macOS"}}}
		native := Native{}
		out := &bytes.Buffer{}
		err := native.Export(project, out)
		assert.NoError(t, err)
		assert.Contains(t, out.String(), `installer-gui-script authoringTool="Paket" authoringToolVersion="0.1.0"`)
		assert.Contains(t, out.String(), `Foo Bar`)
	}
}
