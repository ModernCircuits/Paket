package textconv

import (
	"io/ioutil"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yosssi/gohtml"
)

func MarkdownToHTML(md []byte) []byte {
	maybeUnsafeHTML := markdown.ToHTML(md, nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)
	return gohtml.FormatBytes(html)
}

func MarkdownFileToHTML(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	return MarkdownToHTML(content)
}
