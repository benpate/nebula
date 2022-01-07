package nebula

import (
	"github.com/benpate/derp"
)

type UploadFile struct {
	ItemID int    `json:"itemId" form:"itemId"`
	File   string `json:"file" form:"file"`
	Check  string `json:"hash"   form:"hash"`
}

func (txn UploadFile) Get(library *Library, container *Container) string {
	return ""
}

func (txn UploadFile) Post(library *Library, container *Container) (int, error) {

	// Bounds check
	if container.IsNil(txn.ItemID) {
		return 0, derp.New(500, "content.transaction.UploadFile", "Index out of bounds", txn.ItemID)
	}

	// Find and validate the item
	if err := (*container)[txn.ItemID].Validate(txn.Check); err != nil {
		return 0, derp.Wrap(err, "content.transaction.UploadFile", "Invalid Checksum")
	}

	// Update "file" data.
	(*container)[txn.ItemID].Set("file", txn.File)

	return txn.ItemID, nil
}
