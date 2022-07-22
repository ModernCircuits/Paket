package productbuild_test

import (
	"bytes"
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

func TestInstallerGuiScriptWriteFile(t *testing.T) {
	out := &bytes.Buffer{}
	script := productbuild.NewInstallerGuiScript("Foo Bar")
	err := script.WriteFile(out)
	assert.NoError(t, err)
	assert.Contains(t, out.String(), `installer-gui-script authoringTool="Paket" authoringToolVersion=`)
	assert.Contains(t, out.String(), `Foo Bar`)
}
