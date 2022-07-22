package productbuild_test

import (
	"testing"

	"github.com/moderncircuits/paket/macos/productbuild"
	"github.com/stretchr/testify/assert"
)

func TestNewInstallerGuiScript(t *testing.T) {
	{
		name := "FooBar"
		script := productbuild.NewInstallerGuiScript(name)
		assert.Equal(t, name, script.Title)
		assert.Empty(t, script.License)
		assert.Empty(t, script.Welcome)
		assert.Empty(t, script.Conclusion)
		assert.Empty(t, script.Options)
	}
}
func TestReadInstallerGuiScriptFile(t *testing.T) {
	{
		path := "testdata/noexist.xml"
		script, err := productbuild.ReadInstallerGuiScriptFile(path)
		assert.Error(t, err)
		assert.Nil(t, script)
	}

	{
		path := "testdata/distribution.xml"
		script, err := productbuild.ReadInstallerGuiScriptFile(path)
		assert.NoError(t, err)
		assert.NotNil(t, script)
		assert.Empty(t, script.Title)
		assert.Empty(t, script.License)
		assert.Empty(t, script.Welcome)
		assert.Empty(t, script.Conclusion)
	}
}
