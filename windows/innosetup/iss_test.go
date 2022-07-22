package innosetup_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/windows/innosetup"
	"github.com/stretchr/testify/assert"
)

func TestNewISS(t *testing.T) {
	project, err := paket.ReadProjectConfigFile("../../testdata/full.hcl")
	assert.NoError(t, err)
	assert.NotNil(t, project)

	iss := innosetup.NewISS(*project)
	assert.Equal(t, "Plugin Template", iss.Setup.AppName)
	assert.Equal(t, "Modern Circuits", iss.Setup.AppPublisher)
	assert.Equal(t, "modern", iss.Setup.WizardStyle)
	assert.Equal(t, false, iss.Setup.WizardResizable)

	assert.Empty(t, iss.Setup.AppVersion)
	assert.Empty(t, iss.Setup.LicenseFile)

	buf := new(bytes.Buffer)
	err = iss.WriteFile(buf)
	assert.NoError(t, err)

	str := buf.String()
	assert.Contains(t, str, `AppName="Plugin Template"`)
	assert.Contains(t, str, `AppPublisher="Modern Circuits"`)
	assert.Contains(t, str, `WizardStyle="modern"`)
	assert.Contains(t, str, `WizardResizable=no`)
}

func TestReadFile(t *testing.T) {
	t.Run("NonExist", func(t *testing.T) {
		iss, err := innosetup.ReadFile("testdata/NonExist.iss")
		assert.Error(t, err)
		assert.Nil(t, iss)
	})

	t.Run("ErrNoSetup", func(t *testing.T) {
		iss, err := innosetup.ReadFile("testdata/ErrNoSetup.iss")
		assert.Error(t, err)
		assert.Nil(t, iss)
	})

	t.Run("Example1", func(t *testing.T) {
		iss, err := innosetup.ReadFile("testdata/Example1.iss")
		assert.NoError(t, err)
		assert.NotNil(t, iss)
		assert.Equal(t, "My Program", iss.Setup.AppName)
		assert.Equal(t, "1.5", iss.Setup.AppVersion)
		assert.Equal(t, "modern", iss.Setup.WizardStyle)
		assert.Equal(t, "{autopf}\\My Program", iss.Setup.DefaultDirName)
		assert.Equal(t, "My Program", iss.Setup.DefaultGroupName)
		assert.Equal(t, "{app}\\MyProg.exe", iss.Setup.UninstallDisplayIcon)
		assert.Equal(t, "lzma2", iss.Setup.Compression)
		assert.Equal(t, "userdocs:Inno Setup Examples Output", iss.Setup.OutputDir)
		assert.True(t, iss.Setup.SolidCompression)
	})

}

func TestReadFileNoErrors(t *testing.T) {
	tests := []string{
		"64Bit",
		"Example1",
		"Example2",
		"Example3",
	}

	for _, tc := range tests {
		file := fmt.Sprintf("testdata/%s.iss", tc)
		t.Run(file, func(t *testing.T) {
			setup, err := innosetup.ReadFile(file)
			assert.NoError(t, err)
			assert.NotNil(t, setup)
		})
	}
}

func TestParseWithErrors(t *testing.T) {
	tests := []string{
		"NoKey",
		"NoSetup",
		"NoValue",
	}

	for _, tc := range tests {
		file := fmt.Sprintf("testdata/err/%s.iss", tc)
		t.Run(file, func(t *testing.T) {
			setup, err := innosetup.ReadFile(file)
			assert.Error(t, err)
			assert.Nil(t, setup)
		})
	}
}
