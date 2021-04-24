package content

import (
	"github.com/benpate/html"
	"github.com/benpate/htmlconv"
)

const ItemTypeText = "TEXT"

func TextViewer(lib *Library, b *html.Builder, item *Item) {
	content := item.AsString("text")
	content = htmlconv.FromText(content)
	b.WriteString(content)
}

func TextCreator(lib *Library, b *html.Builder, item *Item) {
	b.Container("textarea").Name(item.Path).Close()
}

func TextEditor(lib *Library, b *html.Builder, item *Item) {
	content := item.AsString("text")
	b.Container("textarea").Name(item.Path).InnerHTML(content).Close()
}
