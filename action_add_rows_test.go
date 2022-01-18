package nebula

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/stretchr/testify/require"
)

func TestAddItem_Rows_BEFORE(t *testing.T) {

	library := NewLibrary()

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"action":   "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "BEFORE",
			"check":    container.GetChecksum(0),
		})

		require.Nil(t, err)
		require.Equal(t, 0, itemID)
		require.Equal(t, 3, len(container))
		require.Equal(t, "LAYOUT", container.GetType(0))
		require.Equal(t, []int{2, 1}, container.GetRefs(0))
		require.Equal(t, "HTML", container.GetType(1))
		require.Equal(t, "TEXT", container.GetType(2))
	}
}

func TestAddItem_Rows_AFTER(t *testing.T) {

	library := NewLibrary()

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"action":   "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "AFTER",
			"check":    container.GetChecksum(0),
		})

		require.Nil(t, err)
		require.Equal(t, 0, itemID)
		require.Equal(t, 3, len(container))
		require.Equal(t, "LAYOUT", container.GetType(0))
		require.Equal(t, []int{1, 2}, container.GetRefs(0))
		require.Equal(t, "HTML", container.GetType(1))
		require.Equal(t, "TEXT", container.GetType(2))
	}
}

func getTestRows() Container {

	return Container{
		{
			Type: ItemTypeLayout,
			Refs: []int{1},
			Data: datatype.Map{
				"style": LayoutStyleRows,
			},
		},
		{
			Type: ItemTypeHTML,
			Data: datatype.Map{
				"html": "FIRST HTML BLOCK",
			},
		},
	}
}
