package nebula

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/benpate/nebula"
	"github.com/benpate/nebula/transaction"
	"github.com/benpate/nebula/vocabulary"
	"github.com/stretchr/testify/require"
)

func getTestColumns() nebula.Container {

	return nebula.Container{
		{
			Type: nebula.ItemTypeLayout,
			Refs: []int{1},
			Data: datatype.Map{
				"style": nebula.LayoutStyleColumns,
			},
		},
		{
			Type: nebula.ItemTypeHTML,
			Data: datatype.Map{
				"html": "FIRST HTML BLOCK",
			},
		},
	}
}

func TestAddItem_Columns_RIGHT(t *testing.T) {

	library := nebula.NewLibrary()
	vocabulary.All(&library)

	{
		container := getTestColumns()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "RIGHT",
			"check":    container.GetItem(0).Check,
		})

		_, err := txn.Execute(&library, &container)

		require.Nil(t, err)
		require.Equal(t, 3, len(container))
		require.Equal(t, "LAYOUT", container[0].Type)
		require.Equal(t, []int{1, 2}, container[0].Refs)
		require.Equal(t, "HTML", container[1].Type)
		require.Equal(t, "TEXT", container[2].Type)
	}

	{
		container := getTestColumns()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "RIGHT",
			"check":    container.GetItem(1).Check,
		})

		_, err := txn.Execute(&library, &container)

		require.Nil(t, err)
		require.Equal(t, 3, len(container))
		require.Equal(t, "LAYOUT", container[0].Type)
		require.Equal(t, []int{1, 2}, container[0].Refs)
		require.Equal(t, "HTML", container[1].Type)
		require.Equal(t, "TEXT", container[2].Type)
	}

}

func TestAddItem_Columns_LEFT(t *testing.T) {

	library := nebula.NewLibrary()
	vocabulary.All(&library)

	{
		container := getTestColumns()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "LEFT",
			"check":    container.GetItem(0).Check,
		})

		_, err := txn.Execute(&library, &container)

		require.Nil(t, err)
		require.Equal(t, 3, len(container))
		require.Equal(t, "LAYOUT", container[0].Type)
		require.Equal(t, []int{2, 1}, container[0].Refs)
		require.Equal(t, "HTML", container[1].Type)
		require.Equal(t, "TEXT", container[2].Type)
	}

	{
		container := getTestColumns()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "LEFT",
			"check":    container.GetItem(1).Check,
		})

		_, err := txn.Execute(&library, &container)

		require.Nil(t, err)
		require.Equal(t, 3, len(container))
		require.Equal(t, "LAYOUT", container[0].Type)
		require.Equal(t, []int{2, 1}, container[0].Refs)
		require.Equal(t, "HTML", container[1].Type)
		require.Equal(t, "TEXT", container[2].Type)
	}

}
