package editor

import (
	"strconv"

	"github.com/benpate/content"
	"github.com/benpate/html"
)

func (e Editor) HTML(b *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)
	result := item.GetString("html")
	idString := strconv.Itoa(id)

	b.Form("post", e.Endpoint)
	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "check").Value(item.Check)
	b.Container("textarea").Name("html").InnerHTML(result)
}
