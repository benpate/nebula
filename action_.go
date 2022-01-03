package nebula

import (
	"github.com/benpate/datatype"
)

func NewAction(in map[string]interface{}) Action {

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

	return NilAction(data)
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
