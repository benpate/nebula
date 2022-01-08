package nebula

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/html"
)

type AddItem struct {
	ItemID    int    `json:"itemId"    form:"itemId"`    // ID of the layout that will hold the new item
	SubItemID int    `json:"subItemID" form:"subItemId"` // ID of the item used for relative positioning
	Place     string `json:"place"     form:"place"`     // Position of the new element (BEFORE, AFTER, ABOVE, BELOW, LEFT, RIGHT) relative to the index
	ItemType  string `json:"itemType"  form:"itemType"`  // Type of content item to add
	Style     string `json:"style"     form:"style"`     // Optional "style" aregument for certain types (like layouts)
	Check     string `json:"check"     form:"check"`     // Checksum to validation transaction.
}

func (txn AddItem) Get(library *Library, container *Container, endpoint string) string {

	b := html.New()

	b.H1().InnerHTML("Add Another Section").Close()

	b.Div().Class("table")

	for _, itemType := range ItemTypes() {
		b.Div().Attr("tabindex", "0")
		b.Form("", "").Data("hx-post", endpoint).Data("hx-trigger", "click")
		b.Input("hidden", "action").Value("add-item").Close()
		b.Input("hidden", "itemId").Value(convert.String(txn.ItemID)).Close()
		b.Input("hidden", "subItemId").Value(convert.String(txn.SubItemID)).Close()
		b.Input("hidden", "itemType").Value(itemType.Code)

		for key, value := range itemType.Data {
			b.Input("hidden", key).Value(convert.String(value)).Close()
		}

		b.Input("hidden", "place").Value(txn.Place).Close()
		b.Input("hidden", "check").Value(txn.Check).Close()
		b.Div().InnerHTML(itemType.Label).Close()
		b.Div().InnerHTML(itemType.Description).Close()
		b.Close() // Form
		b.Close() // Div
	}

	b.Close()

	// Close button (because now we have to do it ourselves)
	b.Div()
	b.Button().Script("on click send closeModal").InnerHTML("Close")
	b.CloseAll()

	return b.String()

}

// Execute performs the AddItem transaction on the provided content structure
func (txn AddItem) Post(library *Library, container *Container) (int, error) {

	/*** If we can't append to the item directly, then try to insert into to its parent instead */

	parent := container.GetItem(txn.ItemID)

	if err := parent.Validate(txn.Check); err != nil {
		return -1, derp.Wrap(err, "nebula.AddItem.Post", "Invalid checksum", txn)
	}

	if container.IsNil(txn.SubItemID) {
		return -1, derp.New(derp.CodeBadRequestError, "nebula.AddItem.Post", "Invalid Reference Item", txn)
	}

	newItemID := container.NewItemWithInit(library, txn.ItemType, nil)

	container.AddReference(txn.ItemID, txn.Place, txn.SubItemID, newItemID)
	container.NewChecksum(txn.ItemID)

	// Since we have moved things around, replace the whole parent
	if txn.ItemID > 0 {
		return txn.ItemID, nil
	}

	return 0, nil
}
