package content

import "github.com/benpate/html"

const ItemTypeRows = "ROWS"

func RowsViewer(library *Library, builder *html.Builder, item *Item) {
	builder.Div().Class("rows rows-" + item.GetString("style"))
	for index := range item.Children {
		library.SubTree(builder, &item.Children[index])
	}
}
