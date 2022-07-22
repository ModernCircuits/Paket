// SPDX-License-Identifier: BSL-1.0

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
	assert.Implements(t, (*paket.Generator)(nil), null)
	assert.Equal(t, "null", null.Info().Tag)

	assert.NoError(t, null.Configure(paket.Project{}, nil, nil))
	assert.NoError(t, null.Build(out))
	assert.NoError(t, null.Run(out))
	assert.Empty(t, out.String())

	_, err := null.Import(out)
	assert.NoError(t, err)
	assert.NoError(t, null.Export(paket.Project{}, nil))
}
