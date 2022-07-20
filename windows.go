package paket

import (
	"encoding/json"
	"io/ioutil"

	"github.com/moderncircuits/paket/innosetup"
)

func runWindowsInnoSetup(project Project) error {
	iss := innosetup.NewInnoSetupScript(project.Name, "Modern Circuits")

	issJSON, err := json.MarshalIndent(iss, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("iss.json", issJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}
