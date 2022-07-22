package paket_test

import (
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/macos"
	"github.com/moderncircuits/paket/windows/innosetup"
	"github.com/stretchr/testify/assert"
)

func TestRegisterGenerator(t *testing.T) {
	t.Run("register duplicate", func(t *testing.T) {
		runner := paket.NewRunner()
		err := runner.RegisterGenerator(paket.NullGenerator{})
		assert.NoError(t, err)
		err = runner.RegisterGenerator(paket.NullGenerator{})
		assert.Error(t, err)
	})
}

func TestReadProjectFile(t *testing.T) {
	runner := paket.NewRunner()
	assert.NotNil(t, runner)
	generators := []paket.Generator{&macos.Native{}, &innosetup.Compiler{}}
	err := runner.RegisterGenerators(generators)
	assert.NoError(t, err)

	t.Run("path/does/no/exist/config.hcl", func(t *testing.T) {
		_, err := runner.ReadProjectFile("path/does/no/exist/config.hcl")
		assert.Error(t, err)
	})

	t.Run("testdata/minimal.hcl", func(t *testing.T) {
		project, err := runner.ReadProjectFile("testdata/minimal.hcl")
		assert.NoError(t, err)
		assert.Equal(t, "Plugin Template", project.Name)
		assert.Equal(t, "Modern Circuits", project.Vendor)
		assert.Equal(t, "0.1.0", project.Version)
		assert.Equal(t, "com.modern-circuits.plugin-template", project.Identifier)
		assert.Empty(t, project.License)
		assert.Empty(t, project.WorkDir)
		assert.Len(t, project.Installers, 1)
	})

	t.Run("testdata/full.hcl", func(t *testing.T) {
		project, err := runner.ReadProjectFile("testdata/full.hcl")
		assert.NoError(t, err)
		assert.Equal(t, "Plugin Template", project.Name)
		assert.Equal(t, "Modern Circuits", project.Vendor)
		assert.Equal(t, "0.1.0", project.Version)
		assert.Equal(t, "com.modern-circuits.plugin-template", project.Identifier)
		assert.Equal(t, "LICENSE.txt", project.License)
		assert.Len(t, project.Installers, 2)
	})

}
