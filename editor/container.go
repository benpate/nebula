package editor

import (
	"fmt"
	"strconv"

	"github.com/benpate/content"
	"github.com/benpate/html"
)

func (e Editor) Container(builder *html.Builder, c content.Content, id int) {

	item := c.GetItem(id)
	idString := strconv.Itoa(id)
	style := item.GetString("style")

	builder.Div().
		Class("container").
		Data("style", style).
		Data("size", strconv.Itoa(len(item.Refs))).
		Data("id", idString)

	// For containers with multiple items, add an insertion point that cross-cuts the beginning of the container
	if id == 0 {
		e.marker(builder, idString, content.ContainerPlaceLeft, item.Check)
		e.marker(builder, idString, content.ContainerPlaceAbove, item.Check)
	}

	for childIndex, childID := range item.Refs {
		childIDString := strconv.Itoa(childID)
		builder.Div().Class("container-item")

		if e.showInsertMarker(c, id, childIndex, content.ContainerPlaceAbove) {
			e.marker(builder, childIDString, content.ContainerPlaceAbove, item.Check)
		}

		if e.showInsertMarker(c, id, childIndex, content.ContainerPlaceLeft) {
			e.marker(builder, childIDString, content.ContainerPlaceLeft, item.Check)
		}

		// Render child item
		e.subTree(builder, c, childID)

		if e.showInsertMarker(c, id, childIndex, content.ContainerPlaceRight) {
			e.marker(builder, childIDString, content.ContainerPlaceRight, item.Check)
		}

		if e.showInsertMarker(c, id, childIndex, content.ContainerPlaceBelow) {
			e.marker(builder, childIDString, content.ContainerPlaceBelow, item.Check)
		}
		builder.Close()
	}

	// For containers with multiple items, add an insertion point the cross-cuts the end of the container
	if id == 0 {
		e.marker(builder, idString, content.ContainerPlaceRight, item.Check)
		e.marker(builder, idString, content.ContainerPlaceBelow, item.Check)
	}

	builder.Close()
}

func (e Editor) marker(b *html.Builder, itemID string, place string, check string) {
	b.Div().
		Class("container-insert").
		Data("hx-get", fmt.Sprintf("/.editor/itemTypes?itemId=%s&place=%s&check=%s", itemID, place, check)).
		Data("itemId", itemID).
		Data("place", place).
		Data("check", check).
		Close()
}

// showInsertMarker returns TRUE if an container insertion marker should be shown at this location
func (e Editor) showInsertMarker(c content.Content, parentID int, childIndex int, place string) bool {

	switch c[parentID].GetString("style") {

	case content.ContainerStyleRows:

		switch place {
		case content.ContainerPlaceAbove:
			return childIndex > 0
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
			return childIndex > 0
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
