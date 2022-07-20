package paket

import (
	"os"

	"github.com/moderncircuits/paket/innosetup"
)

func runWindowsInnoSetup(project Project) error {
	iss := innosetup.NewInnoSetupScript(project.Name, project.Vendor)
	w, err := os.Create("innosetup.iss")
	if err != nil {
		return err
	}
	defer w.Close()
	return iss.Run(w)
}
