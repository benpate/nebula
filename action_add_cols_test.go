package nebula

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/stretchr/testify/require"
)

func getTestColumns() Container {

	return Container{
		{
			Type: ItemTypeLayout,
			Refs: []int{1},
			Data: datatype.Map{
				"style": LayoutStyleColumns,
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

func TestAddItem_Columns_RIGHT(t *testing.T) {

	library := NewLibrary()

	{
		container := getTestColumns()

		_, err := container.Execute(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "RIGHT",
			"check":    container.GetChecksum(0),
		})

		require.Nil(t, err)
		require.Equal(t, 3, len(container))
		{
			item := container.GetItem(0)
			require.Equal(t, "LAYOUT", item.Type)
			require.Equal(t, []int{1, 2}, item.Refs)
		}
		{
			item := container.GetItem(1)
			require.Equal(t, "HTML", item.Type)
		}
		{
			item := container.GetItem(2)
			require.Equal(t, "TEXT", item.Type)
		}
	}

	{
		container := getTestColumns()

		_, err := container.Execute(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "RIGHT",
			"check":    container.GetChecksum(1),
		})

		require.Nil(t, err)
		require.Equal(t, 3, len(container))
		require.Equal(t, "LAYOUT", container.GetType(0))
		require.Equal(t, []int{1, 2}, container.GetRefs(0))
		require.Equal(t, "HTML", container.GetType(1))
		require.Equal(t, "TEXT", container.GetType(2))
	}

}

func TestAddItem_Columns_LEFT(t *testing.T) {

	library := NewLibrary()

	{
		container := getTestColumns()

		_, err := container.Execute(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "LEFT",
			"check":    container.GetChecksum(0),
		})

		require.Nil(t, err)
		require.Equal(t, 3, len(container))
		require.Equal(t, "LAYOUT", container.GetType(0))
		require.Equal(t, []int{2, 1}, container.GetRefs(0))
		require.Equal(t, "HTML", container.GetType(1))
		require.Equal(t, "TEXT", container.GetType(2))
	}

	{
		container := getTestColumns()

		_, err := container.Execute(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "LEFT",
			"check":    container.GetChecksum(1),
		})

		require.Nil(t, err)
		require.Equal(t, 3, len(container))
		require.Equal(t, "LAYOUT", container.GetType(0))
		require.Equal(t, []int{2, 1}, container.GetRefs(0))
		require.Equal(t, "HTML", container.GetType(1))
		require.Equal(t, "TEXT", container.GetType(2))
	}

}
