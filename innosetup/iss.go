package innosetup

import "fmt"

// SetupSection represents the [setup] section of an iss file.
// https://jrsoftware.org/ishelp/index.php?topic=setupsection
type SetupSection struct {
	AppId           string `json:"AppId,omitempty"`
	AppName         string `json:"AppName,omitempty"`
	AppPublisher    string `json:"AppPublisher,omitempty"`
	AppPublisherURL string `json:"AppPublisherURL,omitempty"`
	AppSupportURL   string `json:"AppSupportURL,omitempty"`
	AppUpdatesURL   string `json:"AppUpdatesURL,omitempty"`
	AppVersion      string `json:"AppVersion,omitempty"`

	ArchitecturesAllowed            string `json:"ArchitecturesAllowed,omitempty"`
	ArchitecturesInstallIn64BitMode string `json:"ArchitecturesInstallIn64BitMode,omitempty"`

	Compression      string `json:"Compression,omitempty"`
	DefaultDirName   string `json:"DefaultDirName,omitempty"`
	DefaultGroupName string `json:"DefaultGroupName,omitempty"`

	LicenseFile string `json:"LicenseFile,omitempty"`

	InfoAfterFile  string `json:"InfoAfterFile,omitempty"`
	InfoBeforeFile string `json:"InfoBeforeFile,omitempty"`

	OutputDir          string `json:"OutputDir,omitempty"`
	OutputBaseFilename string `json:"OutputBaseFilename,omitempty"`
	SolidCompression   bool   `json:"SolidCompression,omitempty"`
	SetupLogging       bool   `json:"SetupLogging,omitempty"`

	UninstallDisplayIcon     string `json:"UninstallDisplayIcon,omitempty"`
	UninstallFilesDir        string `json:"UninstallFilesDir,omitempty"`
	UninstallRestartComputer bool   `json:"UninstallRestartComputer,omitempty"`

	SetupIconFile        string `json:"SetupIconFile,omitempty"`
	WizardImageFile      string `json:"WizardImageFile,omitempty"`
	WizardSmallImageFile string `json:"WizardSmallImageFile,omitempty"`
	WizardResizable      bool   `json:"WizardResizable,omitempty"`
	WizardStyle          string `json:"WizardStyle,omitempty"`
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
		return fmt.Errorf("AppId is required")
	}
	return nil
}
