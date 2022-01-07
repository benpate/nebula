package nebula

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/stretchr/testify/require"
)

func TestAddItem_Rows_ABOVE(t *testing.T) {

	library := NewLibrary()

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "ABOVE",
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

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "ABOVE",
			"check":    container.GetChecksum(1),
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

func TestAddItem_Rows_BEFORE(t *testing.T) {

	library := NewLibrary()

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
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

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "BEFORE",
			"check":    container.GetChecksum(1),
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

func TestAddItem_Rows_BELOW(t *testing.T) {

	library := NewLibrary()

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "BELOW",
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

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "BELOW",
			"check":    container.GetChecksum(1),
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

func TestAddItem_Rows_AFTER(t *testing.T) {

	library := NewLibrary()

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
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

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "AFTER",
			"check":    container.GetChecksum(1),
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

func TestAddItem_Rows_LEFT(t *testing.T) {

	library := NewLibrary()

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "LEFT",
			"check":    container.GetChecksum(0),
		})

		require.Nil(t, err)
		require.Equal(t, 0, itemID)
		require.Equal(t, 4, len(container))
		require.Equal(t, "LAYOUT", container.GetType(0))
		require.Equal(t, "COLS", container[0].Data["style"])
		require.Equal(t, []int{2, 3}, container.GetRefs(0))

		require.Equal(t, "HTML", container.GetType(1))
		require.Equal(t, "FIRST HTML BLOCK", container[1].Data["html"])

		require.Equal(t, "TEXT", container.GetType(2))

		require.Equal(t, "LAYOUT", container.GetType(3))
		require.Equal(t, "ROWS", container[3].Data["style"])
		require.Equal(t, []int{1}, container.GetRefs(3))
	}

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "LEFT",
			"check":    container.GetChecksum(0),
		})

		require.Nil(t, err)
		require.Equal(t, 0, itemID)
		require.Equal(t, 4, len(container))
		require.Equal(t, "LAYOUT", container.GetType(0))
		require.Equal(t, "ROWS", container[0].Data["style"])
		require.Equal(t, []int{1}, container.GetRefs(0))

		require.Equal(t, "LAYOUT", container.GetType(1))
		require.Equal(t, "COLS", container[1].Data["style"])
		require.Equal(t, []int{2, 3}, container.GetRefs(1))

		require.Equal(t, "TEXT", container.GetType(2))

		require.Equal(t, "HTML", container.GetType(3))
		require.Equal(t, "FIRST HTML BLOCK", container[3].Data["html"])
	}
}

func TestAddItem_Rows_RIGHT(t *testing.T) {

	library := NewLibrary()

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   0,
			"itemType": "TEXT",
			"place":    "RIGHT",
			"check":    container.GetChecksum(0),
		})

		require.Nil(t, err)
		require.Equal(t, 0, itemID)
		require.Equal(t, 4, len(container))
		require.Equal(t, "LAYOUT", container.GetType(0))
		require.Equal(t, "COLS", container[0].Data["style"])
		require.Equal(t, []int{3, 2}, container.GetRefs(0))

		require.Equal(t, "HTML", container.GetType(1))
		require.Equal(t, "FIRST HTML BLOCK", container[1].Data["html"])

		require.Equal(t, "TEXT", container.GetType(2))

		require.Equal(t, "LAYOUT", container.GetType(3))
		require.Equal(t, "ROWS", container[3].Data["style"])
		require.Equal(t, []int{1}, container.GetRefs(3))
	}

	{
		container := getTestRows()

		itemID, err := container.Post(&library, datatype.Map{
			"type":     "add-item",
			"itemId":   1,
			"itemType": "TEXT",
			"place":    "RIGHT",
			"check":    container.GetChecksum(0),
		})

		require.Nil(t, err)
		require.Equal(t, 0, itemID)
		require.Equal(t, 4, len(container))
		require.Equal(t, "LAYOUT", container.GetType(0))
		require.Equal(t, "ROWS", container[0].Data["style"])
		require.Equal(t, []int{1}, container.GetRefs(0))

		require.Equal(t, "LAYOUT", container.GetType(1))
		require.Equal(t, "COLS", container[1].Data["style"])
		require.Equal(t, []int{3, 2}, container.GetRefs(1))

		require.Equal(t, "TEXT", container.GetType(2))

		require.Equal(t, "HTML", container.GetType(3))
		require.Equal(t, "FIRST HTML BLOCK", container[3].Data["html"])
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
