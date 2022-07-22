// SPDX-License-Identifier: BSL-1.0

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
	tests := []struct {
		file string
		err  bool
	}{
		// Do not exist
		{file: "path/does/no/exist/config.hcl", err: true},
		{file: "local/noexist.hcl", err: true},

		// Syntax/Schema errors
		{file: "testdata/err/invalid_generator.hcl", err: true},
		{file: "testdata/err/no_identifier.hcl", err: true},
		{file: "testdata/err/no_name.hcl", err: true},
		{file: "testdata/err/no_vendor.hcl", err: true},
		{file: "testdata/err/no_version.hcl", err: true},
		{file: "testdata/err/syntax_missing_key.hcl", err: true},
		{file: "testdata/err/syntax_missing_value.hcl", err: true},

		// No errors
		{file: "testdata/full.hcl", err: false},
		{file: "testdata/minimal.hcl", err: false},
	}

	for _, tc := range tests {
		t.Run(tc.file, func(t *testing.T) {
			runner := paket.NewRunner()
			assert.NotNil(t, runner)
			generators := []paket.Generator{&macos.Native{}, &innosetup.Compiler{}}
			err := runner.RegisterGenerators(generators)
			assert.NoError(t, err)

			project, err := runner.ReadProjectFile(tc.file)
			if tc.err {
				assert.Error(t, err)
				assert.Nil(t, project)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, project)
			}

		})
	}

}

func TestReadProjectFileExamples(t *testing.T) {
	runner := paket.NewRunner()
	assert.NotNil(t, runner)
	generators := []paket.Generator{&macos.Native{}, &innosetup.Compiler{}}
	err := runner.RegisterGenerators(generators)
	assert.NoError(t, err)

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
