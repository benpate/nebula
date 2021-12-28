package transaction

import (
	"github.com/benpate/content"
	"github.com/benpate/derp"
)

type MoveItem struct {
	ItemID      int    `json:"itemId"      form:"itemId"`
	NewParentID int    `json:"newParentId" form:"newParentId"`
	Position    int    `json:"position"    form:"position"`
	Check       string `json:"check"       form:"check"`
}

func (txn MoveItem) Execute(c *content.Content) (int, error) {
	return 0, derp.New(derp.CodeBadRequestError, "content.transaction.MoveItem", "Unimplemented")
}

func (txn MoveItem) Description() string {
	return "Move Item"
}
