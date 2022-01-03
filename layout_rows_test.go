package tests

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/benpate/nebula"
	"github.com/benpate/nebula/transaction"
	"github.com/benpate/nebula/vocabulary"
	"github.com/stretchr/testify/require"
)

func TestAddItem_Rows_ABOVE(t *testing.T) {

	library := nebula.NewLibrary()
	vocabulary.All(&library)

	{
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "ABOVE",
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
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "ABOVE",
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

func TestAddItem_Rows_BEFORE(t *testing.T) {

	library := nebula.NewLibrary()
	vocabulary.All(&library)

	{
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "BEFORE",
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
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "BEFORE",
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

func TestAddItem_Rows_BELOW(t *testing.T) {

	library := nebula.NewLibrary()
	vocabulary.All(&library)

	{
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "BELOW",
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
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "BELOW",
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

func TestAddItem_Rows_AFTER(t *testing.T) {

	library := nebula.NewLibrary()
	vocabulary.All(&library)

	{
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "AFTER",
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
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "AFTER",
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

func TestAddItem_Rows_LEFT(t *testing.T) {

	library := nebula.NewLibrary()
	vocabulary.All(&library)

	{
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "LEFT",
			"check":    container.GetItem(0).Check,
		})

		_, err := txn.Execute(&library, &container)

		require.Nil(t, err)
		require.Equal(t, 4, len(container))
		require.Equal(t, "LAYOUT", container[0].Type)
		require.Equal(t, "COLS", container[0].Data["style"])
		require.Equal(t, []int{2, 3}, container[0].Refs)

		require.Equal(t, "HTML", container[1].Type)
		require.Equal(t, "FIRST HTML BLOCK", container[1].Data["html"])

		require.Equal(t, "TEXT", container[2].Type)

		require.Equal(t, "LAYOUT", container[3].Type)
		require.Equal(t, "ROWS", container[3].Data["style"])
		require.Equal(t, []int{1}, container[3].Refs)
	}

	{
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "LEFT",
			"check":    container.GetItem(0).Check,
		})

		_, err := txn.Execute(&library, &container)

		require.Nil(t, err)
		require.Equal(t, 4, len(container))
		require.Equal(t, "LAYOUT", container[0].Type)
		require.Equal(t, "ROWS", container[0].Data["style"])
		require.Equal(t, []int{1}, container[0].Refs)

		require.Equal(t, "LAYOUT", container[1].Type)
		require.Equal(t, "COLS", container[1].Data["style"])
		require.Equal(t, []int{2, 3}, container[1].Refs)

		require.Equal(t, "TEXT", container[2].Type)

		require.Equal(t, "HTML", container[3].Type)
		require.Equal(t, "FIRST HTML BLOCK", container[3].Data["html"])
	}
}

func TestAddItem_Rows_RIGHT(t *testing.T) {

	library := nebula.NewLibrary()
	vocabulary.All(&library)

	{
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "RIGHT",
			"check":    container.GetItem(0).Check,
		})

		_, err := txn.Execute(&library, &container)

		require.Nil(t, err)
		require.Equal(t, 4, len(container))
		require.Equal(t, "LAYOUT", container[0].Type)
		require.Equal(t, "COLS", container[0].Data["style"])
		require.Equal(t, []int{3, 2}, container[0].Refs)

		require.Equal(t, "HTML", container[1].Type)
		require.Equal(t, "FIRST HTML BLOCK", container[1].Data["html"])

		require.Equal(t, "TEXT", container[2].Type)

		require.Equal(t, "LAYOUT", container[3].Type)
		require.Equal(t, "ROWS", container[3].Data["style"])
		require.Equal(t, []int{1}, container[3].Refs)
	}

	{
		container := getTestRows()

		txn := transaction.Parse(datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "RIGHT",
			"check":    container.GetItem(0).Check,
		})

		_, err := txn.Execute(&library, &container)

		require.Nil(t, err)
		require.Equal(t, 4, len(container))
		require.Equal(t, "LAYOUT", container[0].Type)
		require.Equal(t, "ROWS", container[0].Data["style"])
		require.Equal(t, []int{1}, container[0].Refs)

		require.Equal(t, "LAYOUT", container[1].Type)
		require.Equal(t, "COLS", container[1].Data["style"])
		require.Equal(t, []int{3, 2}, container[1].Refs)

		require.Equal(t, "TEXT", container[2].Type)

		require.Equal(t, "HTML", container[3].Type)
		require.Equal(t, "FIRST HTML BLOCK", container[3].Data["html"])
	}
}

func getTestRows() nebula.Container {

	return nebula.Container{
		{
			Type: nebula.ItemTypeLayout,
			Refs: []int{1},
			Data: datatype.Map{
				"style": nebula.LayoutStyleRows,
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
