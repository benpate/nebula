package content

import (
	"github.com/benpate/convert"
	"github.com/benpate/html"
)

const ItemTypeTabs = "TABS"

func TabsViewer(lib *Library, builder *html.Builder, item *Item) {
	labels := convert.SliceOfString(item.AsInterface("labels"))

	builder.Div().Class("tabs")
	for index, label := range labels {
		id := convert.String(index)
		builder.A("#tab-" + id).Class("tabs-label").InnerHTML(label).Close()
	}

	for index := range item.Children {
		id := convert.String(index)
		builder.Div().ID("tab-" + id).EndBracket()
		lib.SubTree(builder, &item.Children[index])
		builder.Close()
	}

	builder.Close()
}

func TabsCreator(lib *Library, builder *html.Builder, path string, item *Item) {
	builder.Container("textarea").Name(path).Close()
}

func TabsEditor(lib *Library, builder *html.Builder, item *Item) {
	content := item.AsString("html")
	builder.Container("textarea").Name(item.Path).InnerHTML(content).Close()
}
