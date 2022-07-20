package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {

	// xmlFile, err := ioutil.ReadFile("productbuild/test_data/distribution.xml")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// {
	// 	html := textconv.MarkdownFileToHTML("test_data/simple_markdown.md")
	// 	fmt.Println(string(html))
	// }
	// {
	// 	html := textconv.MarkdownFileToHTML("LICENSE.txt")
	// 	fmt.Println(string(html))
	// }

	installer := Installer{
		ProductName:    "Plugin Template",
		ProductVersion: "1.0.0",
		Identifier:     "com.moderncircuits.plugin-template",

		MacOS: MacOS{
			Components: []Component{
				{
					Tag:         "AU",
					PayloadPath: "macOS/AU/Plugin Template.component",
					InstallPath: "/Library/Audio/Plug-Ins/AU",
				},
				{
					Tag:         "VST3",
					PayloadPath: "macOS/VST3/Plugin Template.vst3",
					InstallPath: "/Library/Audio/Plug-Ins/VST3",
				},
				{
					Tag:         "CLAP",
					PayloadPath: "macOS/CLAP/Plugin Template.clap",
					InstallPath: "/Library/Audio/Plug-Ins/CLAP",
				},
			},
		},
	}

	return installer.Create()
}
