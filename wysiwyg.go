package content

import (
	"strconv"

	"github.com/benpate/html"
)

type WYSIWYG struct{}

func (widget WYSIWYG) View(b *html.Builder, content *Content, id int) {
	item := content.GetItem(id)
	result := item.GetString("html")
	b.WriteString(result)
}

func (widget WYSIWYG) Edit(b *html.Builder, content *Content, id int, endpoint string) {
	item := content.GetItem(id)
	result := item.GetString("html")
	idString := strconv.Itoa(id)

	b.Form("", "").Data("hx-post", endpoint).Data("hx-trigger", "blur").Script("on blur log me")
	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "check").Value(item.Check)
	b.Div().Script("install wysiwyg").InnerHTML(result)
	// b.Input("hidden", "html")
	// b.Div().Class("ck-editor").InnerHTML(result)
}
