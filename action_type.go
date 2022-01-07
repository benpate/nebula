package nebula

import (
	"github.com/benpate/derp"
)

type ChangeType struct {
	ItemID   int    `json:"itemId"   form:"itemId"`   // ID of the root item that will be added to
	ItemType string `json:"itemType" form:"itemType"` // Type of content item to add
	Check    string `json:"check"    form:"check"`    // Checksum to validation transaction.
}

func (txn ChangeType) Get(library *Library, container *Container) string {
	return ""
}

// Execute performs the ChangeType transaction on the provided content structure
func (txn ChangeType) Post(library *Library, container *Container) (int, error) {

	// Validate transaction
	if err := (*container)[txn.ItemID].Validate(txn.Check); err != nil {
		return -1, derp.New(derp.CodeForbiddenError, "content.transaction.ChangeType", "Invalid Checksum")
	}

	(*container)[txn.ItemID].Type = txn.ItemType

	return txn.ItemID, nil
}
