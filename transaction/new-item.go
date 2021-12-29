package transaction

import (
	"github.com/benpate/derp"
	"github.com/benpate/nebula"
)

const newItemPositionBefore = 0

const newItemPositionAfter = 1

type NewItem struct {
	ItemID   int    `json:"itemId"   form:"itemId"`   // ID of the root item that will be added to
	Place    string `json:"place"    form:"place"`    // ABOVE, BELOW, LEFT, RIGHT
	ItemType string `json:"itemType" form:"itemType"` // Type of content item to add
	Check    string `json:"check"    form:"check"`    // Checksum to validation transaction.
}

// Execute performs the NewItem transaction on the provided content structure
func (txn NewItem) Execute(container *nebula.Container) (int, error) {

	// Bounds check
	if (txn.ItemID < 0) || (txn.ItemID >= container.Len()) {
		return 0, derp.New(500, "content.transaction.NewItem", "Index out of bounds", txn)
	}

	// Find and validate the sibling
	sibling := container.GetItem(txn.ItemID)

	if err := sibling.Validate(txn.Check); err != nil {
		return -1, derp.Wrap(err, "content.transaction.NewItem", "Invalid Transaction")
	}

	// Create a new item to insert into the content
	newItem := nebula.NewItem(txn.ItemType)

	// Insert at head or tail of a container
	if sibling.Type == nebula.ItemTypeLayout {
		switch txn.Place {
		case nebula.LayoutPlaceAbove:
			if sibling.GetString("style") == nebula.LayoutStyleRows {
				addFirstRef(container, txn.ItemID, newItem)
				return txn.ItemID, nil
			}

		case nebula.LayoutPlaceBelow:
			if sibling.GetString("style") == nebula.LayoutStyleRows {
				addLastRef(container, txn.ItemID, newItem)
				return txn.ItemID, nil
			}

		case nebula.LayoutPlaceLeft:
			if sibling.GetString("style") == nebula.LayoutStyleColumns {
				addFirstRef(container, txn.ItemID, newItem)
				return txn.ItemID, nil
			}

		case nebula.LayoutPlaceRight:
			if sibling.GetString("style") == nebula.LayoutStyleColumns {
				addLastRef(container, txn.ItemID, newItem)
				return txn.ItemID, nil
			}
		}
	}

	// Locate the parent
	parentIndex, parent := findParent(container, txn.ItemID)

	if parent.IsEmpty() {
		return -1, derp.New(derp.CodeBadRequestError, "nebula.transaction.NewItem", "Cannot find parent of item", sibling)
	}

	// If the parent is already a container (of the right direction) then
	// we only need to add this new item into it...
	if parent != nil && parent.Type == nebula.ItemTypeLayout {

		switch txn.Place {
		case nebula.LayoutPlaceAbove:
			if parent.GetString("style") == nebula.LayoutStyleRows {
				insertRef(container, parentIndex, txn.ItemID, newItem, newItemPositionBefore)
				return parentIndex, nil
			}

		case nebula.LayoutPlaceBelow:
			if parent.GetString("style") == nebula.LayoutStyleRows {
				insertRef(container, parentIndex, txn.ItemID, newItem, newItemPositionAfter)
				return parentIndex, nil
			}

		case nebula.LayoutPlaceLeft:
			if parent.GetString("style") == nebula.LayoutStyleColumns {
				insertRef(container, parentIndex, txn.ItemID, newItem, newItemPositionBefore)
				return parentIndex, nil
			}

		case nebula.LayoutPlaceRight:
			if parent.GetString("style") == nebula.LayoutStyleColumns {
				insertRef(container, parentIndex, txn.ItemID, newItem, newItemPositionAfter)
				return parentIndex, nil
			}
		}
	}

	// Fall through means that we need to make a new container.
	// ABOVE,BELOW require a ROWS container
	// LEFT,RIGHT require a COLUMNS container

	switch txn.Place {
	case nebula.LayoutPlaceAbove:
		replaceWithLayout(container, nebula.LayoutStyleRows, txn.ItemID, newItem, newItemPositionBefore)
		return txn.ItemID, nil

	case nebula.LayoutPlaceBelow:
		replaceWithLayout(container, nebula.LayoutStyleRows, txn.ItemID, newItem, newItemPositionAfter)
		return txn.ItemID, nil

	case nebula.LayoutPlaceLeft:
		replaceWithLayout(container, nebula.LayoutStyleColumns, txn.ItemID, newItem, newItemPositionBefore)
		return txn.ItemID, nil

	case nebula.LayoutPlaceRight:
		replaceWithLayout(container, nebula.LayoutStyleColumns, txn.ItemID, newItem, newItemPositionAfter)
		return txn.ItemID, nil
	}

	// Something bad happened.  Abort! Abort!
	return 0, derp.New(500, "content.transaction.NewItem", "Invalid transaction", txn)

}

func (txn NewItem) Description() string {
	return "New Item (" + txn.ItemType + ")"
}

func addFirstRef(container *nebula.Container, parentID int, newItem nebula.Item) {
	newID := container.AddItem(newItem)
	oldRefs := (*container)[parentID].Refs
	(*container)[parentID].Refs = append([]int{newID}, oldRefs...)
}

func addLastRef(container *nebula.Container, parentID int, newItem nebula.Item) {
	newID := container.AddItem(newItem)
	(*container)[parentID].Refs = append((*container)[parentID].Refs, newID)
}

// insertRef inserts `newItem` into the content, and places a reference to it inside of the
// `parentID` item, either BEFORE or AFTER the `childID` item
func insertRef(container *nebula.Container, parentID int, childID int, newItem nebula.Item, position int) {
	newID := container.AddItem(newItem)
	newRefs := make([]int, 0)

	for _, itemID := range (*container)[parentID].Refs {
		if itemID == childID {
			if position == newItemPositionBefore {
				newRefs = append(newRefs, newID, itemID)
			} else {
				newRefs = append(newRefs, itemID, newID)
			}
		} else {
			newRefs = append(newRefs, itemID)
		}
	}

	(*container)[parentID].Refs = newRefs
}

// replaceWithContainer replaces an existing content Item with a container (of a specific style),
// moves the original item to the end of the content structure,
// then inserts the new item into correct position of the new container (either BEFORE or AFTER the original Item)
func replaceWithLayout(container *nebula.Container, style string, itemID int, newItem nebula.Item, position int) {

	// reset the checksum on the current item
	(*container)[itemID].NewChecksum()

	// copy the current item to the end of the content structure
	newLocationID := container.Len()
	*container = append(*container, (*container)[itemID])

	// insert a new layout in the spot where the original content was
	layout := nebula.NewItem(nebula.ItemTypeLayout, newLocationID)
	layout.Set("style", style)
	(*container)[itemID] = layout

	// insert a reference to the newItem into the new container
	insertRef(container, itemID, newLocationID, newItem, position)
}
