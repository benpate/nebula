package nebula

import (
	"github.com/benpate/derp"
)

// Delete action removes a single item from a container
type DeleteItem struct {
	ItemID int    `json:"itemId" form:"itemId"`
	Check  string `json:"check"  form:"check"`
}

// Execute removes a single itme from a container
func (txn DeleteItem) Execute(library *Library, container *Container) (int, error) {

	// Find parent index and record
	parentID := container.GetParentID(txn.ItemID)

	if container.IsNil(parentID) {
		return -1, derp.New(derp.CodeBadRequestError, "nebula.DeleteItem.Execute", "Invalid item", txn)
	}

	// Remove parent's reference to this item
	container.DeleteReference(parentID, txn.ItemID)

	// TODO: If a delete results in an empty layout, then remove the layout
	// TODO: If a delete results in a layout with a single item, then remove the layout and promote the item

	// Recursively delete this item and all of its children
	return parentID, deleteItem(container, parentID, txn.ItemID, txn.Check)
}

// DeleteReference removes an item from a parent.
// This separate function is used to make recursive calls efficiently
func deleteItem(container *Container, parentID int, deleteID int, check string) error {

	// TODO: perhaps these could just fail gracefully? Then, we allow cleanup by the "compact" function
	// to handle items that are still bad

	// Bounds check
	if container.IsNil(parentID) {
		return derp.New(500, "content.Create", "Parent index out of bounds", parentID, deleteID)
	}

	// Bounds check
	if container.IsNil(deleteID) {
		return derp.New(500, "content.Create", "Child index out of bounds", parentID, deleteID)
	}

	// validate checksum
	if err := (*container)[parentID].Validate(check); err != nil {
		return derp.Wrap(err, "content.Create", "Invalid Checksum")
	}

	// Remove all children from the content
	if len((*container)[deleteID].Refs) > 0 {
		childCheck := (*container)[deleteID].Check
		for _, childID := range (*container)[deleteID].Refs {
			deleteItem(container, deleteID, childID, childCheck)
		}
	}

	// Remove the deleted item
	(*container)[deleteID] = Item{}

	// Success!
	return nil
}
