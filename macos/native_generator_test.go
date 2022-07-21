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

		project, err := paket.NewProject(path)
		assert.NoError(t, err)

		script, err := createMacInstaller(*project, project.Installers[0])
		assert.NoError(t, err)
		assert.NotNil(t, script)
		assert.Equal(t, "Plugin Template", script.Title)
		assert.Empty(t, script.License)
		assert.Empty(t, script.Welcome)
		assert.Empty(t, script.Conclusion)
		assert.Len(t, script.Choices, 1)
	}

	{
		path := "../testdata/full.hcl"

		project, err := paket.NewProject(path)
		assert.NoError(t, err)

		script, err := createMacInstaller(*project, project.Installers[0])
		assert.NoError(t, err)
		assert.NotNil(t, script)
		assert.Equal(t, "Plugin Template", script.Title)
		assert.NotEmpty(t, script.License)
		assert.NotEmpty(t, script.Welcome)
		assert.NotEmpty(t, script.Conclusion)
		assert.Len(t, script.Choices, 3)
	}

}

func TestNativeGenerator(t *testing.T) {
	null := NativeGenerator{}
	out := &bytes.Buffer{}
	assert.Equal(t, "macOS", null.Tag())
	assert.NoError(t, null.Configure(paket.Project{}, paket.InstallerConfig{}))
	assert.NoError(t, null.Build(out))
	assert.NoError(t, null.Run(out))
	assert.Empty(t, out.String())
}
