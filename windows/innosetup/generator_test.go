package innosetup_test

import (
	"bytes"
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/windows/innosetup"
	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	null := innosetup.Generator{}
	out := &bytes.Buffer{}
	assert.Equal(t, "InnoSetup", null.Tag())
	assert.NoError(t, null.Configure(paket.Project{}, paket.InstallerConfig{}))
	assert.NoError(t, null.Build(out))
	assert.NoError(t, null.Run(out))
	assert.Empty(t, out.String())
}
