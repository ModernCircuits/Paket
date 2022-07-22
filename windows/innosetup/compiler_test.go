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
	assert.Equal(t, "innosetup", inno.Info().Tag)

	assert.NoError(t, inno.Build(out))
	assert.NoError(t, inno.Run(out))
	assert.Empty(t, out.String())

	_, err := inno.Import(out)
	assert.Error(t, err)
	assert.Error(t, inno.Export(paket.Project{}, nil))
	assert.Empty(t, out.String())
}

func TestCompilerExport(t *testing.T) {
	{
		inno := innosetup.Compiler{}
		out := &bytes.Buffer{}
		err := inno.Export(paket.Project{}, out)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "innosetup installer config not found")
	}

	// {
	// 	project := paket.Project{Name: "Foo Bar"}
	// 	inno := innosetup.Compiler{}
	// 	out := &bytes.Buffer{}
	// 	err := inno.Export(project, out)
	// 	assert.NoError(t, err)
	// 	assert.Contains(t, out.String(), `AppName="Foo Bar"`)
	// }
}
