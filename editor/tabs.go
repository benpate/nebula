package editor

import (
	"strconv"

	"github.com/benpate/content"
	"github.com/benpate/html"
)

func (e Editor) Tabs(builder *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)
	labels := item.GetSliceOfString("labels")

	builder.Div().Class("tabs")
	for index, id := range item.Refs {
		nodeID := "#id-" + strconv.Itoa(id)
		label := labels[index]
		builder.A(nodeID).Class("tabs-label").InnerHTML(label).Close()
	}

	for _, id := range item.Refs {
		nodeID := "id-" + strconv.Itoa(id)
		builder.Div().ID(nodeID).EndBracket()
		e.subTree(builder, c, id)
		builder.Close()
	}

	builder.Close()
}

/*
func (widget Tabs) DefaultChildren() []content.Item {
	return []content.Item{
		{
			Type: "CONTAINER",
			Data: datatype.Map{
				"style": "COLUMNS",
			},
		},
		{
			Type: "CONTAINER",
			Data: datatype.Map{
				"style": "COLUMNS",
			},
		},
	}
}
*/
