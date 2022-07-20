package paket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateMacInstaller(t *testing.T) {

	{
		path := "testdata/minimal.hcl"

		project, err := NewProject(path)
		assert.NoError(t, err)

		script, err := createMacInstaller(*project)
		assert.NoError(t, err)
		assert.NotNil(t, script)
		assert.Equal(t, "Plugin Template", script.Title)
		assert.Empty(t, script.License)
		assert.Empty(t, script.Welcome)
		assert.Empty(t, script.Conclusion)
		assert.Len(t, script.Choices, 1)
	}

	{
		path := "testdata/full.hcl"

		project, err := NewProject(path)
		assert.NoError(t, err)

		script, err := createMacInstaller(*project)
		assert.NoError(t, err)
		assert.NotNil(t, script)
		assert.Equal(t, "Plugin Template", script.Title)
		assert.NotEmpty(t, script.License)
		assert.NotEmpty(t, script.Welcome)
		assert.NotEmpty(t, script.Conclusion)
		assert.Len(t, script.Choices, 3)
	}

}
