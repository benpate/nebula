package nebula

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/stretchr/testify/require"
)

func TestItem(t *testing.T) {

	library := NewLibrary()

	container := NewContainer()

	layoutID := container.NewItemWithInit(&library, ItemTypeLayout, datatype.Map{"style": LayoutStyleRows})
	container[1].Data["html"] = "FIRST HTML"

	second := container.NewItemWithInit(&library, ItemTypeWYSIWYG, nil)
	container[second].Data["html"] = "SECOND HTML ITEM"
	container.AddLastReference(layoutID, second)

	_, err := container.Execute(&library, datatype.Map{
		"type":     "add-item",
		"itemId":   2,
		"itemType": "WYSIWYG",
		"place":    "RIGHT",
		"check":    container.GetItem(second).Check,
	})

	require.Nil(t, err)
}
