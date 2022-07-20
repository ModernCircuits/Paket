package innosetup_test

import (
	"bytes"
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/innosetup"
	"github.com/stretchr/testify/assert"
)

func TestNewInnoSetupScript(t *testing.T) {
	{
		project, err := paket.NewProject("../testdata/full.hcl")
		assert.NoError(t, err)

		iss := innosetup.NewInnoSetupScript(project.Name, project.Vendor)
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
}
