package nebula

import (
	"github.com/benpate/datatype"
	"github.com/benpate/derp"
	"github.com/benpate/nebula"
)

type ChangeType struct {
	ItemID   int    `json:"itemId"   form:"itemId"`   // ID of the root item that will be added to
	ItemType string `json:"itemType" form:"itemType"` // Type of content item to add
	Check    string `json:"check"    form:"check"`    // Checksum to validation transaction.
}

// Execute performs the ChangeType transaction on the provided content structure
func (txn ChangeType) Execute(library *nebula.Library, container *nebula.Container) (int, error) {

	// Bounds check
	if (txn.ItemID < 0) || (txn.ItemID >= container.Len()) {
		return -1, derp.New(500, "content.transaction.ChangeType", "Index out of bounds", txn)
	}

	if err := (*container)[txn.ItemID].Validate(txn.Check); err != nil {
		return -1, derp.New(derp.CodeForbiddenError, "content.transaction.ChangeType", "Invalid Checksum")
	}

	(*container)[txn.ItemID].Type = txn.ItemType
	(*container)[txn.ItemID].Data = datatype.Map{}

	return txn.ItemID, nil
}
