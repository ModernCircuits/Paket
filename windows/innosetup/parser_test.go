// SPDX-License-Identifier: BSL-1.0

package innosetup

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserInternals(t *testing.T) {
	f, err := os.Open("testdata/Example1.iss")
	assert.NoError(t, err)
	defer f.Close()

	lines := readAllLines(f)
	assert.NotNil(t, lines)
	assert.Len(t, lines, 23)

	withoutComments := removeAllCommentLines(lines)
	assert.NotNil(t, withoutComments)
	assert.Len(t, withoutComments, 16)

	setup, err := getSetupLines(withoutComments)
	assert.NoError(t, err)
	assert.NotNil(t, setup)
	assert.Len(t, setup, 9)
}

func Test_parseSetup(t *testing.T) {
	tests := []struct {
		line string
		err  bool
	}{
		{line: "AppName=My Program", err: false},
		{line: "NoExist=My Program", err: true},
		{line: "com.foo.bar", err: true},
		{line: "=foobar", err: true},
		{line: "foobar=", err: true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s err: %v", tc.line, tc.err), func(t *testing.T) {
			setup, err := parseSetup([]string{tc.line})
			if tc.err {
				assert.Error(t, err)
				assert.Nil(t, setup)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_setSetupStructField(t *testing.T) {
	t.Run("unimplemnted type", func(t *testing.T) {
		someType := &struct {
			AppName    string
			AppVersion int
		}{}
		r := reflect.ValueOf(someType).Elem()
		assert.NoError(t, setSetupStructField(r, "AppName", "Foo"))

		err := setSetupStructField(r, "AppVersion", "1.0")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unimplemented type int")
	})

}
