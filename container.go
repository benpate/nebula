package nebula

import (
	"net/url"

	"github.com/benpate/html"
)

// Container represents a complete package of container
type Container []Item

// NewContainer returns a fully initialized Container object
func NewContainer() Container {
	return make(Container, 0)
}

/*****************************************
 * USER INTERFACE FUNCTIONS
 *****************************************/

// View returns an HTML string containing the VIEW version of the container
func View(library Library, container *Container) string {
	builder := html.New()
	library.View(builder, container, 0)
	return builder.String()
}

// Edit returns an HTML string containing the EDIT version of the container
func Edit(library Library, container *Container, endpoint string) string {
	builder := html.New()
	library.Edit(builder, container, 0, endpoint)
	return builder.String()
}

// Prop returns an editable property form based on the URL params provided.
func Prop(library Library, container *Container, params url.Values, endpoint string) (string, error) {
	builder := html.New()
	err := library.Prop(builder, container, 0, params, endpoint)
	return builder.String(), err
}

/*****************************************
 * READ FUNCTIONS
 *****************************************/

// Len returns the number of items in the container
func (container *Container) Len() int {
	return len(*container)
}

// IsEmpty returns TRUE if the container container is empty.
func (container *Container) IsEmpty() bool {
	return container.Len() == 0
}

// GetItem returns a pointer to the item at the desired index
func (container *Container) GetItem(id int) *Item {

	// Return empty item if out of bounds
	if (id < 0) || (id >= container.Len()) {
		return &Item{}
	}

	// Return a valid item
	return &(*container)[id]
}

func (container *Container) GetParent(id int) (int, *Item) {

	for itemIndex := range *container {
		for refIndex := range (*container)[itemIndex].Refs {
			if (*container)[itemIndex].Refs[refIndex] == id {
				return itemIndex, &(*container)[itemIndex]
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
func (container *Container) NewItem(library Library, itemType string) int {
	item := NewItem(itemType)
	id := container.AddItem(item)
	library.Init(container, id)
	return id
}

// AddItem adds a new item to this container structure, and returns the new item's index
func (container *Container) AddItem(item Item) int {
	newID := len(*container)

	*container = append(*container, item)

	return newID
}

// Compact removes any unused items in the container slice
// and reorganizes references
func (container *Container) Compact() {
	front := 0
	back := len(*container) - 1

	for front < back {

		if (*container)[front].Type != "" {
			front = front + 1
			continue
		}

		if (*container)[back].Type == "" {
			back = back - 1
			continue
		}

		container.move(back, front)
	}

	if (*container)[back].Type != "" {
		back = back - 1
	}

	*container = (*container)[:back]
}

// move physically moves an item from one index to another (overwriting the target location)
// and updates references
func (container *Container) move(from int, to int) {

	(*container)[to] = (*container)[from]
	(*container)[from] = Item{}

	for index := range *container {
		(*container)[index].UpdateReference(from, to)
	}
}
