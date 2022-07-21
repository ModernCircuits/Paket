package macos

import (
	"bytes"
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/stretchr/testify/assert"
)

func Test_CreateMacInstaller(t *testing.T) {

	{
		path := "../testdata/minimal.hcl"

		project, err := paket.ReadProjectConfigFile(path)
		assert.NoError(t, err)

		script, tasks, err := createMacInstaller(*project, project.Installers[0])
		assert.NoError(t, err)
		assert.NotNil(t, script)
		assert.Equal(t, "Plugin Template", script.Title)
		assert.Empty(t, script.License)
		assert.Empty(t, script.Welcome)
		assert.Empty(t, script.Conclusion)
		assert.Len(t, script.Choices, 1)
		assert.Len(t, tasks, 1)
	}

	{
		path := "../testdata/full.hcl"

		project, err := paket.ReadProjectConfigFile(path)
		assert.NoError(t, err)

		script, tasks, err := createMacInstaller(*project, project.Installers[0])
		assert.NoError(t, err)
		assert.NotNil(t, script)
		assert.Equal(t, "Plugin Template", script.Title)
		assert.NotEmpty(t, script.License)
		assert.NotEmpty(t, script.Welcome)
		assert.NotEmpty(t, script.Conclusion)
		assert.Len(t, script.Choices, 3)
		assert.Len(t, tasks, 3)
	}

}

func TestNative(t *testing.T) {
	native := Native{}
	out := &bytes.Buffer{}
	assert.Implements(t, (*paket.Generator)(nil), &native)
	assert.Equal(t, "macOS", native.Info().Tag)
	assert.NoError(t, native.Configure(paket.ProjectConfig{}, paket.InstallerConfig{}))
	assert.NoError(t, native.Build(out))
	assert.NoError(t, native.Run(out))
	assert.Empty(t, out.String())
}
