package nebula

import (
	"strconv"

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
		Data("id", layoutIDString).
		Data("check", layout.Check)

	for childIndex, childID := range layout.Refs {
		childIDString := strconv.Itoa(childID)

		child := container.GetItem(childID)
		b.Div().Class("nebula-layout-item").Data("id", childIDString)

		if childIndex == 0 {
			layoutInsert(b, layoutIDString, childIDString, LayoutPlaceBefore, layout.Check, endpoint)
		}

		if child.Type != ItemTypeLayout {

			deleteURL := makeURL(endpoint, "action=delete-item", "itemId="+childIDString, "check="+layout.Check)

			b.Div().Class("nebula-layout-controls")
			b.Div().Class("nebula-layout-sortable-handle")
			b.Container("i").Class("ti ti-grip-horizontal").Close()
			b.Close()

			b.Div().Class("nebula-layout-delete").Data("hx-get", deleteURL).Data("hx-vals", "{}")
			b.Container("i").Class("ti ti-circle-x").Close()
			b.Close()
			b.Close()
		}

		w.library.Edit(b, container, childID, endpoint)

		layoutInsert(b, layoutIDString, childIDString, LayoutPlaceAfter, layout.Check, endpoint)

		b.Close()
	}

	b.Close()
}

// Validate cleans the container for invalid content
func (w Layout) Validate(container *Container, index int) {
}

// insertMarker adds a nebula-layout-insert to the html.Builder.
func layoutInsert(b *html.Builder, layoutID string, childID string, place string, check string, endpoint string) {

	url := makeURL(endpoint, "action=add-item", "itemId="+layoutID, "subItemId="+childID, "place="+place, "check="+check)

	b.Span().
		Class("nebula-layout-insert").
		Data("place", place).
		Data("hx-get", url).
		Close()
}
