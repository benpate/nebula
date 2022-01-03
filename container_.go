package nebula

import (
	"github.com/benpate/datatype"
)

// Container represents a complete package of container
type Container []Item

// NewContainer returns a fully initialized Container object
func NewContainer() Container {
	return make(Container, 0)
}

/*****************************************
 * TODO: PATH INTERFACE
 *****************************************/

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

// GetItem returns a copy of the item at the desired index
func (container *Container) GetItem(itemID int) Item {

	// Return empty item if out of bounds
	if (itemID < 0) || (itemID >= container.Len()) {
		return Item{}
	}

	// Return a valid item
	return (*container)[itemID]
}

// GetType returns the item type of the designated item
func (container *Container) GetType(itemID int) string {
	return container.GetItem(itemID).Type
}

// GetRefs returns the child references of the designated item
func (container *Container) GetRefs(itemID int) []int {
	return container.GetItem(itemID).Refs
}

// GetChecksum returns the checksum of the designated item
func (container *Container) GetChecksum(itemID int) string {
	return container.GetItem(itemID).Check
}

// GetParent searches for the ID of a parent item
func (container *Container) GetParent(itemID int) int {

	for itemIndex := range *container {
		for refIndex := range (*container)[itemIndex].Refs {
			if (*container)[itemIndex].Refs[refIndex] == itemID {
				return itemIndex
			}
		}
	}

	return -1
}

/*****************************************
 * WRITE FUNCTIONS
 *****************************************/

// NewItem creates a new item of the designated type
func (container *Container) NewItem(itemType string, data datatype.Map) int {

	if data == nil {
		data = make(datatype.Map)
	}

	// Create a new item using type and data
	itemID := len(*container)
	item := NewItem(itemType)
	item.Data = data

	// Append to the container
	*container = append(*container, item)

	// Done
	return itemID
}

// NewItem creates a new item of the designated type and initializes it
// with the default Init() method from the corresponding widget library
func (container *Container) NewItemWithInit(library *Library, itemType string, data datatype.Map) int {

	// Create the new item
	itemID := container.NewItem(itemType, data)

	// Initialize the item (adding any extra dependencies, like tab containers)
	library.Init(container, itemID)

	// Is this success?
	return itemID
}

func (container *Container) AddFirstReference(itemID int, newItemID int) {
	(*container)[itemID].AddFirstReference(newItemID)
}

func (container *Container) AddLastReference(itemID int, newItemID int) {
	(*container)[itemID].AddLastReference(newItemID)
}

// AddReference links the newItemID into the parent's reference list, placed relative to the referenceID
func (container *Container) AddReference(parentID int, newItemID int, referenceID int, place string) {

	index := (*container)[parentID].findReference(referenceID)

	// If the item cannot be found, then place relative to the whole reference list.
	// With a good UI, this shouldn't happen very much, though.
	if index == -1 {
		if place == LayoutPlaceBefore {
			container.AddLastReference(parentID, newItemID)
		} else {
			container.AddLastReference(parentID, newItemID)
		}
		return
	}

	// To insert AFTER an indexed item, increment the index
	if place == LayoutPlaceAfter {
		index = index + 1
	}

	(*container)[parentID].AddReference(newItemID, index)
}

// ReplaceRefs searches all references from parentID, and replaces oldID with newID if found
func (container *Container) ReplaceReference(parentID int, oldID int, newID int) {

	// Bounds check
	if (parentID < 0) || (parentID >= len(*container)) {
		return
	}

	// scan parent record for references to oldID.  Replace if found
	for index, itemID := range (*container)[parentID].Refs {
		if itemID == oldID {
			(*container)[parentID].Refs[index] = newID
			return
		}
	}
}

// Execute parses and executes a new Action against this container.
func (container *Container) Execute(library *Library, input map[string]interface{}) (int, error) {

	action := NewAction(input)
	return action.Execute(library, container)
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
