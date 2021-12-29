package transaction

import (
	"github.com/benpate/derp"
	"github.com/benpate/nebula"
)

type MoveItem struct {
	ItemID      int    `json:"itemId"      form:"itemId"`
	NewParentID int    `json:"newParentId" form:"newParentId"`
	Position    int    `json:"position"    form:"position"`
	Check       string `json:"check"       form:"check"`
}

func (txn MoveItem) Execute(container *nebula.Container) (int, error) {
	return -1, derp.New(derp.CodeBadRequestError, "content.transaction.MoveItem", "Unimplemented")
}

func (txn MoveItem) Description() string {
	return "Move Item Unimplemented"
}
