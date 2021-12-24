package editor

import (
	"strconv"

	"github.com/benpate/content"
	"github.com/benpate/html"
)

func (e Editor) WYSIWYG(b *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)
	result := item.GetString("html")
	idString := strconv.Itoa(id)

	b.Form("", "").
		Data("hx-post", e.Endpoint).
		Data("hx-trigger", "quill:blur").
		Data("hx-swap", "none")

	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "check").Value(item.Check)
	b.Input("hidden", "html").Value("") // Filled by Javascript on quill:blur
	b.Div().Style("min-height:1em;").Script("on click 1 call makeWYSIWYG(me)").InnerHTML(result)

	b.CloseAll()
}
