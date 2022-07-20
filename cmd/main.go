package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/productbuild"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var project paket.Project
	err := hclsimple.DecodeFile("test_data/config.hcl", nil, &project)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %v", err)
	}

	jsonFile, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("test.json", jsonFile, 0644)
	if err != nil {
		return err
	}

	err = productbuild.CreateDistributionXML(project)
	if err != nil {
		return err
	}

	return nil
}

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
