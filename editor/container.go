package editor

import (
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
		builder.Div().Script("install containerInsert").Data("itemId", idString).Data("place", content.ContainerPlaceLeft).Data("check", item.Check).Close()
		builder.Div().Script("install containerInsert").Data("itemId", idString).Data("place", content.ContainerPlaceAbove).Data("check", item.Check).Close()
	}

	for childIndex, childID := range item.Refs {
		childIDString := strconv.Itoa(childID)
		builder.Div().Class("container-item")

		if e.showInsertMarker(c, id, childIndex, content.ContainerPlaceAbove) {
			builder.Div().Script("install containerInsert").Data("itemId", childIDString).Data("place", content.ContainerPlaceAbove).Data("check", item.Check).Close()
		}

		if e.showInsertMarker(c, id, childIndex, content.ContainerPlaceLeft) {
			builder.Div().Script("install containerInsert").Data("itemId", childIDString).Data("place", content.ContainerPlaceLeft).Data("check", item.Check).Close()
		}

		// Render child item
		e.subTree(builder, c, childID)

		if e.showInsertMarker(c, id, childIndex, content.ContainerPlaceRight) {
			builder.Div().Script("install containerInsert").Data("itemId", childIDString).Data("place", content.ContainerPlaceRight).Data("check", item.Check).Close()
		}

		if e.showInsertMarker(c, id, childIndex, content.ContainerPlaceBelow) {
			builder.Div().Script("install containerInsert").Data("itemId", childIDString).Data("place", content.ContainerPlaceBelow).Data("check", item.Check).Close()
		}
		builder.Close()
	}

	// For containers with multiple items, add an insertion point the cross-cuts the end of the container
	if id == 0 {
		builder.Div().Script("install containerInsert").Data("itemId", idString).Data("place", content.ContainerPlaceRight).Data("check", item.Check).Close()
		builder.Div().Script("install containerInsert").Data("itemId", idString).Data("place", content.ContainerPlaceBelow).Data("check", item.Check).Close()
	}

	builder.Close()
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
