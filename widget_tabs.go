package nebula

import (
	"fmt"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/maps"
)

// ItemTypeTabs describes a multi-tab container layout
const ItemTypeTabs = "TABS"

type Tabs struct {
	library *Library
}

// Init appends three empty tabs into this tab control.
func (w Tabs) Init(container *Container, id int) {

	// Let's add THREE new tabs
	for index := 1; index <= 3; index++ {
		itemID := container.NewItemWithInit(w.library, ItemTypeLayout, maps.Map{
			"style": LayoutStyleRows,
			"label": "Tab " + convert.String(index),
		})
		(*container)[id].AddReference(itemID, 0) // Add a reference to the new container into me
	}
}

// View displays that tabs in HTML, along with sub-elements
func (w Tabs) View(b *html.Builder, container *Container, id int) {

	item := container.GetItem(id)

	b.Div().Class("tabs").Script("install TabContainer")

	b.Div().Role("tablist").EndBracket()
	for index, id := range item.Refs {
		item := container.GetItem(id)
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

		w.library.View(b, container, id)
		b.Close()
	}

	b.CloseAll()
}

func (w Tabs) Edit(b *html.Builder, container *Container, id int, endpoint string) {

	item := container.GetItem(id)

	b.Div().Class("tabs").Script("install TabContainer")

	b.Div().Role("tablist").EndBracket()
	for index, id := range item.Refs {
		item := container.GetItem(id)
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
			Data("hx-post", endpoint).
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
		Data("hx-post", endpoint).
		Data("hx-vals", fmt.Sprintf("{'action':'new-item', 'itemId':'%s', 'place':'RIGHT', 'itemType':'CONTAINER', 'check':'%s'}", convert.String(id), item.GetString("check"))).
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

		w.library.Edit(b, container, id, endpoint)
		b.Close()
	}

	b.CloseAll()
}

// Validate cleans the container for invalid content
func (w Tabs) Validate(container *Container, index int) {

	for _, id := range (*container)[index].Refs {
		item := container.GetItem(id)
		originalText := item.GetString("label")
		cleanText := textPolicy.Sanitize(originalText)
		item.Data.SetString("label", cleanText)
	}
}
