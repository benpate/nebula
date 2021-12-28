package vocabulary

import (
	"github.com/benpate/content"
	"github.com/benpate/convert"
	"github.com/benpate/html"
	"github.com/benpate/htmlconv"
)

const ItemTypeText = "TEXT"

type Text struct{}

// View returns the Text widget, rendered as HTML
func (w Text) View(b *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)
	result := item.GetString("text")
	result = htmlconv.FromText(result)
	b.WriteString(result)
}

// Edit returns a textarea where the text content can be edited
func (w Text) Edit(b *html.Builder, c content.Content, id int, endpoint string) {
	item := c.GetItem(id)

	b.Form("", "").Data("hx-post", endpoint).Data("hx-trigger", "blur")
	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(convert.String(id))
	b.Input("hidden", "check").Value(item.Check)
	b.Container("textarea").Name("text").InnerHTML(item.GetString("text"))
}