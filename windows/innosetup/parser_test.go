package innosetup

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserInternals(t *testing.T) {
	f, err := os.Open("testdata/Example1.iss")
	assert.NoError(t, err)
	defer f.Close()

	lines := readAllLines(f)
	assert.NotEmpty(t, lines)
	assert.Len(t, lines, 23)

	withoutComments := removeAllCommentLines(lines)
	assert.NotEmpty(t, withoutComments)
	assert.Len(t, withoutComments, 16)

	setupStart := findSetupStartIndex(withoutComments)
	assert.Equal(t, 0, setupStart)

	setupEnd := findSetupEndIndex(withoutComments, setupStart)
	assert.Equal(t, 10, setupEnd)

	setup, err := getSetupSectionLines(withoutComments)
	assert.NoError(t, err)
	assert.NotEmpty(t, setup)
	assert.Len(t, setup, 10)
}
