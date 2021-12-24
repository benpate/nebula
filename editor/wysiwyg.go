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

	b.Form("", "").Data("hx-post", e.Endpoint).Data("hx-trigger", "blur").Script("on blur log me")
	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "check").Value(item.Check)
	b.Div().Script("install wysiwyg").InnerHTML(result)
	// b.Input("hidden", "html")
	// b.Div().Class("ck-editor").InnerHTML(result)
}
