package nebula

import (
	"github.com/benpate/html"
)

type Library map[string]Widget

func NewLibrary() Library {
	library := make(Library)

	library.Register(ItemTypeLayout, Layout{library: &library})
	library.Register(ItemTypeHTML, HTML{})
	library.Register(ItemTypeOEmbed, OEmbed{})
	library.Register(ItemTypeTabs, Tabs{library: &library})
	library.Register(ItemTypeText, Text{})
	library.Register(ItemTypeWYSIWYG, WYSIWYG{})
	library.Register(ItemTypeMarkdown, Markdown{})

	return library
}

// Register adds a new named widget to the library
func (library *Library) Register(name string, widget Widget) {
	(*library)[name] = widget
}

// Widget looks up a widget in the library by its name.
// If no matching widget is found, than an empty "Nil" widget is returned
func (library *Library) Widget(name string) Widget {

	if widget, ok := (*library)[name]; ok {
		return widget
	}

	return NilWidget{}
}

func (library *Library) Init(container *Container, id int) {
	item := (*container)[id]
	widget := library.Widget(item.Type)

	if initer, ok := widget.(WidgetIniter); ok {
		initer.Init(container, id)
	}
}

// View safely renders a widget's View method (including any sub-widgets)
func (library *Library) View(builder *html.Builder, container *Container, id int) {

	subBuilder := builder.SubTree()
	item := container.GetItem(id)
	widget := library.Widget(item.Type)

	// Render the sub-widget using a sub-builder...
	widget.View(subBuilder, container, id)
	subBuilder.CloseAll()
}

// Edit safely renders a widget's Edit method (including any sub-widgets)
func (library *Library) Edit(builder *html.Builder, container *Container, id int, endpoint string) {

	subBuilder := builder.SubTree()
	item := container.GetItem(id)
	widget := library.Widget(item.Type)

	// Render the sub-widget using a sub-builder...
	widget.Edit(subBuilder, container, id, endpoint)
	subBuilder.CloseAll()
}

// Validate scans all content in the container, removing invalid content before it is saved.
func (library *Library) Validate(container *Container) {
	for index := range *container {
		item := container.GetItem(index)
		widget := library.Widget(item.Type)
		widget.Validate(container, index)
	}
}
