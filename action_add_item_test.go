package nebula

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/stretchr/testify/require"
)

func TestItem(t *testing.T) {

	library := NewLibrary()

	container := NewContainer()

	firstID := container.NewItemWithInit(&library, ItemTypeHTML, datatype.Map{"html": "FIRST HTML ITEM"})

	itemID, err := container.Execute(&library, datatype.Map{
		"type":     "add-item",
		"itemId":   2,
		"itemType": "WYSIWYG",
		"place":    "RIGHT",
		"check":    container.GetChecksum(firstID),
	})

	require.Nil(t, err)
	require.Equal(t, 0, itemID)
}
