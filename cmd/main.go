package main

import (
	"fmt"
	"os"

	"github.com/moderncircuits/paket/runtime"
)

// These variables get set during link time. See Makefile.
var (
	hostOS string
	commit string
	date   string
)

func init() {
	runtime.BuildCommit = commit
	runtime.BuildDate = date
	runtime.BuildOS = hostOS
}

func main() {
	if err := execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
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
