package markdown

import "github.com/gomarkdown/markdown"

func ToHTMLBytes(text string) []byte {
	md := []byte(text)
	return markdown.ToHTML(md, nil, nil)
}

func ToHTMLString(text string) string {
	return string(ToHTMLBytes(text))
}
