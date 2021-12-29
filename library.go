package nebula

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

func (lib Library) Init(container Container, id int) {
	item := container.GetItem(id)
	widget := lib.Widget(item.Type)

	if initer, ok := widget.(WidgetIniter); ok {
		initer.Init(container, id)
	}
}

// View safely renders a widget's View method (including any sub-widgets)
func (lib Library) View(builder *html.Builder, container Container, id int) {

	subBuilder := builder.SubTree()
	item := container.GetItem(id)
	widget := lib.Widget(item.Type)

	// Render the sub-widget using a sub-builder...
	widget.View(subBuilder, container, id)
	subBuilder.CloseAll()
}

// Edit safely renders a widget's Edit method (including any sub-widgets)
func (lib Library) Edit(builder *html.Builder, container Container, id int, endpoint string) {

	subBuilder := builder.SubTree()
	item := container.GetItem(id)
	widget := lib.Widget(item.Type)

	// Render the sub-widget using a sub-builder...
	widget.Edit(subBuilder, container, id, endpoint)
	subBuilder.CloseAll()
}

// Prop safely renders a widget's Prop method (including any sub-widgets)
func (lib Library) Prop(builder *html.Builder, container Container, id int, params url.Values, endpoint string) error {

	item := container.GetItem(id)
	widget := lib.Widget(item.Type)

	// Render the sub-widget using a sub-builder...
	if propertyEditor, ok := widget.(PropertyEditor); ok {
		return propertyEditor.Prop(builder, container, id, params, endpoint)
	}

	return derp.New(derp.CodeBadRequestError, "container.Library.Prop", "Widget does not support property panels", item)
}
