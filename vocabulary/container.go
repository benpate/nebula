package vocabulary

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/benpate/content"
	"github.com/benpate/html"
)

const ItemTypeContainer = "CONTAINER"

type Container struct {
	library content.Library
}

// Init adds a child WYSIWYG element
func (w Container) Init(c *content.Content, id int) {
	me := c.GetItem(id)
	wysiwyg := c.NewItem(w.library, ItemTypeWYSIWYG)
	me.AddReference(wysiwyg, 0)
}

// View dsplays the container and its children.
func (w Container) View(b *html.Builder, c content.Content, id int) {

	item := c.GetItem(id)

	if len(item.Refs) == 0 {
		return
	}

	b.Div().
		Class("container").
		Data("style", item.GetString("style")).
		Data("size", strconv.Itoa(len(item.Refs)))

	for _, id := range item.Refs {
		b.Div().Class("container-item")
		w.library.View(b, c, id)
		b.Close()
	}
}

func (w Container) Edit(b *html.Builder, c content.Content, id int, endpoint string) {

	item := c.GetItem(id)
	idString := strconv.Itoa(id)
	style := item.GetString("style")

	b.Div().
		Class("container").
		Data("style", style).
		Data("size", strconv.Itoa(len(item.Refs))).
		Data("id", idString)

	// For containers with multiple items, add an insertion point that cross-cuts the beginning of the container
	if id == 0 {
		marker(b, idString, content.ContainerPlaceLeft, item.Check)
		marker(b, idString, content.ContainerPlaceAbove, item.Check)
	}

	for childIndex, childID := range item.Refs {
		childIDString := strconv.Itoa(childID)
		b.Div().Class("container-item")

		if showInsertMarker(c, id, childIndex, content.ContainerPlaceAbove) {
			marker(b, childIDString, content.ContainerPlaceAbove, item.Check)
		}

		if showInsertMarker(c, id, childIndex, content.ContainerPlaceLeft) {
			marker(b, childIDString, content.ContainerPlaceLeft, item.Check)
		}

		// Render child item
		w.library.Edit(b, c, childID, endpoint)

		if showInsertMarker(c, id, childIndex, content.ContainerPlaceRight) {
			marker(b, childIDString, content.ContainerPlaceRight, item.Check)
		}

		if showInsertMarker(c, id, childIndex, content.ContainerPlaceBelow) {
			marker(b, childIDString, content.ContainerPlaceBelow, item.Check)
		}
		b.Close()
	}

	// For containers with multiple items, add an insertion point the cross-cuts the end of the container
	if id == 0 {
		marker(b, idString, content.ContainerPlaceRight, item.Check)
		marker(b, idString, content.ContainerPlaceBelow, item.Check)
	}

	b.Close()
}

func (w Container) Prop(c *content.Content, id int, endpoint string, params url.Values) (string, error) {

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
		Class("container-insert").
		Data("hx-get", fmt.Sprintf("/.editor/itemTypes?itemId=%s&place=%s&check=%s", itemID, place, check)).
		Data("itemId", itemID).
		Data("place", place).
		Data("check", check).
		Close()
}

// showInsertMarker returns TRUE if an container insertion marker should be shown at this location
func showInsertMarker(c content.Content, parentID int, childIndex int, place string) bool {

	switch c[parentID].GetString("style") {

	case content.ContainerStyleRows:

		switch place {
		case content.ContainerPlaceAbove:
			return true
		case content.ContainerPlaceBelow:
			return false
		case content.ContainerPlaceLeft:
			return true
		case content.ContainerPlaceRight:
			return true
		}

	case content.ContainerStyleColumns:

		switch place {
		case content.ContainerPlaceLeft:
			return true
		case content.ContainerPlaceRight:
			return false
		case content.ContainerPlaceAbove:
			return true
		case content.ContainerPlaceBelow:
			return true
		}
	}

	return false
}
