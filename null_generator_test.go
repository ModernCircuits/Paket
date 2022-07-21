package paket_test

import (
	"bytes"
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/stretchr/testify/assert"
)

func TestNullGenerator(t *testing.T) {
	null := paket.NullGenerator{}
	out := &bytes.Buffer{}
	assert.Equal(t, "null", null.Tag())
	assert.NoError(t, null.Configure(paket.Project{}, paket.InstallerConfig{}))
	assert.NoError(t, null.Build(out))
	assert.NoError(t, null.Run(out))
}
