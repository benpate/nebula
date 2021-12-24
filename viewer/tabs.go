package viewer

import (
	"strconv"

	"github.com/benpate/content"
	"github.com/benpate/html"
)

const ItemTypeTabs = "TABS"

func Tabs(b *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)
	labels := item.GetSliceOfString("labels")

	b.Div().Class("tabs")
	for index, id := range item.Refs {
		nodeID := "#id-" + strconv.Itoa(id)
		label := labels[index]
		b.A(nodeID).Class("tabs-label").InnerHTML(label).Close()
	}

	for _, id := range item.Refs {
		nodeID := "id-" + strconv.Itoa(id)
		b.Div().ID(nodeID).EndBracket()
		subTree(b, c, id)
		b.Close()
	}

	b.Close()
}
