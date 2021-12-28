package viewer

import (
	"strconv"

	"github.com/benpate/content"
	"github.com/benpate/html"
)

const ItemTypeContainer = "CONTAINER"

func (v Viewer) Container(b *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)

	b.Div().
		Class("container").
		Data("style", item.GetString("style")).
		Data("size", strconv.Itoa(len(item.Refs)))

	for _, id := range item.Refs {
		b.Div().Class("container-item")
		v.subTree(b, c, id)
		b.Close()
	}
}
