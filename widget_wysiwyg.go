package nebula

import (
	"net/url"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/html"
)

// ItemTypeWYSIWYG describes a what-you-see-is-what-you-get content editor
const ItemTypeWYSIWYG = "WYSIWYG"

type WYSIWYG struct{}

// Init sets up an empty "html" property
func (w WYSIWYG) Init(container *Container, id int) {
	(*container)[id].Set("html", "")
}

func (w WYSIWYG) View(b *html.Builder, container *Container, id int) {
	item := container.GetItem(id)
	result := item.GetString("html")
	b.WriteString(result)
}

func (w WYSIWYG) Edit(b *html.Builder, container *Container, id int, endpoint string) {
	item := container.GetItem(id)
	result := item.GetString("html")
	idString := convert.String(id)

	b.Form("", "").
		Data("hx-post", endpoint).
		Data("hx-trigger", "autosave").
		Data("hx-swap", "none")

	// Form fields here
	b.Input("hidden", "type").Value("update-item")
	b.Input("hidden", "itemId").Value(idString)
	b.Input("hidden", "html").Value(idString)
	b.Input("hidden", "check").Value(item.Check)

	b.Div().Class("wysiwyg").Script("install wysiwyg(name:'html') install hotkey")
	b.Div().Class("wysiwyg-toolbar").Attr("hidden", "true").ID("toolbar-" + idString)
	{
		b.Span().Class("wysiwyg-toolbar-group").EndBracket()
		b.Button().Data("command", "formatBlock").Data("command-value", "h1").InnerHTML("H1").Close()
		b.Button().Data("command", "formatBlock").Data("command-value", "h2").InnerHTML("H2").Close()
		b.Button().Data("command", "formatBlock").Data("command-value", "h3").InnerHTML("H3").Close()
		b.Button().Data("command", "formatBlock").Data("command-value", "p").InnerHTML("P").Close()
		b.Close()
	}
	{
		b.Span().Class("wysiwyg-toolbar-group").EndBracket()
		b.Button().Data("command", "bold").Data("hotkey", "b").InnerHTML("B").Close()
		b.Button().Data("command", "italic").Data("hotkey", "i").InnerHTML("I").Close()
		b.Button().Data("command", "underline").Data("hotkey", "u").InnerHTML("U").Close()
		b.Close()
	}
	{
		b.Span().Class("wysiwyg-toolbar-group").EndBracket()
		b.Button().Script("on click log me get prompt('Enter URL') log it call document.execCommand('link', result)").InnerHTML("Link").Close()
		b.Button().Data("command", "unlink").InnerHTML("Unlink").Close()
		b.Close()
	}
	{
		b.Span().Class("wysiwyg-toolbar-group").EndBracket()
		b.Button().Data("command", "undo").Data("hotkey", "z").InnerHTML("Undo").Close()
		b.Button().Data("command", "redo").Data("hotkey", "Z").InnerHTML("Redo").Close()
		b.Close()
	}
	b.Close()

	b.Div().Class("wysiwyg-editor").Script("install Autosave").InnerHTML(result)
	b.CloseAll()
}

func (w WYSIWYG) Prop(container *Container, id int, params url.Values) (string, error) {
	return "", derp.New(derp.CodeNotFoundError, "content.WYSIWYG.Prop", "Unrecognized panel", params)
}
