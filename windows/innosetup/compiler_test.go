package innosetup_test

import (
	"bytes"
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/windows/innosetup"
	"github.com/stretchr/testify/assert"
)

func TestCompiler(t *testing.T) {
	inno := innosetup.Compiler{}
	out := &bytes.Buffer{}
	assert.Implements(t, (*paket.Generator)(nil), &inno)
	assert.Equal(t, "InnoSetup", inno.Info().Tag)
	assert.NoError(t, inno.Configure(paket.ProjectConfig{}, paket.InstallerConfig{}))
	assert.NoError(t, inno.Build(out))
	assert.NoError(t, inno.Run(out))
	assert.Empty(t, out.String())
}
