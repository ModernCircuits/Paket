package paket_test

import (
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/stretchr/testify/assert"
)

func Test_ReadProjectHCL(t *testing.T) {
	tests := []struct {
		file string
		err  bool
	}{
		{file: "testdata/noexist.hcl", err: true},
		{file: "testdata/minimal.hcl", err: false},
		{file: "testdata/full.hcl", err: false},
	}

	for _, tc := range tests {
		t.Run(tc.file, func(t *testing.T) {
			runner := paket.NewRunner()
			project, err := runner.ReadProjectHCL(tc.file)
			if tc.err {
				assert.Error(t, err)
				assert.Nil(t, project)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
