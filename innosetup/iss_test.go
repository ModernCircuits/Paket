package innosetup_test

import (
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
	}
}
