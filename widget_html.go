package nebula

import (
	"github.com/benpate/convert"
	"github.com/benpate/html"
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

	b.Form("", "").Data("hx-post", endpoint).Data("hx-trigger", "blur")
	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "check").Value(item.Check)
	b.Container("textarea").Name("html").InnerHTML(result)
}
