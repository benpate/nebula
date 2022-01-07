package nebula

import (
	"net/url"
	"strconv"

	"github.com/benpate/convert"
	"github.com/benpate/html"
)

// ItemTypeLayout represents a Layout widget
const ItemTypeLayout = "LAYOUT"

/************************
 *  Layout Styles
 ************************/

// LayoutStyleRows represents a group of content items organized into rows
const LayoutStyleRows = "ROWS"

// LayoutStyle Colums represents a group of content items organized into responsive columns
const LayoutStyleColumns = "COLS"

/************************
 *  Layout Placements
 ************************/

// LayoutPlaceBefore represents the placement of a new item before the reference item
const LayoutPlaceBefore = "BEFORE"

// LayoutPlaceAfter represents the placement of a new item after the reference item
const LayoutPlaceAfter = "AFTER"

// LayoutPlaceAbove represents the placement of a new item above the reference item
const LayoutPlaceAbove = "ABOVE"

// LayoutPlaceAbove represents the placement of a new item below the reference item
const LayoutPlaceBelow = "BELOW"

// LayoutPlaceAbove represents the placement of a new item to the left of the reference item
const LayoutPlaceLeft = "LEFT"

// LayoutPlaceAbove represents the placement of a new item to the right of the reference item
const LayoutPlaceRight = "RIGHT"

type Layout struct {
	library *Library
}

// Init adds a child WYSIWYG element
func (w Layout) Init(container *Container, id int) {
	wysiwyg := container.NewItemWithInit(w.library, ItemTypeWYSIWYG, nil)
	(*container)[id].AddReference(wysiwyg, 0)
}

// View dsplays the layout and its children.
func (w Layout) View(b *html.Builder, container *Container, layoutID int) {

	item := container.GetItem(layoutID)

	if len(item.Refs) == 0 {
		return
	}

	b.Div().
		Class("nebula-layout").
		Data("style", item.GetString("style")).
		Data("size", strconv.Itoa(len(item.Refs)))

	for _, id := range item.Refs {
		b.Div().Class("nebula-layout-item")
		w.library.View(b, container, id)
		b.Close()
	}
}

func (w Layout) Edit(b *html.Builder, container *Container, layoutID int, endpoint string) {

	layout := container.GetItem(layoutID)
	layoutIDString := strconv.Itoa(layoutID)

	b.Div().
		Class("nebula-layout").
		Data("style", "ROWS").
		Data("size", strconv.Itoa(len(layout.Refs))).
		Data("id", layoutIDString)

	for childIndex, childID := range layout.Refs {
		childIDString := strconv.Itoa(childID)

		child := container.GetItem(childID)
		b.Div().Class("nebula-layout-item")

		if childIndex == 0 {
			layoutInsert(b, layoutIDString, childIDString, LayoutPlaceAbove, child.Check, endpoint)
		}

		if child.Type != ItemTypeLayout {
			b.Div().Class("nebula-layout-controls")
			b.Div().Class("nebula-layout-sortable-handle")
			b.Container("i").Class("fa-solid fa-grip-horizontal").Close()
			b.Close()

			b.Div().Class("nebula-layout-delete")
			b.Container("i").Class("fa-solid fa-circle-xmark").Close()
			b.Close()
			b.Close()
		}

		w.library.Edit(b, container, childID, endpoint)

		layoutInsert(b, layoutIDString, childIDString, LayoutPlaceBelow, child.Check, endpoint)

		b.Close()
	}

	b.Close()
}

func (w Layout) Prop(b *html.Builder, container *Container, id int, endpoint string, params url.Values) error {

	b.H1().InnerHTML("Add Another Section").Close()

	b.Div().Class("table")
	for _, itemType := range ItemTypes() {
		b.Div().Attr("tabindex", "0")
		b.Form("", "").Data("hx-post", endpoint).Data("hx-trigger", "click")
		b.Input("hidden", "type").Value("add-item").Close()
		b.Input("hidden", "itemId").Value(params.Get("subItemId")).Close()
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

	b.Span().
		Class("nebula-layout-insert").
		Data("place", place).
		Data("hx-get", url).
		Close()
}
