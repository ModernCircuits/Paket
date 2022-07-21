package paket_test

import (
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/stretchr/testify/assert"
)

func TestRegisterGenerator(t *testing.T) {
	project := paket.Project{}
	err := project.RegisterGenerator(paket.NullGenerator{})
	assert.NoError(t, err)
	err = project.RegisterGenerator(paket.NullGenerator{})
	assert.Error(t, err)

	{
		project, err := paket.NewProject("testdata/minimal.hcl")
		assert.NoError(t, err)
		err = project.RunTag("null") // unimplemented
		assert.Error(t, err)

		err = project.RegisterGenerator(paket.NullGenerator{})
		assert.NoError(t, err)

		err = project.RunTag("null")
		assert.NoError(t, err)
	}
}
