package nebula

import (
	"github.com/benpate/derp"
	"github.com/benpate/rosetta/compare"
)

type SortChildren struct {
	ItemID   int    `json:"itemId"   form:"itemId"`   // ID of the layout that will hold the new item
	ChildIDs []int  `json:"childIds" form:"childIds"` // Sorted list of IDs to use for child references
	Check    string `json:"check"    form:"check"`    // Checksum to validation transaction.
}

func (txn SortChildren) Get(library *Library, container *Container, endpoint string) string {
	return ""
}

// Execute performs the SortChildren transaction on the provided content structure
func (txn SortChildren) Post(library *Library, container *Container) (int, error) {

	/*** If we can't append to the item directly, then try to insert into to its parent instead */

	parent := container.GetItem(txn.ItemID)

	// Validate the request
	if err := parent.Validate(txn.Check); err != nil {
		return -1, derp.Wrap(err, "nebula.SortChildren.Post", "Invalid checksum", txn)
	}

	// Validate that lengths are equal
	if len(txn.ChildIDs) != len(parent.Refs) {
		return -1, derp.NewBadRequestError("nebula.SortChildren.Post", "Invalid sort list", txn.ChildIDs)
	}

	// Validate that all items are the same (just sorted differently)
	for _, childID := range txn.ChildIDs {
		if !compare.Contains(parent.Refs, childID) {
			return -1, derp.NewBadRequestError("nebula.SortChildren.Post", "Invalid sort list.  Value does not appear in existing references", childID)
		}
	}

	// Validation passed, set the value
	(*container)[txn.ItemID].Refs = txn.ChildIDs
	// container.NewChecksum(txn.ItemID)

	return txn.ItemID, nil
}
