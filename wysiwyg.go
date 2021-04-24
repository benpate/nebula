package content

import (
	"github.com/benpate/html"
)

const ItemTypeWYSIWYG = "WYSIWYG"

func WYSIWYGViewer(lib *Library, b *html.Builder, item *Item) {
	content := item.GetString("html")
	b.WriteString(content)
}

func WYSIWYGEditor(lib *Library, b *html.Builder, item *Item) {
	content := item.GetString("html")
	b.Container("textarea").Name(item.Path).InnerHTML(content).Close()
}
