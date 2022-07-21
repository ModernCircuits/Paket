package innosetup

import (
	"fmt"
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

	setup, err := getSetupSectionLines(withoutComments)
	assert.NoError(t, err)
	assert.NotEmpty(t, setup)
	assert.Len(t, setup, 9)
}

func Test_parseSetupSection(t *testing.T) {
	tests := [][]string{
		{"com.foo.bar"},
		{"=foobar"},
		{"foobar="},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			setup, err := parseSetupSection(tc)
			assert.Error(t, err)
			assert.Empty(t, setup)
		})
	}
}
