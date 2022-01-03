package nebula

import (
	"github.com/benpate/derp"
)

type MoveItem struct {
	ItemID      int    `json:"itemId"      form:"itemId"`
	NewParentID int    `json:"newParentId" form:"newParentId"`
	Position    int    `json:"position"    form:"position"`
	Check       string `json:"check"       form:"check"`
}

func (txn MoveItem) Execute(library *Library, container *Container) (int, error) {
	return -1, derp.New(derp.CodeBadRequestError, "content.transaction.MoveItem", "Unimplemented")
}