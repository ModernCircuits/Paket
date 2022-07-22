package innosetup

import (
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/moderncircuits/paket"
)

type ISS struct {
	Setup Setup `json:"setup"`
}

func NewISS(project paket.ProjectConfig) ISS {
	return ISS{
		Setup: Setup{
			AppName:      project.Name,
			AppPublisher: project.Vendor,

			DefaultGroupName:   project.Name,
			DefaultDirName:     fmt.Sprintf("{commonpf}\\%s\\%s", project.Vendor, project.Name),
			OutputBaseFilename: fmt.Sprintf("%s Installer", project.Name),

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

func (iss ISS) WriteFile(w io.Writer) error {
	e := reflect.ValueOf(&iss.Setup).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		switch varType.Kind() {
		case reflect.String:
			if str := varValue.(string); str != "" {
				fmt.Fprintf(w, "%v=%q\n", varName, str)
			}
		case reflect.Bool:
			txt := "no"
			if varValue.(bool) {
				txt = "yes"
			}
			fmt.Fprintf(w, "%v=%v\n", varName, txt)
		default:
			return fmt.Errorf("unimplemented type %v", varType)
		}
	}

	return nil
}

func ReadFile(path string) (*ISS, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Parse(f)
}

func Parse(r io.Reader) (*ISS, error) {
	lines := removeAllCommentLines(readAllLines(r))

	setupLines, err := getSetupLines(lines)
	if err != nil {
		return nil, err
	}

	setup, err := parseSetup(setupLines)
	if err != nil {
		return nil, err
	}

	return &ISS{
		Setup: *setup,
	}, nil
}

// Setup represents the [Setup] section of an iss file.
// https://jrsoftware.org/ishelp/index.php?topic=setupsection
type Setup struct {
	AppId           string `iss:"AppId,omitempty" json:"AppId,omitempty"`
	AppName         string `iss:"AppName,omitempty" json:"AppName,omitempty"`
	AppPublisher    string `iss:"AppPublisher,omitempty" json:"AppPublisher,omitempty"`
	AppPublisherURL string `iss:"AppPublisherURL,omitempty" json:"AppPublisherURL,omitempty"`
	AppSupportURL   string `iss:"AppSupportURL,omitempty" json:"AppSupportURL,omitempty"`
	AppUpdatesURL   string `iss:"AppUpdatesURL,omitempty" json:"AppUpdatesURL,omitempty"`
	AppVersion      string `iss:"AppVersion,omitempty" json:"AppVersion,omitempty"`

	ArchitecturesAllowed            string `iss:"ArchitecturesAllowed,omitempty" json:"ArchitecturesAllowed,omitempty"`
	ArchitecturesInstallIn64BitMode string `iss:"ArchitecturesInstallIn64BitMode,omitempty" json:"ArchitecturesInstallIn64BitMode,omitempty"`

	ChangesAssociations     bool   `iss:"ChangesAssociations,omitempty" json:"ChangesAssociations,omitempty"`
	Compression             string `iss:"Compression,omitempty" json:"Compression,omitempty"`
	DefaultDirName          string `iss:"DefaultDirName,omitempty" json:"DefaultDirName,omitempty"`
	DefaultGroupName        string `iss:"DefaultGroupName,omitempty" json:"DefaultGroupName,omitempty"`
	DisableProgramGroupPage bool   `iss:"DisableProgramGroupPage,omitempty" json:"DisableProgramGroupPage,omitempty"`

	LicenseFile string `iss:"LicenseFile,omitempty" json:"LicenseFile,omitempty"`

	InfoAfterFile  string `iss:"InfoAfterFile,omitempty" json:"InfoAfterFile,omitempty"`
	InfoBeforeFile string `iss:"InfoBeforeFile,omitempty" json:"InfoBeforeFile,omitempty"`

	OutputDir                          string `iss:"OutputDir,omitempty" json:"OutputDir,omitempty"`
	OutputBaseFilename                 string `iss:"OutputBaseFilename,omitempty" json:"OutputBaseFilename,omitempty"`
	PrivilegesRequiredOverridesAllowed bool   `iss:"PrivilegesRequiredOverridesAllowed,omitempty" json:"PrivilegesRequiredOverridesAllowed,omitempty"`

	SolidCompression bool `iss:"SolidCompression,omitempty" json:"SolidCompression,omitempty"`
	SetupLogging     bool `iss:"SetupLogging,omitempty" json:"SetupLogging,omitempty"`

	UninstallDisplayIcon     string `iss:"UninstallDisplayIcon,omitempty" json:"UninstallDisplayIcon,omitempty"`
	UninstallFilesDir        string `iss:"UninstallFilesDir,omitempty" json:"UninstallFilesDir,omitempty"`
	UninstallRestartComputer bool   `iss:"UninstallRestartComputer,omitempty" json:"UninstallRestartComputer,omitempty"`

	UserInfoPage bool `iss:"UserInfoPage,omitempty" json:"UserInfoPage,omitempty"`

	SetupIconFile        string `iss:"SetupIconFile,omitempty" json:"SetupIconFile,omitempty"`
	WizardImageFile      string `iss:"WizardImageFile,omitempty" json:"WizardImageFile,omitempty"`
	WizardSmallImageFile string `iss:"WizardSmallImageFile,omitempty" json:"WizardSmallImageFile,omitempty"`
	WizardResizable      bool   `iss:"WizardResizable,omitempty" json:"WizardResizable,omitempty"`
	WizardStyle          string `iss:"WizardStyle,omitempty" json:"WizardStyle,omitempty"`
}
