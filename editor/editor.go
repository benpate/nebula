package editor

import (
	"github.com/benpate/content"
	"github.com/benpate/html"
)

type Editor struct {
	Endpoint string
}

func New(endpoint string) Editor {
	return Editor{
		Endpoint: endpoint,
	}
}

type widgetFunc func(*html.Builder, content.Content, int)

// Draw returns an HTML string containing the VIEW version of the content
func (e Editor) Draw(c content.Content) string {
	builder := html.New()
	widgetFunc := e.getWidget(builder, c[0])
	widgetFunc(builder, c, 0)
	return builder.String()
}

// ItemTypes implements the content.EditorWidget interface, and returns
// a slice of all Item Types for users to select.
func (e Editor) ItemTypes() []content.ItemType {
	return []content.ItemType{
		{
			Code:  content.ItemTypeWYSIWYG,
			Label: "Rich Text Content",
		},
		{
			Code:  content.ItemTypeText,
			Label: "Plain Text Content",
		},
		{
			Code:  content.ItemTypeOEmbed,
			Label: "Embed Image or Video",
		},
		{
			Code:  content.ItemTypeTabs,
			Label: "Tabs",
		},
	}
}

func (e Editor) getWidget(b *html.Builder, item content.Item) widgetFunc {

	switch item.Type {

	case content.ItemTypeContainer:
		return e.Container
	case content.ItemTypeHTML:
		return e.HTML
	case content.ItemTypeOEmbed:
		return e.OEmbed
	case content.ItemTypeTabs:
		return e.Tabs
	case content.ItemTypeText:
		return e.Text
	case content.ItemTypeWYSIWYG:
		return e.WYSIWYG
	default:
		return e.Nil
	}
}

// subTree safely renders a sub-widget.
func (e Editor) subTree(b *html.Builder, c content.Content, id int) {
	subBuilder := b.SubTree()
	widgetFunc := e.getWidget(subBuilder, c[id])

	widgetFunc(subBuilder, c, id)
	subBuilder.CloseAll()
}
