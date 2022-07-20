package paket

import (
	"github.com/moderncircuits/paket/innosetup"
)

func runWindowsInnoSetup(project Project) error {
	iss := innosetup.NewInnoSetupScript(project.Name, project.Vendor)
	return iss.Run()
}
