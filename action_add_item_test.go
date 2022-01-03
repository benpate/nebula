package nebula

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/stretchr/testify/require"
)

func TestItem_Split_ABOVE(t *testing.T) {

	library := NewLibrary()

	container := NewContainer()

	firstID := container.NewItemWithInit(&library, ItemTypeHTML, datatype.Map{"html": "FIRST HTML ITEM"})

	itemID, err := container.Execute(&library, datatype.Map{
		"type":     "add-item",
		"itemId":   0,
		"itemType": "WYSIWYG",
		"place":    "ABOVE",
		"check":    container.GetChecksum(firstID),
	})

	require.Nil(t, err)
	require.Equal(t, 0, itemID)
	require.Equal(t, 3, container.Len())

	{
		item := container.GetItem(0)
		require.Equal(t, "LAYOUT", item.Type)
		require.Equal(t, "ROWS", item.GetString("style"))
		require.Equal(t, []int{1, 2}, item.Refs)
	}
	{
		item := container.GetItem(1)
		require.Equal(t, "WYSIWYG", item.Type)
		require.Equal(t, "", item.GetString("html"))
	}
	{
		item := container.GetItem(2)
		require.Equal(t, "HTML", item.Type)
		require.Equal(t, "FIRST HTML ITEM", item.GetString("html"))
	}
}

func TestItem_Split_BELOW(t *testing.T) {

	library := NewLibrary()

	container := NewContainer()

	firstID := container.NewItemWithInit(&library, ItemTypeHTML, datatype.Map{"html": "FIRST HTML ITEM"})

	itemID, err := container.Execute(&library, datatype.Map{
		"type":     "add-item",
		"itemId":   0,
		"itemType": "WYSIWYG",
		"place":    "BELOW",
		"check":    container.GetChecksum(firstID),
	})

	require.Nil(t, err)
	require.Equal(t, 0, itemID)
	require.Equal(t, 3, container.Len())

	{
		item := container.GetItem(0)
		require.Equal(t, "LAYOUT", item.Type)
		require.Equal(t, "ROWS", item.GetString("style"))
		require.Equal(t, []int{2, 1}, item.Refs)
	}
	{
		item := container.GetItem(1)
		require.Equal(t, "WYSIWYG", item.Type)
		require.Equal(t, "", item.GetString("html"))
	}
	{
		item := container.GetItem(2)
		require.Equal(t, "HTML", item.Type)
		require.Equal(t, "FIRST HTML ITEM", item.GetString("html"))
	}
}

func TestItem_Split_LEFT(t *testing.T) {

	library := NewLibrary()

	container := NewContainer()

	firstID := container.NewItemWithInit(&library, ItemTypeHTML, datatype.Map{"html": "FIRST HTML ITEM"})

	itemID, err := container.Execute(&library, datatype.Map{
		"type":     "add-item",
		"itemId":   0,
		"itemType": "WYSIWYG",
		"place":    "LEFT",
		"check":    container.GetChecksum(firstID),
	})

	require.Nil(t, err)
	require.Equal(t, 0, itemID)
	require.Equal(t, 3, container.Len())

	{
		item := container.GetItem(0)
		require.Equal(t, "LAYOUT", item.Type)
		require.Equal(t, "COLS", item.GetString("style"))
		require.Equal(t, []int{1, 2}, item.Refs)
	}
	{
		item := container.GetItem(1)
		require.Equal(t, "WYSIWYG", item.Type)
		require.Equal(t, "", item.GetString("html"))
	}
	{
		item := container.GetItem(2)
		require.Equal(t, "HTML", item.Type)
		require.Equal(t, "FIRST HTML ITEM", item.GetString("html"))
	}
}

func TestItem_Split_RIGHT(t *testing.T) {

	library := NewLibrary()

	container := NewContainer()

	firstID := container.NewItemWithInit(&library, ItemTypeHTML, datatype.Map{"html": "FIRST HTML ITEM"})

	itemID, err := container.Execute(&library, datatype.Map{
		"type":     "add-item",
		"itemId":   0,
		"itemType": "WYSIWYG",
		"place":    "RIGHT",
		"check":    container.GetChecksum(firstID),
	})

	require.Nil(t, err)
	require.Equal(t, 0, itemID)
	require.Equal(t, 3, container.Len())

	{
		item := container.GetItem(0)
		require.Equal(t, "LAYOUT", item.Type)
		require.Equal(t, "COLS", item.GetString("style"))
		require.Equal(t, []int{2, 1}, item.Refs)
	}
	{
		item := container.GetItem(1)
		require.Equal(t, "WYSIWYG", item.Type)
		require.Equal(t, "", item.GetString("html"))
	}
	{
		item := container.GetItem(2)
		require.Equal(t, "HTML", item.Type)
		require.Equal(t, "FIRST HTML ITEM", item.GetString("html"))
	}
}

func TestItem_Split_BEFORE(t *testing.T) {

	library := NewLibrary()

	container := NewContainer()

	firstID := container.NewItemWithInit(&library, ItemTypeHTML, datatype.Map{"html": "FIRST HTML ITEM"})

	itemID, err := container.Execute(&library, datatype.Map{
		"type":     "add-item",
		"itemId":   0,
		"itemType": "WYSIWYG",
		"place":    "BEFORE",
		"check":    container.GetChecksum(firstID),
	})

	require.Nil(t, err)
	require.Equal(t, 0, itemID)
	require.Equal(t, 3, container.Len())

	{
		item := container.GetItem(0)
		require.Equal(t, "LAYOUT", item.Type)
		require.Equal(t, "ROWS", item.GetString("style"))
		require.Equal(t, []int{1, 2}, item.Refs)
	}
	{
		item := container.GetItem(1)
		require.Equal(t, "WYSIWYG", item.Type)
		require.Equal(t, "", item.GetString("html"))
	}
	{
		item := container.GetItem(2)
		require.Equal(t, "HTML", item.Type)
		require.Equal(t, "FIRST HTML ITEM", item.GetString("html"))
	}
}

func TestItem_Split_AFTER(t *testing.T) {

	library := NewLibrary()

	container := NewContainer()

	firstID := container.NewItemWithInit(&library, ItemTypeHTML, datatype.Map{"html": "FIRST HTML ITEM"})

	itemID, err := container.Execute(&library, datatype.Map{
		"type":     "add-item",
		"itemId":   0,
		"itemType": "WYSIWYG",
		"place":    "AFTER",
		"check":    container.GetChecksum(firstID),
	})

	require.Nil(t, err)
	require.Equal(t, 0, itemID)
	require.Equal(t, 3, container.Len())

	{
		item := container.GetItem(0)
		require.Equal(t, "LAYOUT", item.Type)
		require.Equal(t, "ROWS", item.GetString("style"))
		require.Equal(t, []int{2, 1}, item.Refs)
	}
	{
		item := container.GetItem(1)
		require.Equal(t, "WYSIWYG", item.Type)
		require.Equal(t, "", item.GetString("html"))
	}
	{
		item := container.GetItem(2)
		require.Equal(t, "HTML", item.Type)
		require.Equal(t, "FIRST HTML ITEM", item.GetString("html"))
	}
}