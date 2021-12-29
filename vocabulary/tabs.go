package vocabulary

import (
	"fmt"

	"github.com/benpate/convert"
	"github.com/benpate/html"
	"github.com/benpate/nebula"
)

const ItemTypeTabs = "TABS"

type Tabs struct {
	library nebula.Library
}

// Init appends three empty tabs into this tab control.
func (w Tabs) Init(container *nebula.Container, id int) {

	me := container.GetItem(id)

	// Let's add THREE new tabs
	for index := 1; index <= 3; index++ {
		itemID := container.NewItem(w.library, ItemTypeLayout) // Make a new container
		item := container.GetItem(itemID)                      // Get a pointer to the item in the content structure
		item.Set("label", "Tab "+convert.String(index))        // Add a custom "label" for the tab
		me.AddReference(itemID, 0)                             // Add a reference to the new container into me
	}
}

// View displays that tabs in HTML, along with sub-elements
func (w Tabs) View(b *html.Builder, container *nebula.Container, id int) {

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

func (w Tabs) Edit(b *html.Builder, container *nebula.Container, id int, endpoint string) {

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

		w.library.Edit(b, container, id, endpoint)
		b.Close()
	}

	b.CloseAll()
}
