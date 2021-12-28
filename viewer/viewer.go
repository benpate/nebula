package viewer

import (
	"github.com/benpate/content"
	"github.com/benpate/html"
)

type Viewer struct{}

func New() Viewer {
	return Viewer{}
}

type widgetFunc func(*html.Builder, content.Content, int)

// View returns an HTML string containing the VIEW version of the content
func (v Viewer) Draw(c content.Content) string {
	if len(c) == 0 {
		return ""
	}
	item := c.GetItem(0)
	builder := html.New()
	widgetFunc := v.getWidget(builder, item)
	widgetFunc(builder, c, 0)
	return builder.String()
}

func (v Viewer) getWidget(b *html.Builder, item *content.Item) widgetFunc {

	switch item.Type {

	case ItemTypeContainer:
		return v.Container
	case ItemTypeHTML:
		return v.HTML
	case ItemTypeOEmbed:
		return v.OEmbed
	case ItemTypeTabs:
		return v.Tabs
	case ItemTypeText:
		return v.Text
	case ItemTypeWYSIWYG:
		return v.WYSIWYG
	default:
		return v.Nil
	}
}

// subTree safely renders a sub-widget.
func (v Viewer) subTree(b *html.Builder, c content.Content, id int) {

	item := c.GetItem(id)
	subBuilder := b.SubTree()
	widgetFunc := v.getWidget(subBuilder, item)

	// Render the sub-widget using a sub-builder...
	widgetFunc(subBuilder, c, id)
	subBuilder.CloseAll()
}
