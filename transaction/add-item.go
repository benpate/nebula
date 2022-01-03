package transaction

import (
	"github.com/benpate/datatype"
	"github.com/benpate/derp"
	"github.com/benpate/first"
	"github.com/benpate/nebula"
)

type AddItem struct {
	ItemID   int    `json:"itemId"   form:"itemId"`   // ID of the root item that will be added to
	Place    string `json:"place"    form:"place"`    // Position of the new element (BEFORE, AFTER, ABOVE, BELOW, LEFT, RIGHT) relative to the index
	ItemType string `json:"itemType" form:"itemType"` // Type of content item to add
	Style    string `json:"style"    form:"style"`    // Optional "style" aregument for certain types (like layouts)
	Check    string `json:"check"    form:"check"`    // Checksum to validation transaction.
}

// Execute performs the AddItem transaction on the provided content structure
func (txn AddItem) Execute(library *nebula.Library, container *nebula.Container) (int, error) {

	/*** Validate the transa tion */

	// Try to get a copy of the item to be modified (this will return a Nil item, if not found)
	item := container.GetItem(txn.ItemID)

	// Validate the item can be manipulated
	if err := item.Validate(txn.Check); err != nil {
		return -1, derp.Wrap(err, "nebula.transaction.AddItem", "Invalid item", txn)
	}

	// Create the new item to insert into the container
	newItemID := container.NewItemWithInit(library, txn.ItemType, nil)

	/*** Try to append directly to the layout, if possible */

	// If we can append to this item, do it
	switch canAppendLayout(&item, txn.Place) {

	case nebula.LayoutPlaceBefore:
		container.AddFirstReference(txn.ItemID, newItemID)
		return txn.ItemID, nil

	case nebula.LayoutPlaceAfter:
		container.AddLastReference(txn.ItemID, newItemID)
		return txn.ItemID, nil
	}

	/*** If we can't append to the item directly, then try to insert into to its parent instead */

	parentID := findParent(container, txn.ItemID)
	parent := container.GetItem(parentID)

	// If we can append to this layout, do it
	if place := canAppendLayout(&parent, txn.Place); place != "" {
		container.AddReference(parentID, newItemID, txn.ItemID, place)
		return parentID, nil
	}

	/*** Fall through means that we'll need to split/replace the existing item with a layout */

	// create a new layout item that will contain the new items
	newLayoutID := container.NewItem(nebula.ItemTypeLayout, datatype.Map{
		"style": getLayoutStyleFromPlace(txn.Place),
	})
	newLayout := container.GetItem(newLayoutID)

	// Swap original parent with new layout
	(*container)[txn.ItemID] = newLayout
	(*container)[newLayoutID] = item

	// Wrap the original item in the new layout
	container.AddFirstReference(txn.ItemID, newLayoutID)

	place := canAppendLayout(&newLayout, txn.Place)

	// Add the new item into the new layout
	container.AddReference(txn.ItemID, newItemID, newLayoutID, place)

	// Since we have moved things around, replace the whole parent
	return parentID, nil
}

// findParent returns the ID of the designated item's parent.
// If the item has no parent, then this function returns -1.
func findParent(container *nebula.Container, itemID int) int {

	for itemIndex := range *container {
		for _, refID := range (*container)[itemIndex].Refs {
			if refID == itemID {
				return itemIndex
			}
		}
	}

	return -1
}

// canAppendLayout returns TRUE if the placement matches the item's layout style
func canAppendLayout(item *nebula.Item, place string) string {

	if item.Type != nebula.ItemTypeLayout {
		return ""
	}

	switch getLayoutStyle(item) {
	case nebula.LayoutStyleRows:

		switch place {

		case nebula.LayoutPlaceBefore, nebula.LayoutPlaceAbove:
			return nebula.LayoutPlaceBefore

		case nebula.LayoutPlaceAfter, nebula.LayoutPlaceBelow:
			return nebula.LayoutPlaceAfter
		}

	case nebula.LayoutStyleColumns:

		switch place {

		case nebula.LayoutPlaceBefore, nebula.LayoutPlaceLeft:
			return nebula.LayoutPlaceBefore

		case nebula.LayoutPlaceAfter, nebula.LayoutPlaceRight:
			return nebula.LayoutPlaceAfter
		}
	}

	return ""
}

// getLayoutStyleFromPlace returns the correct layout style to use
// when placing new items relative to an existing one.
func getLayoutStyleFromPlace(place string) string {

	switch place {

	case nebula.LayoutPlaceAbove:
		return nebula.LayoutStyleRows

	case nebula.LayoutPlaceAfter:
		return nebula.LayoutStyleRows

	case nebula.LayoutPlaceBefore:
		return nebula.LayoutStyleRows

	case nebula.LayoutPlaceBelow:
		return nebula.LayoutStyleRows

	case nebula.LayoutPlaceLeft:
		return nebula.LayoutStyleColumns

	case nebula.LayoutPlaceRight:
		return nebula.LayoutStyleColumns

	}

	return nebula.LayoutStyleRows
}

// getLayoutStyle returns a valid layout style for all layout items.
// If a non-layout item is passed, then it returns ""
func getLayoutStyle(item *nebula.Item) string {

	if item.Type == nebula.ItemTypeLayout {
		return first.String(item.GetString("style"), nebula.LayoutStyleRows)
	}

	return ""
}
