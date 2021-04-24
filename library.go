package content

import (
	"github.com/benpate/html"
)

type Widget func(*Library, *html.Builder, *Item)

type Library struct {
	widgets map[string]Widget
}

func NewLibrary() Library {
	return Library{
		widgets: make(map[string]Widget),
	}
}

// ViewerLibrary generates a fully populated library
// containing all of the default controls.
func ViewerLibrary() Library {
	result := NewLibrary()

	result.Register(ItemTypeHTML, HTMLViewer)
	result.Register(ItemTypeText, TextViewer)
	result.Register(ItemTypeTabs, TabsViewer)

	return result
}

func CreatorLibrary() Library {
	result := NewLibrary()

	return result
}

func EditorLibrary() Library {
	result := NewLibrary()

	return result
}

///////////////////////////////
// Library Methods

func (library *Library) Register(class string, widget Widget) *Library {
	library.widgets[class] = widget
	return library
}

// Render returns the HTML for a specific content.Item, based on the RenderType requested
func (library *Library) Render(item *Item) string {

	builder := html.New()

	if widget, ok := library.widgets[item.Type]; ok {
		widget(library, builder, item)
	}

	return builder.String()
}

// RenderToBuilder uses the widget library to safely append values to an existing html.Builder
func (library *Library) SubTree(builder *html.Builder, item *Item) {
	if widget, ok := library.widgets[item.Type]; ok {
		subTree := builder.SubTree()
		widget(library, subTree, item)
		subTree.CloseAll()
	}
}
