package viewer

import (
	"github.com/benpate/content"
	"github.com/benpate/convert"
	"github.com/benpate/html"
)

const ItemTypeTabs = "TABS"

func (v Viewer) Tabs(b *html.Builder, c content.Content, id int) {

	item := c.GetItem(id)

	b.Div().Class("tabs").Script("install TabContainer")

	b.Div().Role("tablist").EndBracket()
	for index, id := range item.Refs {
		item := c.GetItem(id)
		idString := convert.String(id)

		b.Button().
			Role("tab").
			ID("tab-"+idString).
			Aria("controls", "panel-"+idString).
			Aria("selected", convert.String(index == 0)).
			InnerHTML(item.GetString("label")).
			Close()
	}
	b.Close()

	for index, id := range item.Refs {
		idString := convert.String(id)

		b.Div().
			Role("tabpanel").
			ID("panel-"+idString).
			Aria("labelledby", "tab-"+idString).
			Attr("hidden", convert.String(index != 0)).
			EndBracket()

		v.subTree(b, c, id)
		b.Close()
	}

	b.CloseAll()
}
