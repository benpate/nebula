package nebula

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
)

// ItemTypeHTML describes an HTML code editor
const ItemTypeHTML = "HTML"

// HTML widget displays plain HTML content
type HTML struct{}

// View renders the HTML content of the widget
func (w HTML) View(b *html.Builder, container *Container, id int) {
	item := container.GetItem(id)
	result := item.GetString("html")
	b.WriteString(result)
}

// Edit renders a plain text editor for editing HTML code
func (w HTML) Edit(b *html.Builder, container *Container, id int, endpoint string) {
	item := container.GetItem(id)
	result := item.GetString("html")
	idString := convert.String(id)

	b.Form("", "").
		Data("hx-post", endpoint).
		Data("hx-trigger", "autosave").
		Data("hx-swap", "none")

	b.Input("hidden", "action").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "check").Value(item.Check)
	b.Container("textarea").Name("html").Script("install Autosize install Autosave").InnerHTML(result)
}

// Validate cleans the container for invalid content
func (w HTML) Validate(container *Container, index int) {
	originalHTML := (*container)[index].GetString("html")
	cleanHTML := htmlPolicy.Sanitize(originalHTML)
	(*container)[index].Data.SetString("html", cleanHTML)
}
