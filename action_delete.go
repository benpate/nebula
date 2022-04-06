package nebula

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/html"
)

// Delete action removes a single item from a container
type DeleteItem struct {
	ItemID int    `json:"itemId" form:"itemId"`
	Check  string `json:"check"  form:"check"`
}

func (txn DeleteItem) Get(library *Library, container *Container, endpoint string) string {

	b := html.New()

	b.H1().InnerHTML("Remove This Content Item?")
	b.Div()
	b.Form("", "").Data("hx-post", endpoint)
	b.Input("hidden", "action").Value("delete-item")
	b.Input("hidden", "itemId").Value(convert.String(txn.ItemID))
	b.Input("hidden", "check").Value(txn.Check)
	b.Button().Class("warning").Type("submit").InnerHTML("Delete")
	b.Button().Type("button").Script("on click send closeModal").InnerHTML("Cancel")

	return b.String()
}

// Execute removes a single itme from a container
func (txn DeleteItem) Post(library *Library, container *Container) (int, error) {

	// Find parent index and record
	parentID := container.GetParentID(txn.ItemID)

	if container.IsNil(parentID) {
		return -1, derp.NewBadRequestError("nebula.DeleteItem.Post", "Invalid item", txn)
	}

	parent := container.GetItem(parentID)

	if err := parent.Validate(txn.Check); err != nil {
		return -1, derp.Wrap(err, "nebula.DeleteItem.Post", "Invalid checksum")
	}

	// Remove parent's reference to this item
	container.DeleteReference(parentID, txn.ItemID)
	container.NewChecksum(parentID)

	// TODO: If a delete results in an empty layout, then remove the layout
	// TODO: If a delete results in a layout with a single item, then remove the layout and promote the item

	// Recursively delete this item and all of its children
	return 0, deleteItem(container, parentID, txn.ItemID)
}

// DeleteReference removes an item from a parent.
// This separate function is used to make recursive calls efficiently
func deleteItem(container *Container, parentID int, deleteID int) error {

	// TODO: perhaps these could just fail gracefully? Then, we allow cleanup by the "compact" function
	// to handle items that are still bad

	// Bounds check
	if container.IsNil(parentID) {
		return derp.New(500, "nebula.deleteItem", "Parent index out of bounds", container, parentID)
	}

	// Bounds check
	if container.IsNil(deleteID) {
		return derp.New(500, "nebula.deleteItem", "Child index out of bounds", container, deleteID)
	}

	// Remove all children from the content
	if len((*container)[deleteID].Refs) > 0 {
		for _, childID := range (*container)[deleteID].Refs {
			deleteItem(container, deleteID, childID)
		}
	}

	// Remove the deleted item
	(*container)[deleteID] = Item{}

	// Success!
	return nil
}
