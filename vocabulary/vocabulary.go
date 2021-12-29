package vocabulary

import "github.com/benpate/nebula"

func All(library nebula.Library) {
	library.Register("layout", Layout{library: library})
	library.Register("html", HTML{})
	library.Register("oembed", OEmbed{})
	library.Register("tabs", Tabs{library: library})
	library.Register("text", Text{})
	library.Register("wysiwyg", WYSIWYG{})
}
