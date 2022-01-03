package nebula

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/benpate/nebula"
	"github.com/benpate/nebula/transaction"
	"github.com/benpate/nebula/vocabulary"
)

func TestContainer(t *testing.T) {

	library := nebula.NewLibrary()
	vocabulary.All(&library)

	container := nebula.NewContainer()

	layoutID := container.NewItemWithInit(&library, nebula.ItemTypeLayout, datatype.Map{"style": nebula.LayoutStyleRows})
	container[1].Data["html"] = "FIRST HTML"

	second := container.NewItemWithInit(&library, nebula.ItemTypeWYSIWYG, nil)
	container[second].Data["html"] = "SECOND HTML ITEM"
	container.AddLastReference(layoutID, second)

	txn := transaction.Parse(datatype.Map{
		"type":     "add-item",
		"itemId":   2,
		"itemType": "WYSIWYG",
		"place":    "RIGHT",
		"check":    container.GetItem(second).Check,
	})

	txn.Execute(&library, &container)
}
