package transaction

import (
	"github.com/benpate/nebula"
)

// findParent returns the item that contains the provided itemID.  If no
// container is found, then (-1, nil) is returned
func findParent(container *nebula.Container, itemID int) (int, *nebula.Item) {

	for itemIndex := range *container {
		if refIndex := findChildPosition(container, itemIndex, itemID); refIndex != -1 {
			return itemIndex, &(*container)[itemIndex]
		}
	}

	return -1, nil
}

// findChildPosition returns the position of a childID in the ref array.
func findChildPosition(container *nebula.Container, itemID int, childID int) int {
	for refIndex := range (*container)[itemID].Refs {
		if (*container)[itemID].Refs[refIndex] == childID {
			return refIndex
		}
	}

	return -1
}
