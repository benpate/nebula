package viewer

import (
	"github.com/benpate/content"
	"github.com/benpate/html"
)

type Viewer struct{}

func New() content.Widget {
	return Viewer{}
}

type widgetFunc func(*html.Builder, content.Content, int)

// View returns an HTML string containing the VIEW version of the content
func (v Viewer) Draw(c content.Content) string {
	item := c.GetItem(0)
	builder := html.New()
	widgetFunc := getWidget(builder, item)
	widgetFunc(builder, c, 0)
	return builder.String()
}

func getWidget(b *html.Builder, item *content.Item) widgetFunc {

	switch item.Type {

	case ItemTypeContainer:
		return Container
	case ItemTypeHTML:
		return HTML
	case ItemTypeOEmbed:
		return OEmbed
	case ItemTypeTabs:
		return Tabs
	case ItemTypeText:
		return Text
	case ItemTypeWYSIWYG:
		return WYSIWYG
	default:
		return Nil
	}
}

// subTree safely renders a sub-widget.
func subTree(b *html.Builder, c content.Content, id int) {
	item := c.GetItem(id)
	subBuilder := b.SubTree()
	widgetFunc := getWidget(subBuilder, item)

	widgetFunc(subBuilder, c, id)
	subBuilder.CloseAll()
}
