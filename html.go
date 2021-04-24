package content

import (
	"github.com/benpate/html"
)

const ItemTypeHTML = "HTML"

func HTMLViewer(lib *Library, b *html.Builder, item *Item) {
	content := item.AsString("html")
	b.WriteString(content)
}

func HTMLCreator(lib *Library, b *html.Builder, item *Item) {
	b.Container("textarea").Name(item.Path).Close()
}

func HTMLEditor(lib *Library, b *html.Builder, item *Item) {
	content := item.AsString("html")
	b.Container("textarea").Name(item.Path).InnerHTML(content).Close()
}
