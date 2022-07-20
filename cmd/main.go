package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/moderncircuits/paket"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	project, err := paket.NewProject("testdata/full.hcl")
	if err != nil {
		return err
	}

	err = project.Run(runtime.GOOS)
	if err != nil {
		return err
	}

	err = project.Run("darwin")
	if err != nil {
		return err
	}

	jsonFile, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("test.json", jsonFile, 0644)
	if err != nil {
		return err
	}

	return nil
}

// xmlFile, err := ioutil.ReadFile("productbuild/testdata/distribution.xml")
// if err != nil {
// 	fmt.Println(err)
// }

// {
// 	html := textconv.MarkdownFileToHTML("testdata/simple_markdown.md")
// 	fmt.Println(string(html))
// }
// {
// 	html := textconv.MarkdownFileToHTML("LICENSE.txt")
// 	fmt.Println(string(html))
// }
