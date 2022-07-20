package innosetup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// SetupSection represents the [setup] section of an iss file.
// https://jrsoftware.org/ishelp/index.php?topic=setupsection
type SetupSection struct {
	AppId           string `iss:"AppId,omitempty" json:"AppId,omitempty"`
	AppName         string `iss:"AppName,omitempty" json:"AppName,omitempty"`
	AppPublisher    string `iss:"AppPublisher,omitempty" json:"AppPublisher,omitempty"`
	AppPublisherURL string `iss:"AppPublisherURL,omitempty" json:"AppPublisherURL,omitempty"`
	AppSupportURL   string `iss:"AppSupportURL,omitempty" json:"AppSupportURL,omitempty"`
	AppUpdatesURL   string `iss:"AppUpdatesURL,omitempty" json:"AppUpdatesURL,omitempty"`
	AppVersion      string `iss:"AppVersion,omitempty" json:"AppVersion,omitempty"`

	ArchitecturesAllowed            string `iss:"ArchitecturesAllowed,omitempty" json:"ArchitecturesAllowed,omitempty"`
	ArchitecturesInstallIn64BitMode string `iss:"ArchitecturesInstallIn64BitMode,omitempty" json:"ArchitecturesInstallIn64BitMode,omitempty"`

	Compression      string `iss:"Compression,omitempty" json:"Compression,omitempty"`
	DefaultDirName   string `iss:"DefaultDirName,omitempty" json:"DefaultDirName,omitempty"`
	DefaultGroupName string `iss:"DefaultGroupName,omitempty" json:"DefaultGroupName,omitempty"`

	LicenseFile string `iss:"LicenseFile,omitempty" json:"LicenseFile,omitempty"`

	InfoAfterFile  string `iss:"InfoAfterFile,omitempty" json:"InfoAfterFile,omitempty"`
	InfoBeforeFile string `iss:"InfoBeforeFile,omitempty" json:"InfoBeforeFile,omitempty"`

	OutputDir          string `iss:"OutputDir,omitempty" json:"OutputDir,omitempty"`
	OutputBaseFilename string `iss:"OutputBaseFilename,omitempty" json:"OutputBaseFilename,omitempty"`
	SolidCompression   bool   `iss:"SolidCompression,omitempty" json:"SolidCompression,omitempty"`
	SetupLogging       bool   `iss:"SetupLogging,omitempty" json:"SetupLogging,omitempty"`

	UninstallDisplayIcon     string `iss:"UninstallDisplayIcon,omitempty" json:"UninstallDisplayIcon,omitempty"`
	UninstallFilesDir        string `iss:"UninstallFilesDir,omitempty" json:"UninstallFilesDir,omitempty"`
	UninstallRestartComputer bool   `iss:"UninstallRestartComputer,omitempty" json:"UninstallRestartComputer,omitempty"`

	SetupIconFile        string `iss:"SetupIconFile,omitempty" json:"SetupIconFile,omitempty"`
	WizardImageFile      string `iss:"WizardImageFile,omitempty" json:"WizardImageFile,omitempty"`
	WizardSmallImageFile string `iss:"WizardSmallImageFile,omitempty" json:"WizardSmallImageFile,omitempty"`
	WizardResizable      bool   `iss:"WizardResizable,omitempty" json:"WizardResizable,omitempty"`
	WizardStyle          string `iss:"WizardStyle,omitempty" json:"WizardStyle,omitempty"`
}

type InnoSetupScript struct {
	Setup SetupSection `json:"setup"`
}

func NewInnoSetupScript(projectName string, vendor string) InnoSetupScript {
	return InnoSetupScript{
		Setup: SetupSection{
			AppName:      projectName,
			AppPublisher: vendor,

			DefaultGroupName:   projectName,
			DefaultDirName:     fmt.Sprintf("{commonpf}\\%s\\%s", vendor, projectName),
			OutputBaseFilename: fmt.Sprintf("%s Installer", projectName),

			Compression:      "lzma2",
			SolidCompression: true,
			SetupLogging:     true,

			ArchitecturesAllowed:            "x64",
			ArchitecturesInstallIn64BitMode: "x64",

			UninstallDisplayIcon:     "{app}\\Plugin Template.exe",
			UninstallFilesDir:        "{app}\\uninst",
			UninstallRestartComputer: false,

			WizardResizable: false,
			WizardStyle:     "modern",
		},
	}
}

func (iss InnoSetupScript) Run() error {
	if iss.Setup.AppId == "" {
		id, err := NewInnoSetupGUID()
		if err != nil {
			return err
		}
		iss.Setup.AppId = id
	}

	js, err := json.MarshalIndent(iss, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("iss.json", js, 0644)
	if err != nil {
		return err
	}

	issType := reflect.TypeOf(iss)
	examiner(issType, 0)

	return nil
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("  ", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("  ", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Type.Kind() == reflect.Struct {
				examiner(f.Type, depth+1)
			}
			if f.Tag.Get("iss") != "" {
				indent := strings.Repeat("  ", depth+1)
				idx := i + 1
				fmt.Printf("%s %d: %s Type: %s\n", indent, idx, f.Name, f.Type.Name())
			}
		}
	}
}
