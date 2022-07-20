package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/hashicorp/hcl/v2"
	"github.com/moderncircuits/paket"
	"github.com/moderncircuits/paket/innosetup"
	"github.com/zclconf/go-cty/cty"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	windowsConstants := innosetup.DirectoryConstants()
	windowsVars := map[string]cty.Value{}
	for _, constant := range windowsConstants {
		windowsVars[constant] = cty.StringVal(fmt.Sprintf("{%s}", constant))
	}
	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"windows": cty.ObjectVal(windowsVars),
		},
	}

	project, err := paket.ReadFile("testdata/minimal.hcl", ctx)
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

	err = project.Run(runtime.GOOS)
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
