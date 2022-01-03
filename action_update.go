package nebula

import (
	"github.com/benpate/derp"
)

type UpdateItem struct {
	ItemID int                    `json:"itemId" form:"itemId"`
	Data   map[string]interface{} `json:"data"   form:"data"`
	Check  string                 `json:"hash"   form:"hash"`
}

func (txn UpdateItem) Execute(library *Library, container *Container) (int, error) {

	// Bounds check
	if (txn.ItemID < 0) || (txn.ItemID >= container.Len()) {
		return 0, derp.New(500, "content.transaction.UpdateItem", "Index out of bounds", txn.ItemID)
	}

	// Find and validate the item
	if err := (*container)[txn.ItemID].Validate(txn.Check); err != nil {
		return 0, derp.Wrap(err, "content.transaction.UpdateItem", "Invalid Checksum")
	}

	// Update data
	for key, value := range txn.Data {
		(*container)[txn.ItemID].Set(key, value)
	}

	return txn.ItemID, nil
}
