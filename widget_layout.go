package vocabulary

import (
	"net/url"
	"strconv"

	"github.com/benpate/convert"
	"github.com/benpate/first"
	"github.com/benpate/html"
	"github.com/benpate/nebula"
)

const ItemTypeLayout = "LAYOUT"

type Layout struct {
	library *nebula.Library
}

// Init adds a child WYSIWYG element
func (w Layout) Init(container *nebula.Container, id int) {
	wysiwyg := container.NewItemWithInit(w.library, ItemTypeWYSIWYG, nil)
	(*container)[id].AddReference(wysiwyg, 0)
}

// View dsplays the layout and its children.
func (w Layout) View(b *html.Builder, container *nebula.Container, layoutID int) {

	item := container.GetItem(layoutID)

	if len(item.Refs) == 0 {
		return
	}

	b.Div().
		Class("nabula-layout").
		Data("style", item.GetString("style")).
		Data("size", strconv.Itoa(len(item.Refs)))

	for _, id := range item.Refs {
		b.Div().Class("nebula-layout-item")
		w.library.View(b, container, id)
		b.Close()
	}
}

func (w Layout) Edit(b *html.Builder, container *nebula.Container, layoutID int, endpoint string) {

	layout := container.GetItem(layoutID)
	layoutIDString := strconv.Itoa(layoutID)
	style := first.String(layout.GetString("style"), nebula.LayoutStyleRows)

	b.Div().
		Class("nebula-layout").
		Data("style", style).
		Data("size", strconv.Itoa(len(layout.Refs))).
		Data("id", layoutIDString)

	layoutInsert(b, layoutIDString, layoutIDString, nebula.LayoutPlaceAbove, layout.Check, endpoint)
	layoutInsert(b, layoutIDString, layoutIDString, nebula.LayoutPlaceLeft, layout.Check, endpoint)

	for _, childID := range layout.Refs {
		childIDString := strconv.Itoa(childID)

		b.Div().Class("nebula-layout-item")

		layoutInsert(b, layoutIDString, childIDString, nebula.LayoutPlaceAbove, layout.Check, endpoint)
		layoutInsert(b, layoutIDString, childIDString, nebula.LayoutPlaceLeft, layout.Check, endpoint)

		w.library.Edit(b, container, childID, endpoint)

		layoutInsert(b, layoutIDString, childIDString, nebula.LayoutPlaceBelow, layout.Check, endpoint)
		layoutInsert(b, layoutIDString, childIDString, nebula.LayoutPlaceRight, layout.Check, endpoint)

		b.Close()
	}

	layoutInsert(b, layoutIDString, layoutIDString, nebula.LayoutPlaceBelow, layout.Check, endpoint)
	layoutInsert(b, layoutIDString, layoutIDString, nebula.LayoutPlaceRight, layout.Check, endpoint)

	b.Close()
}

func (w Layout) Prop(b *html.Builder, container *nebula.Container, id int, endpoint string, params url.Values) error {

	b.H1().InnerHTML("Add Another Section").Close()

	b.Div().Class("table")
	for _, itemType := range ItemTypes() {
		b.Div().Attr("tabindex", "0")
		b.Form("", "").Data("hx-post", endpoint).Data("hx-trigger", "click")
		b.Input("hidden", "type").Value("add-item").Close()
		b.Input("hidden", "itemId").Value(params.Get("itemId")).Close()
		b.Input("hidden", "itemType").Value(itemType.Code)

		for key, value := range itemType.Data {
			b.Input("hidden", key).Value(convert.String(value)).Close()
		}

		b.Input("hidden", "place").Value(params.Get("place")).Close()
		b.Input("hidden", "check").Value(params.Get("check")).Close()
		b.Div().InnerHTML(itemType.Label).Close()
		b.Div().InnerHTML(itemType.Description).Close()
		b.Close() // Form
		b.Close() // Div
	}
	b.CloseAll()

	return nil
}

// insertMarker adds a nebula-layout-insert to the html.Builder.
func layoutInsert(b *html.Builder, layoutID string, widgetID string, place string, check string, endpoint string) {

	url := makeURL(endpoint, "prop=insert", "itemId="+layoutID, "subItemId="+widgetID, "place="+place, "check="+check)

	b.Div().
		Class("nebula-layout-insert").
		Data("place", place).
		Data("hx-get", url).
		Close()
}
