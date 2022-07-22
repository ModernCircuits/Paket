// SPDX-License-Identifier: BSL-1.0

package paket_test

import (
	"runtime"
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/macos"
	"github.com/moderncircuits/paket/windows/innosetup"
	"github.com/stretchr/testify/assert"
)

func fileDoesNotExistMessage() string {
	if runtime.GOOS == "windows" {
		return "The system cannot find the path specified."
	}
	return "no such file or directory"
}

func TestRegisterGenerator(t *testing.T) {
	t.Run("register duplicate", func(t *testing.T) {
		runner := paket.NewRunner()
		err := runner.RegisterGenerator(paket.NullGenerator{})
		assert.NoError(t, err)
		err = runner.RegisterGenerator(paket.NullGenerator{})
		assert.Error(t, err)
	})
}

func TestRegisterGenerators(t *testing.T) {
	t.Run("register duplicate", func(t *testing.T) {
		runner := paket.NewRunner()
		err := runner.RegisterGenerators([]paket.Generator{paket.NullGenerator{}, paket.NullGenerator{}})
		assert.Error(t, err)
	})
}

func TestReadProjectFile(t *testing.T) {
	tests := []struct {
		file string
		err  string
	}{
		// Do not exist
		{file: "/path/does/no/exist/config.hcl", err: fileDoesNotExistMessage()},
		{file: "local/noexist.hcl", err: fileDoesNotExistMessage()},

		// // Syntax/Schema errors
		{file: "testdata/err/invalid_generator.hcl", err: "no generator registered for tag: macos-magic"},
		{file: "testdata/err/invalid_installer_block.hcl", err: `An argument named "unknown"`},
		{file: "testdata/err/missing_identifier.hcl", err: `The argument "identifier"`},
		{file: "testdata/err/missing_name.hcl", err: `The argument "name"`},
		{file: "testdata/err/missing_vendor.hcl", err: `The argument "vendor"`},
		{file: "testdata/err/missing_version.hcl", err: `The argument "version"`},
		{file: "testdata/err/syntax_missing_key.hcl", err: "Argument or block definition required"},
		{file: "testdata/err/syntax_missing_value.hcl", err: "Invalid expression"},

		// // No errors
		{file: "testdata/full.hcl", err: ""},
		{file: "testdata/minimal.hcl", err: ""},
	}

	for _, tc := range tests {
		t.Run(tc.file, func(t *testing.T) {
			runner := paket.NewRunner()
			assert.NotNil(t, runner)
			generators := []paket.Generator{&macos.Native{}, &innosetup.Compiler{}}
			err := runner.RegisterGenerators(generators)
			assert.NoError(t, err)

			project, err := runner.ReadProjectFile(tc.file)
			if tc.err != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.err)
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
