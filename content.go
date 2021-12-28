package content

import (
	"net/url"

	"github.com/benpate/html"
)

// Content represents a complete package of content
type Content []Item

// New returns a fully initialized Content object
func New(library Library) Content {
	return make(Content, 0)
}

/*****************************************
 * USER INTERFACE FUNCTIONS
 *****************************************/

// Init initializes an empty container with default content.
func Init(library Library, content Content) {

	if len(content) > 0 {
		return
	}

	content.NewItem(library, "container")
}

// View returns an HTML string containing the VIEW version of the content
func View(library Library, content Content) string {
	builder := html.New()
	library.View(builder, content, 0)
	return builder.String()
}

// Edit returns an HTML string containing the EDIT version of the content
func Edit(library Library, content Content, endpoint string) string {
	builder := html.New()
	library.Edit(builder, content, 0, endpoint)
	return builder.String()
}

// Prop returns an editable property form based on the URL params provided.
func Prop(library Library, content Content, params url.Values, endpoint string) (string, error) {
	builder := html.New()
	err := library.Prop(builder, content, 0, params, endpoint)
	return builder.String(), err
}

/*****************************************
 * READ FUNCTIONS
 *****************************************/

// IsEmpty returns TRUE if the content container is empty.
func (content *Content) IsEmpty() bool {
	return len(*content) == 0
}

// GetItem returns a pointer to the item at the desired index
func (content *Content) GetItem(id int) *Item {

	// Return empty item if out of bounds
	if id >= len(*content) {
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

/*****************************************
 * WRITE FUNCTIONS
 *****************************************/

// NewItem creates a new item of the designated type and initializes it
// with the default Init() method from the corresponding widget library
func (content *Content) NewItem(library Library, itemType string) int {
	item := NewItem(itemType)
	id := content.AddItem(item)
	library.Init(*content, id)
	return id
}

// AddItem adds a new item to this content structure, and returns the new item's index
func (content *Content) AddItem(item Item) int {
	newID := len(*content)

	*content = append(*content, item)

	return newID
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
