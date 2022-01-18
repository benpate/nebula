package nebula

import (
	"strings"

	"github.com/benpate/convert"
	"github.com/benpate/datatype"
)

// Action interface wraps the "Execute" method, which takes some
// non-reversible, non-idempotent action on the container.
type Action interface {
	Get(library *Library, container *Container, endpoint string) string

	// Execute performs an update to the content data.  It returns
	// the ID of the element to be re-rendered on the client, along
	// with an error (if present).  Implementations can use this to
	// selectively re-render portions of the content structure without
	// reloading the entire page.
	Post(library *Library, container *Container) (int, error)

	// TODO: ExecuteMongo() -- generates an efficient update statement for a mongodb collection.
}

func NewAction(in map[string]interface{}) Action {

	data := datatype.Map(in)

	switch data.GetString("action") {

	case "add-item":

		return AddItem{
			ItemID:    data.GetInt("itemId"),
			SubItemID: data.GetInt("subItemId"),
			Place:     data.GetString("place"),
			ItemType:  data.GetString("itemType"),
			Style:     data.GetString("style"),
			Check:     data.GetString("check"),
		}

	case "change-type":

		return ChangeType{
			ItemID:   data.GetInt("itemId"),
			ItemType: data.GetString("itemType"),
			Check:    data.GetString("check"),
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

	case "sort-children":

		return SortChildren{
			ItemID:   data.GetInt("itemId"),
			ChildIDs: convert.SliceOfInt(strings.Split(data.GetString("childIds"), ",")),
			Check:    data.GetString("check"),
		}

	case "update-item":

		return UpdateItem{
			ItemID: data.GetInt("itemId"),
			Data:   extractData(in),
			Check:  data.GetString("check"),
		}

	case "upload-file":

		return UploadFile{
			ItemID: data.GetInt("itemId"),
			File:   data.GetString("file"),
			Check:  data.GetString("check"),
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
