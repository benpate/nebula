package viewer

import (
	"github.com/benpate/content"
	"github.com/benpate/html"
	"github.com/benpate/htmlconv"
)

const ItemTypeText = "TEXT"

func (v Viewer) Text(b *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)
	result := item.GetString("text")
	result = htmlconv.FromText(result)
	b.WriteString(result)
}
