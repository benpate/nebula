package nebula

import (
	"github.com/benpate/convert"
	"github.com/benpate/datatype"
	"github.com/benpate/derp"
	"github.com/benpate/path"
)

// Item represents a single piece of content.  It will be rendered by one of several rendering
// Libraries, using the custom data it contains.
type Item struct {
	Type  string       `json:"type"           bson:"type"`  // The type of contem item (WYSIWYG, CONTAINER, OEMBED, ETC...)
	Check string       `json:"check"          bson:"check"` // A random code or nonce to authenticate requests
	Refs  []int        `json:"refs,omitempty" bson:"refs"`  // Indexes of sub-items contained by this item
	Data  datatype.Map `json:"data,omitempty" bson:"data"`  // Additional data specific to this item type.
}

// NewItem returns a fully initialized Item
func NewItem(t string, refs ...int) Item {
	return Item{
		Type:  t,
		Data:  make(datatype.Map),
		Refs:  refs,
		Check: NewChecksum(),
	}
}

// IsEmpty returns TRUE if this item does not have a valid item.Type
func (item *Item) IsEmpty() bool {
	return (item == nil) || (item.Type == "")
}

// Validate checks that an item has a type (meaning it has been found by container.GetItem)
// and has a valid checksum
func (item *Item) Validate(checksum string) error {

	if item.IsEmpty() {
		return derp.New(derp.CodeBadRequestError, "nebula.Item.Validate", "Item is empty")
	}

	if item.Check != checksum {
		return derp.New(derp.CodeForbiddenError, "nebula.Item.Validate", "Invalid checksum", checksum)
	}

	return nil
}

// AddFirstReference adds an itemID to the beginning of the reference list
func (item *Item) AddFirstReference(id int) {
	item.Refs = append([]int{id}, item.Refs...)
}

// AddLastReference adds and itemID to the end of the reference list
func (item *Item) AddLastReference(id int) {
	item.Refs = append(item.Refs, id)
}

// AddReference adds a new item into the middle of the reference list
func (item *Item) AddReference(id int, index int) {

	// special case for empty refs.  No need to do all that work.
	if len(item.Refs) == 0 {
		item.Refs = []int{id}
		return
	}

	// efficient insert for already-populated refs.
	item.Refs = append(item.Refs, 0)
	copy(item.Refs[index+1:], item.Refs[index:])
	item.Refs[index] = id
}

// UpdateReference migrates references from an old value to a new one
func (item *Item) UpdateReference(from int, to int) {
	for index := range item.Refs {
		if item.Refs[index] == from {
			item.Refs[index] = to
			return
		}
	}
}

// DeleteReference removes a reference from this Item.
func (item *Item) DeleteReference(id int) {
	for index := range item.Refs {
		if item.Refs[index] == id {
			item.Refs = append(item.Refs[:index], item.Refs[index+1:]...)
			return
		}
	}
}

// UnmarshalMap extracts data from a map[string]interface{} to populate this Item
func (item *Item) UnmarshalMap(value map[string]interface{}) {
	item.Type = convert.String(value["type"])
	item.Refs = convert.SliceOfInt(value["refs"])
	item.Data = convert.MapOfInterface("data")
	item.Check = NewChecksum()
}

/*****************************************
 * Data Accessors
 *****************************************/

func (item *Item) GetPath(p path.Path) (interface{}, error) {
	return item.Data.GetPath(p)
}

func (item *Item) GetString(key string) string {
	return item.Data.GetString(key)
}

func (item *Item) GetInt(key string) int {
	return item.Data.GetInt(key)
}

func (item *Item) GetSliceOfInt(key string) []int {
	return item.Data.GetSliceOfInt(key)
}

func (item *Item) GetSliceOfString(key string) []string {
	return item.Data.GetSliceOfString(key)
}

func (item *Item) GetInterface(key string) interface{} {
	return item.Data.GetInterface(key)
}

func (item *Item) SetPath(p path.Path, value interface{}) error {
	return item.Data.SetPath(p, value)
}

func (item *Item) Set(key string, value interface{}) *Item {
	item.Data[key] = value
	return item
}

/*****************************************
 * Other Utilities
 *****************************************/

// findReference searches for another itemID in the item's reference list, and returns it place.
// if the item does not exist in this list, then -1 is returned instead
func (item *Item) findReference(itemID int) int {

	for index, refID := range item.Refs {
		if refID == itemID {
			return index
		}
	}

	return -1
}
