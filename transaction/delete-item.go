package transaction

import (
	"github.com/benpate/derp"
	"github.com/benpate/nebula"
)

type DeleteItem struct {
	ItemID int    `json:"itemId" form:"itemId"`
	Check  string `json:"check"  form:"check"`
}

func (txn DeleteItem) Execute(container *nebula.Container) (int, error) {

	// Find parent index and record
	parentID := container.GetParent(txn.ItemID)

	// Remove parent's reference to this item
	(*container)[parentID].DeleteReference(txn.ItemID)

	// Recursively delete this item and all of its children
	return parentID, deleteItem(container, parentID, txn.ItemID, txn.Check)
}

func (txn DeleteItem) Description() string {
	return "Delete Item"
}

// DeleteReference removes an item from a parent
func deleteItem(container *nebula.Container, parentID int, deleteID int, check string) error {

	// Bounds check
	if (parentID < 0) || (parentID >= container.Len()) {
		return derp.New(500, "content.Create", "Parent index out of bounds", parentID, deleteID)
	}

	// Bounds check
	if (deleteID < 0) || (deleteID >= container.Len()) {
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
	(*container)[deleteID] = nebula.Item{}

	// Success!
	return nil
}
