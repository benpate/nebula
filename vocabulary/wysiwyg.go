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

	b.Div().Class("wysiwyg").Script("install wysiwyg(name:'html') install hotkey")
	b.Div().Class("wysiwyg-toolbar").ID("toolbar-" + idString)
	{
		b.Span().Class("wysiwyg-toolbar-group").EndBracket()
		b.Button().Data("command", "bold").Data("hotkey", "b").InnerHTML("B").Close()
		b.Button().Data("command", "italic").Data("hotkey", "i").InnerHTML("I").Close()
		b.Button().Data("command", "underline").Data("hotkey", "u").InnerHTML("U").Close()
		b.Close()
	}
	{
		b.Span().Class("wysiwyg-toolbar-group").EndBracket()
		b.Button().Data("command", "formatBlock").Data("command-value", "h1").InnerHTML("H1").Close()
		b.Button().Data("command", "formatBlock").Data("command-value", "h2").InnerHTML("H2").Close()
		b.Button().Data("command", "formatBlock").Data("command-value", "h3").InnerHTML("H3").Close()
		b.Button().Data("command", "formatBlock").Data("command-value", "p").InnerHTML("P").Close()

		b.Close()
	}
	b.Close()

	b.Div().Class("wysiwyg-editor").InnerHTML(result)
	// b.Div().Style("min-height:1em;").Script("on click 1 call makeWYSIWYG(me)").InnerHTML(result)

	b.CloseAll()
}

func (w WYSIWYG) Prop(container *nebula.Container, id int, params url.Values) (string, error) {
	return "", derp.New(derp.CodeNotFoundError, "content.WYSIWYG.Prop", "Unrecognized panel", params)
}
