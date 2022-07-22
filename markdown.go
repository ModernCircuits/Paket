// SPDX-License-Identifier: BSL-1.0

package paket

import (
	"io/ioutil"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yosssi/gohtml"
)

// MarkdownToHTML converts a markdown byte slice to HTML.
// The output is sanitized using bluemonday.
func MarkdownToHTML(md []byte) []byte {
	maybeUnsafeHTML := markdown.ToHTML(md, nil, nil)
	return bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)
}

// MarkdownFileToHTML converts a markdown file to HTML.
// The output is sanitized using bluemonday.
func MarkdownFileToHTML(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	return MarkdownToHTML(content)
}

// FormatHTML converts HTML to pretty HTML.
func FormatHTML(html []byte) []byte {
	return gohtml.FormatBytes(html)
}
