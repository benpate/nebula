package content

import (
	"github.com/benpate/html"
)

const ItemTypeColumns = "COLUMNS"

func ColumnsViewer(library *Library, builder *html.Builder, item *Item) {
	builder.Div().Class("columns columns-" + item.GetString("style"))
	for index := range item.Children {
		builder.Div().Class("columns-column").EndBracket()
		library.SubTree(builder, &item.Children[index])
		builder.Close()
	}
	builder.Close()
}
