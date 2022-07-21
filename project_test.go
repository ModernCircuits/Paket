package paket_test

import (
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	tests := []struct {
		file     string
		hasError bool
	}{
		{file: "noexist/noexist.hcl", hasError: true},
		{file: "testdata/minimal.hcl", hasError: false},
		{file: "testdata/minimal.hcl", hasError: false},
	}

	for _, tc := range tests {
		t.Run(tc.file, func(t *testing.T) {
			_, err := paket.NewProject(tc.file)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestRegisterGenerator(t *testing.T) {
	t.Run("register duplicate", func(t *testing.T) {
		project := paket.Project{}
		err := project.RegisterGenerator(paket.NullGenerator{})
		assert.NoError(t, err)
		err = project.RegisterGenerator(paket.NullGenerator{})
		assert.Error(t, err)
	})

	t.Run("run registerd", func(t *testing.T) {
		project, err := paket.NewProject("testdata/minimal.hcl")
		assert.NoError(t, err)

		err = project.RunTag("null") // unimplemented
		assert.Error(t, err)

		err = project.RegisterGenerator(paket.NullGenerator{})
		assert.NoError(t, err)

		err = project.RunTag("null")
		assert.NoError(t, err)
	})
}
