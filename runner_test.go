package paket_test

import (
	"testing"

	"github.com/moderncircuits/paket"
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

	t.Run("run registerd", func(t *testing.T) {
		runner := paket.NewRunner()
		assert.NotNil(t, runner)

		config, err := paket.ReadProjectFile("testdata/minimal.hcl")
		assert.NoError(t, err)

		err = runner.RunTag(*config, "null") // unimplemented
		assert.Error(t, err)

		err = runner.RegisterGenerator(paket.NullGenerator{})
		assert.NoError(t, err)

		err = runner.RunTag(*config, "null")
		assert.NoError(t, err)
	})
}
