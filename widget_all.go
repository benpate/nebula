package vocabulary

import "github.com/benpate/nebula"

func All(library *nebula.Library) {
	library.Register(nebula.ItemTypeLayout, Layout{library: library})
	library.Register(nebula.ItemTypeHTML, HTML{})
	library.Register(nebula.ItemTypeOEmbed, OEmbed{})
	library.Register(nebula.ItemTypeTabs, Tabs{library: library})
	library.Register(nebula.ItemTypeText, Text{})
	library.Register(nebula.ItemTypeWYSIWYG, WYSIWYG{})
}
