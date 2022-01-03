package nebula

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {

	library := NewLibrary()
	container := getTestContainer()

	itemID, err := container.Execute(&library, datatype.Map{
		"type":   "delete-item",
		"itemId": "3",
		"check":  container.GetChecksum(3),
	})

	require.Nil(t, err)
	require.Equal(t, 4, container.Len()) // container still has 4 items, because the removed one is just empty
	require.Equal(t, 0, itemID)

	{
		item := container.GetItem(0)
		require.Equal(t, []int{1, 2}, item.Refs)
	}
	{
		item := container.GetItem(3)
		require.True(t, item.IsEmpty())
	}

	container.Compact()

	require.Equal(t, 3, container.Len()) // empty item removed after compact
}

func getTestContainer() Container {

	container := NewContainer()
	zero := container.NewItem("LAYOUT", datatype.Map{"style": "ROWS"})
	first := container.NewItem("HTML", datatype.Map{"html": "FIRST HTML ITEM"})
	second := container.NewItem("TEXT", datatype.Map{"html": "SECOND TEXT ITEM"})
	third := container.NewItem("WYSIWYG", datatype.Map{"html": "THIRD WYSIWYG ITEM"})

	container.AddFirstReference(zero, first)
	container.AddLastReference(zero, second)
	container.AddLastReference(zero, third)

	return container
}
