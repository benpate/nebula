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
	b.Input("hidden", "action").Value("update-item")
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
		b.Button().Data("command", "bold").Aria("keyshortcuts", "Ctrl+B").InnerHTML("B").Close()
		b.Button().Data("command", "italic").Aria("keyshortcuts", "Ctrl+I").InnerHTML("I").Close()
		b.Button().Data("command", "underline").Aria("keyshortcuts", "Ctrl+U").InnerHTML("U").Close()
		b.Close()
	}
	{
		b.Span().Class("wysiwyg-toolbar-group").EndBracket()
		b.Button().Data("command", "createLink").Aria("keyshortcuts", "Ctrl+K")
		b.Container("i").Class("fa-solid", "fa-link").Close()
		b.Close()
		b.Button().Data("command", "unlink").Aria("keyshortcuts", "Ctrl+Shift+K")
		b.Container("i").Class("fa-solid", "fa-unlink").Close()
		b.Close()
		b.Close()
	}
	{
		b.Span().Class("wysiwyg-toolbar-group").Attr("hidden", "true").EndBracket()
		b.Button().Data("command", "cut").Aria("keyshortcuts", "Ctrl+X").InnerHTML("Cut").Close()
		b.Button().Data("command", "copy").Aria("keyshortcuts", "Ctrl+C").InnerHTML("Copy").Close()
		b.Button().Data("command", "paste").Aria("keyshortcuts", "Ctrl+V").InnerHTML("Paste").Close()
		b.Button().Data("command", "undo").Aria("keyshortcuts", "Ctrl+Z").InnerHTML("Undo").Close()
		b.Button().Data("command", "redo").Aria("keyshortcuts", "Ctrl+Shift+Z").InnerHTML("Redo").Close()
		b.Close()
	}
	b.Close()

	b.Div().Class("wysiwyg-editor").Script("install Autosave").InnerHTML(result)
	b.CloseAll()
}

func (w WYSIWYG) Prop(container *Container, id int, params url.Values) (string, error) {
	return "", derp.NewNotFoundError("content.WYSIWYG.Prop", "Unrecognized panel", params)
}
