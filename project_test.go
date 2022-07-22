package paket_test

import (
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/stretchr/testify/assert"
)

func TestReadProjectFile(t *testing.T) {
	{
		runner := paket.NewRunner()
		assert.NotNil(t, runner)
		_, err := runner.ReadProjectFile("path/does/no/exist/config.hcl")
		assert.Error(t, err)
	}

	// {
	// 	runner := paket.NewRunner()
	// 	assert.NotNil(t, runner)

	// 	err := runner.RegisterGenerator(Nat)
	// 	project, err := runner.ReadProjectFile("testdata/minimal.hcl")
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, "Plugin Template", project.Name)
	// 	assert.Equal(t, "Modern Circuits", project.Vendor)
	// 	assert.Equal(t, "0.1.0", project.Version)
	// 	assert.Equal(t, "com.modern-circuits.plugin-template", project.Identifier)
	// 	assert.Empty(t, project.License)
	// 	assert.Empty(t, project.WorkDir)
	// 	assert.Len(t, project.Installers, 1)
	// }

	// {
	// 	runner := paket.NewRunner()
	// 	assert.NotNil(t, runner)
	// 	project, err := runner.ReadProjectFile("testdata/full.hcl")
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, "Plugin Template", project.Name)
	// 	assert.Equal(t, "Modern Circuits", project.Vendor)
	// 	assert.Equal(t, "0.1.0", project.Version)
	// 	assert.Equal(t, "com.modern-circuits.plugin-template", project.Identifier)
	// 	assert.Equal(t, "LICENSE.txt", project.License)
	// 	assert.Len(t, project.Installers, 2)
	// }
}
