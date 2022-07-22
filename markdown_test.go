// SPDX-License-Identifier: BSL-1.0

package paket_test

import (
	"testing"

	"github.com/moderncircuits/paket"
	"github.com/stretchr/testify/assert"
)

func TestMarkdownToHTML(t *testing.T) {
	// Paragraph
	p := paket.MarkdownToHTML([]byte("This is some text"))
	assert.NotEmpty(t, p)
	assert.Contains(t, string(p), "<p>This is some text</p>")

	// Heading
	h1 := paket.MarkdownToHTML([]byte("# Foo"))
	assert.NotEmpty(t, h1)
	assert.Contains(t, string(h1), "<h1>Foo</h1>")

	h2 := paket.MarkdownToHTML([]byte("## Foo"))
	assert.NotEmpty(t, h2)
	assert.Contains(t, string(h2), "<h2>Foo</h2>")
}

func TestMarkdownFileToHTML(t *testing.T) {
	noexist := paket.MarkdownFileToHTML("/noexist/noexist.md")
	assert.Empty(t, noexist)

	file := paket.MarkdownFileToHTML("testdata/txt/simple_markdown.md")
	assert.NotEmpty(t, file)
	assert.Contains(t, string(file), "ASIC Filter")
	assert.Contains(t, string(file), "Foo Bar Baz")

	formatted := paket.FormatHTML(file)
	assert.NotEmpty(t, formatted)
	assert.Contains(t, string(formatted), "ASIC Filter")
	assert.Contains(t, string(formatted), "Foo Bar Baz")
}
