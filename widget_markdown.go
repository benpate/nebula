package nebula

import (
	"bytes"
	"net/url"

	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/yuin/goldmark"
)

// ItemTypeMarkdown describes a what-you-see-is-what-you-get content editor
const ItemTypeMarkdown = "Markdown"

type Markdown struct{}

// Init sets up an empty "html" property
func (w Markdown) Init(container *Container, id int) {
	(*container)[id].Set("markdown", "")
	(*container)[id].Set("html", "")
}

func (w Markdown) View(b *html.Builder, container *Container, id int) {
	item := container.GetItem(id)
	result := item.GetString("html")
	b.WriteString(result)
}

func (w Markdown) Edit(b *html.Builder, container *Container, id int, endpoint string) {
	item := container.GetItem(id)
	result := item.GetString("markdown")
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

	b.Textarea("markdown").Class("code").Script("install Autosize install Autosave").InnerHTML(result).Close()
	b.CloseAll()
}

// Validate cleans the container for invalid content
func (w Markdown) Validate(container *Container, index int) {

	var buffer bytes.Buffer

	// Convert Markdown to HTML
	originalMarkdown := (*container)[index].GetBytes("markdown")

	if err := goldmark.Convert(originalMarkdown, &buffer); err != nil {
		derp.Wrap(err, "content.Markdown.Validate", "Failed to convert markdown to HTML")
	}

	// Sanitize HTML and save
	cleanHTML := htmlPolicy.Sanitize(buffer.String())
	(*container)[index].Data.SetString("html", cleanHTML)
}

func (w Markdown) Prop(container *Container, id int, params url.Values) (string, error) {
	return "", derp.NewNotFoundError("content.Markdown.Prop", "Unrecognized panel", params)
}
