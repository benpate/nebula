package nebula

import (
	"testing"

	"github.com/benpate/datatype"
)

func TestItem(t *testing.T) {

	library := NewLibrary()

	container := NewContainer()

	layoutID := container.NewItemWithInit(&library, ItemTypeLayout, datatype.Map{"style": LayoutStyleRows})
	container[1].Data["html"] = "FIRST HTML"

	second := container.NewItemWithInit(&library, ItemTypeWYSIWYG, nil)
	container[second].Data["html"] = "SECOND HTML ITEM"
	container.AddLastReference(layoutID, second)

	txn := ParseAction(datatype.Map{
		"type":     "add-item",
		"itemId":   2,
		"itemType": "WYSIWYG",
		"place":    "RIGHT",
		"check":    container.GetItem(second).Check,
	})

	txn.Execute(&library, &container)
}
