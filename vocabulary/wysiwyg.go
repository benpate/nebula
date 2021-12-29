package vocabulary

import (
	"net/url"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/nebula"
)

const ItemTypeWYSIWYG = "WYSIWYG"

type WYSIWYG struct{}

func (w WYSIWYG) View(b *html.Builder, container *nebula.Container, id int) {
	item := container.GetItem(id)
	result := item.GetString("html")
	b.WriteString(result)
}

func (w WYSIWYG) Edit(b *html.Builder, container *nebula.Container, id int, endpoint string) {
	item := container.GetItem(id)
	result := item.GetString("html")
	idString := convert.String(id)

	b.Form("", "").
		Data("hx-post", endpoint).
		Data("hx-trigger", "save").
		Data("hx-swap", "none")

	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "check").Value(item.Check)
	b.Input("hidden", "html").Value("").ID("html-" + idString)
	b.Div().Script("install wysiwyg(id:'html-" + idString + "')").InnerHTML(result).Close()
	// b.Div().Style("min-height:1em;").Script("on click 1 call makeWYSIWYG(me)").InnerHTML(result)

	b.CloseAll()
}

func (w WYSIWYG) Prop(container *nebula.Container, id int, params url.Values) (string, error) {
	return "", derp.New(derp.CodeNotFoundError, "content.WYSIWYG.Prop", "Unrecognized panel", params)
}
