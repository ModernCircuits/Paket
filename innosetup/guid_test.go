package innosetup_test

import (
	"regexp"
	"testing"

	"github.com/moderncircuits/paket/innosetup"
	"github.com/stretchr/testify/assert"
)

func TestNewInnoSetupGUID(t *testing.T) {
	uuid, err := innosetup.NewInnoSetupGUID()
	assert.NotEmpty(t, uuid)
	assert.NoError(t, err)

	matched, err := regexp.Match("[0-F]{8}-([0-F]{4}-){3}[0-F]{12}", []byte(uuid))
	assert.NoError(t, err)
	assert.True(t, matched)
}
