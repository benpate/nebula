package content

import (
	"github.com/benpate/html"
	"github.com/benpate/htmlconv"
)

const ItemTypeText = "TEXT"

func TextViewer(lib *Library, b *html.Builder, item *Item) {
	content := item.GetString("text")
	content = htmlconv.FromText(content)
	b.WriteString(content)
}

func TextEditor(lib *Library, b *html.Builder, item *Item) {
	content := item.GetString("text")
	b.Container("textarea").Name(item.Path).InnerHTML(content).Close()
}
