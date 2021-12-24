package content

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/benpate/datatype"
)

// Content represents a complete package of content
type Content []Item

func New() Content {
	return make(Content, 0)
}

func Default() Content {
	return Content{
		{Type: "CONTAINER", Refs: []int{1}, Data: datatype.Map{"style": "ROWS"}},
		{Type: "WYSIWYG", Data: datatype.Map{"html": "Start typing here..."}},
	}
}

// IsEmpty returns TRUE if the content container is empty.
func (content *Content) IsEmpty() bool {
	return len(*content) == 0
}

// AddItem adds a new item to this content structure, and returns the new item's index
func (content *Content) AddItem(item Item) int {
	newID := len(*content)

	*content = append(*content, item)

	return newID
}

// GetItem returns a pointer to the item at the desired index
func (content *Content) GetItem(id int) *Item {

	// Return empty item if out of bounds
	length := len(*content)
	if (id < 0) || (id >= length) {
		return &Item{}
	}

	// Return a valid item
	return &(*content)[id]
}

func (content *Content) GetParent(id int) (int, *Item) {

	for itemIndex := range *content {
		for refIndex := range (*content)[itemIndex].Refs {
			if (*content)[itemIndex].Refs[refIndex] == id {
				return itemIndex, &(*content)[itemIndex]
			}
		}
	}

	return -1, nil
}

// Compact removes any unused items in the content slice
// and reorganizes references
func (content *Content) Compact() {
	front := 0
	back := len(*content) - 1

	for front < back {

		if (*content)[front].Type != "" {
			front = front + 1
			continue
		}

		if (*content)[back].Type == "" {
			back = back - 1
			continue
		}

		content.move(back, front)
	}

	if (*content)[back].Type != "" {
		back = back - 1
	}

	*content = (*content)[:back]
}

// move physically moves an item from one index to another (overwriting the target location)
// and updates references
func (content *Content) move(from int, to int) {

	(*content)[to] = (*content)[from]
	(*content)[from] = Item{}

	for index := range *content {
		(*content)[index].UpdateReference(from, to)
	}
}

// NewChecksum generates a new checksum value to be inserted into a content.Item
func NewChecksum() string {
	seed := time.Now().Unix()
	source := rand.NewSource(seed)
	return strconv.FormatInt(source.Int63(), 36) + strconv.FormatInt(source.Int63(), 36)
}
