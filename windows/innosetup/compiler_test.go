// SPDX-License-Identifier: BSL-1.0

package innosetup_test

import (
	"bytes"
	"testing"

	"github.com/hashicorp/hcl/v2/hclparse"
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

func TestCompilerConfigure(t *testing.T) {
	tests := []struct {
		name string
		src  string
		err  bool
	}{
		{name: "empty", src: "", err: false},
		{name: "unkown", src: `uuid="com.example.foo"`, err: false},

		{name: "unkown key", src: `test_var="4"`, err: true},
		{name: "unkown var", src: `uuid="${some_var}"`, err: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parser := hclparse.NewParser()
			src, diag := parser.ParseHCL([]byte(tc.src), tc.name)
			assert.False(t, diag.HasErrors())

			inno := innosetup.Compiler{}
			err := inno.Configure(paket.Project{}, nil, src.Body)
			assert.Equal(t, tc.err, err != nil)
		})
	}

}

func TestCompilerExport(t *testing.T) {
	inno := innosetup.Compiler{}
	out := &bytes.Buffer{}
	err := inno.Export(paket.Project{}, out)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "innosetup installer config not found")
}
