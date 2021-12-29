package vocabulary

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/benpate/html"
	"github.com/benpate/nebula"
)

const ItemTypeLayout = "LAYOUT"

type Layout struct {
	library nebula.Library
}

// Init adds a child WYSIWYG element
func (w Layout) Init(container *nebula.Container, id int) {
	me := container.GetItem(id)
	wysiwyg := container.NewItem(w.library, ItemTypeWYSIWYG)
	me.AddReference(wysiwyg, 0)
}

// View dsplays the layout and its children.
func (w Layout) View(b *html.Builder, container nebula.Container, id int) {

	item := container.GetItem(id)

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

func (w Layout) Edit(b *html.Builder, container nebula.Container, id int, endpoint string) {

	item := container.GetItem(id)
	idString := strconv.Itoa(id)
	style := item.GetString("style")

	b.Div().
		Class("nebula-layout").
		Data("style", style).
		Data("size", strconv.Itoa(len(item.Refs))).
		Data("id", idString)

	// For layouts with multiple items, add an insertion point that cross-cuts the beginning of the layout
	if id == 0 {
		marker(b, idString, nebula.LayoutPlaceLeft, item.Check)
		marker(b, idString, nebula.LayoutPlaceAbove, item.Check)
	}

	for childIndex, childID := range item.Refs {
		childIDString := strconv.Itoa(childID)
		b.Div().Class("nebula-layout-item")

		if showInsertMarker(container, id, childIndex, nebula.LayoutPlaceAbove) {
			marker(b, childIDString, nebula.LayoutPlaceAbove, item.Check)
		}

		if showInsertMarker(container, id, childIndex, nebula.LayoutPlaceLeft) {
			marker(b, childIDString, nebula.LayoutPlaceLeft, item.Check)
		}

		// Render child item
		w.library.Edit(b, container, childID, endpoint)

		if showInsertMarker(container, id, childIndex, nebula.LayoutPlaceRight) {
			marker(b, childIDString, nebula.LayoutPlaceRight, item.Check)
		}

		if showInsertMarker(container, id, childIndex, nebula.LayoutPlaceBelow) {
			marker(b, childIDString, nebula.LayoutPlaceBelow, item.Check)
		}
		b.Close()
	}

	// For layouts with multiple items, add an insertion point the cross-cuts the end of the layout
	if id == 0 {
		marker(b, idString, nebula.LayoutPlaceRight, item.Check)
		marker(b, idString, nebula.LayoutPlaceBelow, item.Check)
	}

	b.Close()
}

func (w Layout) Prop(container *nebula.Container, id int, endpoint string, params url.Values) (string, error) {

	b := html.New()

	b.H1().InnerHTML("Add Another Section").Close()

	b.Div().Class("table")
	for _, itemType := range ItemTypes() {
		b.Div()
		b.Form("", "").Data("hx-post", endpoint).Data("hx-trigger", "click")
		b.Input("hidden", "type").Value("add-item").Close()
		b.Input("hidden", "itemId").Value(params.Get("itemId")).Close()
		b.Input("hidden", "itemType").Value(itemType.Code)
		b.Input("hidden", "place").Value(params.Get("place")).Close()
		b.Input("hidden", "check").Value(params.Get("check")).Close()
		b.Div().InnerHTML(itemType.Label).Close()
		b.Div().InnerHTML(itemType.Description).Close()
		b.Close() // Form
		b.Close() // Div
	}
	b.Close()

	return b.String(), nil
}

func marker(b *html.Builder, itemID string, place string, check string) {
	b.Div().
		Class("layout-insert").
		Data("hx-get", fmt.Sprintf("/.editor/itemTypes?itemId=%s&place=%s&check=%s", itemID, place, check)).
		Data("itemId", itemID).
		Data("place", place).
		Data("check", check).
		Close()
}

// showInsertMarker returns TRUE if an layout insertion marker should be shown at this location
func showInsertMarker(c nebula.Container, parentID int, childIndex int, place string) bool {

	switch c[parentID].GetString("style") {

	case nebula.LayoutStyleRows:

		switch place {
		case nebula.LayoutPlaceAbove:
			return true
		case nebula.LayoutPlaceBelow:
			return false
		case nebula.LayoutPlaceLeft:
			return true
		case nebula.LayoutPlaceRight:
			return true
		}

	case nebula.LayoutStyleColumns:

		switch place {
		case nebula.LayoutPlaceLeft:
			return true
		case nebula.LayoutPlaceRight:
			return false
		case nebula.LayoutPlaceAbove:
			return true
		case nebula.LayoutPlaceBelow:
			return true
		}
	}

	return false
}
