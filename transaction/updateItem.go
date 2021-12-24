package transaction

import (
	"github.com/benpate/content"
	"github.com/benpate/derp"
	"github.com/davecgh/go-spew/spew"
)

type UpdateItem struct {
	ItemID int                    `json:"itemId" form:"itemId"`
	Data   map[string]interface{} `json:"data"   form:"data"`
	Check  string                 `json:"hash"   form:"hash"`
}

func (txn UpdateItem) Execute(c *content.Content) error {

	spew.Dump(txn, c)
	// Bounds check
	if (txn.ItemID < 0) || (txn.ItemID >= len(*c)) {
		return derp.New(500, "content.transaction.UpdateItem", "Index out of bounds", txn.ItemID)
	}

	// Validate checksum
	if txn.Check != (*c)[txn.ItemID].Check {
		return derp.New(derp.CodeForbiddenError, "content.transaction.UpdateItem", "Invalid Checksum")
	}

	// Update data
	(*c)[txn.ItemID].Data = txn.Data
	return nil
}

func (txn UpdateItem) Description() string {
	return "Update Item"
}
