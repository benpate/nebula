package viewer

import (
	"github.com/benpate/content"
	"github.com/benpate/html"
)

const ItemTypeWYSIWYG = "WYSIWYG"

func (v Viewer) WYSIWYG(b *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)
	result := item.GetString("html")
	b.WriteString(result)
}
