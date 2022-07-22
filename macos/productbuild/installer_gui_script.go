package productbuild

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

type InstallerGuiScript struct {
	XMLName xml.Name `xml:"installer-gui-script" json:"-"`

	AuthoringTool        string `xml:"authoringTool,attr,omitempty" json:"authoringTool,omitempty"`
	AuthoringToolVersion string `xml:"authoringToolVersion,attr,omitempty" json:"authoringToolVersion,omitempty"`
	AuthoringToolBuild   string `xml:"authoringToolBuild,attr,omitempty" json:"authoringToolBuild,omitempty"`
	MinSpecVersion       string `xml:"minSpecVersion,attr,omitempty" json:"minSpecVersion,omitempty"`

	Title      string     `xml:"title,omitempty" json:"title,omitempty"`
	License    License    `xml:"license,omitempty" json:"license,omitempty"`
	Welcome    Welcome    `xml:"welcome,omitempty" json:"welcome,omitempty"`
	Conclusion Conclusion `xml:"conclusion,omitempty" json:"conclusion,omitempty"`
	Options    Options    `xml:"options,omitempty" json:"options,omitempty"`

	ChoicesOutline ChoicesOutline `xml:"choices-outline" json:"choicesOutline,omitempty"`
	Choices        []Choice       `xml:"choice" json:"choice,omitempty"`
	PkgRefs        []PkgRef       `xml:"pkg-ref" json:"pkgRef,omitempty"`
}

func NewInstallerGuiScript(name string) InstallerGuiScript {
	return InstallerGuiScript{
		AuthoringTool:        "Paket",
		AuthoringToolVersion: "0.1.0",
		AuthoringToolBuild:   "git",

		Title:   name,
		Choices: make([]Choice, 0),
		PkgRefs: make([]PkgRef, 0),
	}
}

func ReadInstallerGuiScriptFile(path string) (*InstallerGuiScript, error) {
	xmlBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	script := &InstallerGuiScript{}
	return script, xml.Unmarshal(xmlBytes, script)
}

func (s InstallerGuiScript) WriteFile(w io.Writer) error {
	file, err := xml.MarshalIndent(s, "  ", "    ")
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(xml.Header + string(file)))
	return err
}

type Options struct {
	Customize         string `xml:"customize,attr,omitempty" json:"customize,omitempty"`
	RequireScripts    string `xml:"require-scripts,attr,omitempty" json:"requireScripts,omitempty"`
	HostArchitectures string `xml:"hostArchitectures,attr,omitempty" json:"hostArchitectures,omitempty"`
}

type PkgRef struct {
	ID            string `xml:"id,attr,omitempty" json:"id,omitempty"`
	Version       string `xml:"version,attr,omitempty" json:"version,omitempty"`
	Auth          string `xml:"auth,attr,omitempty" json:"auth,omitempty"`
	InstallKBytes string `xml:"installKBytes,attr,omitempty" json:"installKBytes,omitempty"`
}

type Choice struct {
	ID          string `xml:"id,attr,omitempty" json:"id,omitempty"`
	Title       string `xml:"title,attr,omitempty" json:"title,omitempty"`
	Description string `xml:"description,attr,omitempty" json:"description,omitempty"`

	PkgRefs []PkgRef `xml:"pkg-ref,omitempty" json:"pkgRef,omitempty"`
}

type Line struct {
	Choice string  `xml:"choice,attr,omitempty" json:"choice,omitempty"`
	Lines  []*Line `xml:"line,omitempty" json:"lines,omitempty"`
}

type ChoicesOutline struct {
	Lines []Line `xml:"line" json:"line,omitempty"`
}

type License struct {
	File string `xml:"file,attr,omitempty" json:"file,omitempty"`
}

type Welcome struct {
	File string `xml:"file,attr,omitempty" json:"file,omitempty"`
}

type Conclusion struct {
	File string `xml:"file,attr,omitempty" json:"file,omitempty"`
}
