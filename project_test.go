package paket_test

import (
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/stretchr/testify/assert"
)

func TestReadProjectFile(t *testing.T) {
	{
		_, err := paket.ReadProjectFile("path/does/no/exist/config.hcl")
		assert.Error(t, err)
	}

	{
		project, err := paket.ReadProjectFile("testdata/minimal.hcl")
		assert.NoError(t, err)
		assert.Equal(t, "Plugin Template", project.Name)
		assert.Equal(t, "Modern Circuits", project.Vendor)
		assert.Equal(t, "0.1.0", project.Version)
		assert.Equal(t, "com.modern-circuits.plugin-template", project.Identifier)
		assert.Empty(t, project.License)
		assert.Empty(t, project.WorkDir)
		assert.Len(t, project.Installers, 1)
	}

	{
		project, err := paket.ReadProjectFile("testdata/full.hcl")
		assert.NoError(t, err)
		assert.Equal(t, "Plugin Template", project.Name)
		assert.Equal(t, "Modern Circuits", project.Vendor)
		assert.Equal(t, "0.1.0", project.Version)
		assert.Equal(t, "com.modern-circuits.plugin-template", project.Identifier)
		assert.Equal(t, "LICENSE.txt", project.License)
		assert.Len(t, project.Installers, 2)
	}
}
