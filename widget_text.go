package nebula

import (
	"github.com/benpate/convert"
	"github.com/benpate/html"
	"github.com/benpate/htmlconv"
)

// ItemTypeText describes a plain text editor
const ItemTypeText = "TEXT"

type Text struct{}

// View returns the Text widget, rendered as HTML
func (w Text) View(b *html.Builder, container *Container, itemID int) {
	item := container.GetItem(itemID)
	result := item.GetString("text")
	result = htmlconv.FromText(result)
	b.WriteString(result)
}

// Edit returns a textarea where the text content can be edited
func (w Text) Edit(b *html.Builder, container *Container, itemID int, endpoint string) {
	item := container.GetItem(itemID)
	idString := convert.String(itemID)

	b.Form("", "").
		Data("id", idString).
		Data("hx-post", endpoint).
		Data("hx-trigger", "autosave").
		Data("hx-swap", "none")

	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "check").Value(item.Check)
	b.Container("textarea").Name("text").Script("install Autosize install Autosave").InnerHTML(item.GetString("text"))
}
