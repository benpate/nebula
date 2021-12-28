package vocabulary

import (
	"net/url"

	"github.com/benpate/content"
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/html"
)

const ItemTypeWYSIWYG = "WYSIWYG"

type WYSIWYG struct{}

func (w WYSIWYG) View(b *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)
	result := item.GetString("html")
	b.WriteString(result)
}

func (w WYSIWYG) Edit(b *html.Builder, c content.Content, id int, endpoint string) {
	item := c.GetItem(id)
	result := item.GetString("html")
	idString := convert.String(id)

	b.Form("", "").
		Data("hx-post", endpoint).
		Data("hx-trigger", "quill:blur").
		Data("hx-swap", "none")

	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "check").Value(item.Check)
	b.Input("hidden", "html").Value("") // Filled by Javascript on quill:blur
	b.Div().Style("min-height:1em;").Script("on click 1 call makeWYSIWYG(me)").InnerHTML(result)

	b.CloseAll()
}

func (w WYSIWYG) Prop(c content.Content, id int, params url.Values) (string, error) {
	return "", derp.New(derp.CodeNotFoundError, "content.WYSIWYG.Prop", "Unrecognized panel", params)
}