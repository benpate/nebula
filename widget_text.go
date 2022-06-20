package nebula

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	htmlconv "github.com/benpate/rosetta/html"
)

// ItemTypeText describes a plain text editor
const ItemTypeText = "TEXT"

type Text struct{}

// View returns the Text widget, rendered as HTML
func (w Text) View(b *html.Builder, container *Container, id int) {
	item := container.GetItem(id)
	result := item.GetString("text")
	result = htmlconv.FromText(result)
	b.WriteString(result)
}

// Edit returns a textarea where the text content can be edited
func (w Text) Edit(b *html.Builder, container *Container, id int, endpoint string) {
	item := container.GetItem(id)

	b.Form("", "").
		Data("hx-post", endpoint).
		Data("hx-trigger", "autosave").
		Data("hx-swap", "none")

	b.Input("hidden", "action").Value("update-item")
	b.Input("hidden", "itemId").Value(convert.String(id))
	b.Input("hidden", "check").Value(item.Check)
	b.Container("textarea").Name("text").Script("install Autosize install Autosave").InnerHTML(item.GetString("text"))
}

// Validate cleans the container for invalid content
func (w Text) Validate(container *Container, index int) {
	originalText := (*container)[index].GetString("text")
	cleanText := textPolicy.Sanitize(originalText)
	(*container)[index].Data.SetString("text", cleanText)
}
