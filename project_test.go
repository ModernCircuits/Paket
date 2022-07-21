package paket_test

import (
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	{
		_, err := paket.NewProject("path/does/no/exist/config.hcl")
		assert.Error(t, err)
	}

	{
		project, err := paket.NewProject("testdata/minimal.hcl")
		assert.NoError(t, err)
		assert.Equal(t, "Plugin Template", project.Name)
		assert.Equal(t, "Modern Circuits", project.Vendor)
		assert.Equal(t, "0.1.0", project.Version)
		assert.Equal(t, "com.modern-circuits.plugin-template", project.Identifier)
		assert.Empty(t, project.License)
		assert.Empty(t, project.WorkDir)
		assert.Len(t, project.Installer, 1)
	}

	{
		project, err := paket.NewProject("testdata/full.hcl")
		assert.NoError(t, err)
		assert.Equal(t, "Plugin Template", project.Name)
		assert.Equal(t, "Modern Circuits", project.Vendor)
		assert.Equal(t, "0.1.0", project.Version)
		assert.Equal(t, "com.modern-circuits.plugin-template", project.Identifier)
		assert.Equal(t, "LICENSE.txt", project.License)
		assert.Len(t, project.Installer, 2)
	}

	{
		project, err := paket.NewProject("testdata/minimal.hcl")
		assert.NoError(t, err)
		err = project.RunTag("null") // unimplemented
		assert.Error(t, err)

		err = project.RegisterGenerator(paket.NullGenerator{})
		assert.NoError(t, err)

		err = project.RunTag("null") // unimplemented
		assert.NoError(t, err)
	}
}

func TestRegisterGenerator(t *testing.T) {
	project := paket.Project{}
	err := project.RegisterGenerator(paket.NullGenerator{})
	assert.NoError(t, err)
	err = project.RegisterGenerator(paket.NullGenerator{})
	assert.Error(t, err)
}
