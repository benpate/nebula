package vocabulary

import (
	"github.com/benpate/content"
)

func All(library content.Library) {
	library.Register("container", Container{library: library})
	library.Register("html", HTML{})
	library.Register("oembed", OEmbed{})
	library.Register("tabs", Tabs{library: library})
	library.Register("text", Text{})
	library.Register("wysiwyg", WYSIWYG{})
}
