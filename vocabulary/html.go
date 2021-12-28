package vocabulary

import (
	"github.com/benpate/content"
	"github.com/benpate/convert"
	"github.com/benpate/html"
)

const ItemTypeHTML = "HTML"

type HTML struct{}

func (w HTML) View(b *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)
	result := item.GetString("html")
	b.WriteString(result)
}

func (w HTML) Edit(b *html.Builder, c content.Content, id int, endpoint string) {
	item := c.GetItem(id)
	result := item.GetString("html")
	idString := convert.String(id)

	b.Form("", "").Data("hx-post", endpoint).Data("hx-trigger", "blur")
	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "check").Value(item.Check)
	b.Container("textarea").Name("html").InnerHTML(result)
}
