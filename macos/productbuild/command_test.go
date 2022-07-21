package productbuild_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/moderncircuits/paket/macos/productbuild"
	"github.com/stretchr/testify/assert"
)

func TestCommandMissingBinary(t *testing.T) {
	script := productbuild.Command{
		Distribution: productbuild.InstallerGuiScript{},
		ResourcePath: ".",
		PackagePath:  "./pkg",
		OutputFile:   "foo.pkg",
	}
	script.SetExecutable("/bin/doesnotexist", nil)
	err := script.Run()
	assert.Error(t, err)
}

func TestCommandRunMissingArgs(t *testing.T) {
	script := productbuild.InstallerGuiScript{
		Title: "Foo",
	}
	tests := []productbuild.Command{
		{},
		{Distribution: script},
		{Distribution: script, ResourcePath: "."},
		{Distribution: script, ResourcePath: ".", PackagePath: "./pkg"},
		{Distribution: script, ResourcePath: ".", PackagePath: "./pkg"},
		{Distribution: script, ResourcePath: ".", OutputFile: "foo.pkg"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			err := tc.Run()
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "productbuild")
			assert.Contains(t, err.Error(), "is required")
		})
	}

}

func TestCommandRunEchoArgs(t *testing.T) {
	tests := []productbuild.Command{
		{
			Distribution: productbuild.InstallerGuiScript{},
			ResourcePath: ".",
			PackagePath:  "./pkg",
			OutputFile:   "foo.pkg",
		},
		{
			Distribution: productbuild.InstallerGuiScript{},
			ResourcePath: ".",
			PackagePath:  "./pkg",
			OutputFile:   "foo.pkg",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.Distribution), func(t *testing.T) {
			tc.SetExecutable(pythonExe(), []string{"../../testdata/bin/echo.py"})

			out, err := tc.GetCombinedOutput()
			assert.Error(t, err)
			assert.Empty(t, out, tc.OutputFile)

			err = tc.Run()
			assert.NoError(t, err)

			out, err = tc.GetCombinedOutput()
			assert.NoError(t, err)
			assert.Contains(t, out, tc.OutputFile)
			assert.Contains(t, out, fmt.Sprintf("--distribution %s", "dist.xml"))
			assert.Contains(t, out, fmt.Sprintf("--resources %s", tc.ResourcePath))
			assert.Contains(t, out, fmt.Sprintf("--package-path %s", tc.PackagePath))
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
