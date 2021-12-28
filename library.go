package content

import (
	"net/url"

	"github.com/benpate/derp"
	"github.com/benpate/html"
)

type Library map[string]Widget

func NewLibrary() Library {
	return make(Library)
}

// Register adds a new named widget to the library
func (lib Library) Register(name string, widget Widget) {
	lib[name] = widget
}

// Widget looks up a widget in the library by its name.
// If no matching widget is found, than an empty "Nil" widget is returned
func (lib Library) Widget(name string) Widget {

	if widget, ok := lib[name]; ok {
		return widget
	}

	return NilWidget{}
}

func (lib Library) Init(content Content, id int) {
	item := content.GetItem(id)
	widget := lib.Widget(item.Type)

	if initer, ok := widget.(WidgetIniter); ok {
		initer.Init(content, id)
	}
}

// View safely renders a widget's View method (including any sub-widgets)
func (lib Library) View(builder *html.Builder, content Content, id int) {

	subBuilder := builder.SubTree()
	item := content.GetItem(id)
	widget := lib.Widget(item.Type)

	// Render the sub-widget using a sub-builder...
	widget.View(subBuilder, content, id)
	subBuilder.CloseAll()
}

// Edit safely renders a widget's Edit method (including any sub-widgets)
func (lib Library) Edit(builder *html.Builder, content Content, id int, endpoint string) {

	subBuilder := builder.SubTree()
	item := content.GetItem(id)
	widget := lib.Widget(item.Type)

	// Render the sub-widget using a sub-builder...
	widget.Edit(subBuilder, content, id, endpoint)
	subBuilder.CloseAll()
}

// Prop safely renders a widget's Prop method (including any sub-widgets)
func (lib Library) Prop(builder *html.Builder, content Content, id int, params url.Values, endpoint string) error {

	item := content.GetItem(id)
	widget := lib.Widget(item.Type)

	// Render the sub-widget using a sub-builder...
	if propertyEditor, ok := widget.(PropertyEditor); ok {
		return propertyEditor.Prop(builder, content, id, params, endpoint)
	}

	return derp.New(derp.CodeBadRequestError, "content.Library.Prop", "Widget does not support property panels", item)
}
