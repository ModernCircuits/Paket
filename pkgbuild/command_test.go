package pkgbuild_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/moderncircuits/paket/pkgbuild"
	"github.com/stretchr/testify/assert"
)

func TestCommandRunMissingArgs(t *testing.T) {
	tests := []pkgbuild.Command{
		{},
		{Identifier: "com.foo.bar"},
		{Identifier: "com.foo.bar", Version: "1.0"},
		{Identifier: "com.foo.bar", Version: "1.0", Component: "foo.app"},
		{Identifier: "com.foo.bar", Version: "1.0", Component: "foo.app", InstallLocation: "/usr/local/foo"},
		{Identifier: "com.foo.bar", Version: "1.0", InstallLocation: "/usr/local/foo", Output: "foo.pkg"},
		{Identifier: "com.foo.bar", Version: "1.0", Component: "foo.app", Output: "foo.pkg"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("id: %s", tc.Identifier), func(t *testing.T) {
			err := tc.Run()
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "pkgbuild")
			assert.Contains(t, err.Error(), "is required")
		})
	}

}

func TestCommandRunEchoArgs(t *testing.T) {
	tests := []pkgbuild.Command{
		{
			Identifier:      "com.foo.app",
			Version:         "1.0",
			Component:       "foo.app",
			InstallLocation: "/usr/local/foo",
			Output:          "foo.pkg",
		},
		{
			Identifier:      "com.foo.bar",
			Version:         "1.0",
			Component:       "foo.app",
			InstallLocation: "/usr/local/foo",
			Output:          "foo.pkg",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("id: %s", tc.Identifier), func(t *testing.T) {
			tc.SetExecutable(pythonExe(), []string{"../testdata/bin/echo.py"})

			out, err := tc.GetCombinedOutput()
			assert.Error(t, err)
			assert.Empty(t, out, tc.Output)

			err = tc.Run()
			assert.NoError(t, err)

			out, err = tc.GetCombinedOutput()
			assert.NoError(t, err)
			assert.Contains(t, out, tc.Output)
			assert.Contains(t, out, fmt.Sprintf("--identifier %s", tc.Identifier))
			assert.Contains(t, out, fmt.Sprintf("--version %s", tc.Version))
			assert.Contains(t, out, fmt.Sprintf("--component %s", tc.Component))
			assert.Contains(t, out, fmt.Sprintf("--install-location %s", tc.InstallLocation))
		})
	}

}

func pythonExe() string {
	if runtime.GOOS == "darwin" {
		return "python3"
	} else {
		return "python"
	}
}
