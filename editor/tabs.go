package editor

import (
	"fmt"

	"github.com/benpate/content"
	"github.com/benpate/convert"
	"github.com/benpate/html"
)

const ItemTypeTabs = "TABS"

func (e Editor) Tabs(b *html.Builder, c content.Content, id int) {

	item := c.GetItem(id)

	b.Div().Class("tabs").Script("install TabContainer")

	b.Div().Role("tablist").EndBracket()
	for index, id := range item.Refs {
		item := c.GetItem(id)
		idString := convert.String(id)

		b.Span().
			Role("tab").
			ID("tab-"+idString).
			Aria("controls", "panel-"+idString).
			Aria("selected", convert.String(index == 0)).
			EndBracket()

		b.WriteString(item.GetString("label"))

		b.Form("", "").
			Style("display:inline").
			Data("hx-post", e.Endpoint).
			Data("hx-trigger", "click").
			Data("hx-confirm", "Remove this tab?")

		b.Input("hidden", "type").Value("delete-item").Close()
		b.Input("hidden", "itemId").Value(idString).Close()
		b.Input("hidden", "check").Value(item.GetString("check")).Close()

		b.Container("i").
			Class("fa-regular fa-circle-xmark", "space-left").
			Close()

		b.Close()
		b.Close()
	}

	// Add "new" tab
	b.Span().
		Role("tab").
		ID("tab-new").
		Data("hx-post", e.Endpoint).
		Data("hx-vals", fmt.Sprintf("{'type':'new-item', 'itemId':'%s', 'place':'RIGHT', 'itemType':'CONTAINER', 'check':'%s'}", convert.String(id), item.GetString("check"))).
		EndBracket()

	b.Container("i").Class("fa-regular fa-circle-plus").Close()

	b.Close()
	b.Close()

	for index, id := range item.Refs {
		idString := convert.String(id)

		b.Div().
			Role("tabpanel").
			ID("panel-"+idString).
			Aria("labelledby", "tab-"+idString).
			Attr("hidden", convert.String(index != 0)).
			EndBracket()

		e.subTree(b, c, id)
		b.Close()
	}

	b.CloseAll()
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
