package paket

import (
	"io/ioutil"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yosssi/gohtml"
)

func MarkdownToHTML(md []byte) []byte {
	maybeUnsafeHTML := markdown.ToHTML(md, nil, nil)
	return bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)
}

func MarkdownFileToHTML(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	return MarkdownToHTML(content)
}

func FormatHTML(html []byte) []byte {
	return gohtml.FormatBytes(html)
}
