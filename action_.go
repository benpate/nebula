package nebula

import (
	"github.com/benpate/datatype"
	"github.com/benpate/nebula"
)

type Transaction interface {

	// Execute performs an update to the content data.  It returns
	// the ID of the element to be re-rendered on the client, along
	// with an error (if present).  Implementations can use this to
	// selectively re-render portions of the content structure without
	// reloading the entire page.
	Execute(*nebula.Library, *nebula.Container) (int, error)

	// TODO: ExecuteMongo() -- generates an efficient update statement for a mongodb collection.
}

func Parse(in map[string]interface{}) Transaction {

	data := datatype.Map(in)

	switch data.GetString("type") {

	case "change-type":

		return ChangeType{
			ItemID:   data.GetInt("itemId"),
			ItemType: data.GetString("itemType"),
			Check:    data.GetString("check"),
		}

	case "add-item":

		return AddItem{
			ItemID:   data.GetInt("itemId"),
			Place:    data.GetString("place"),
			ItemType: data.GetString("itemType"),
			Style:    data.GetString("style"),
			Check:    data.GetString("check"),
		}

	case "update-item":

		return UpdateItem{
			ItemID: data.GetInt("itemId"),
			Data:   extractData(in),
			Check:  data.GetString("check"),
		}

	case "delete-item":

		return DeleteItem{
			ItemID: data.GetInt("itemId"),
			Check:  data.GetString("check"),
		}

	case "move-item":

		return MoveItem{
			ItemID:      data.GetInt("itemId"),
			NewParentID: data.GetInt("newParentId"),
			Position:    data.GetInt("position"),
			Check:       data.GetString("check"),
		}
	}

	return NilTransaction(data)
}

func extractData(input map[string]interface{}) map[string]interface{} {

	result := make(map[string]interface{})

	for key, value := range input {
		switch key {
		case "type", "hash", "refs":
			continue
		default:
			result[key] = value
		}
	}

	return result
}
