package nebula

func All(library *Library) {
	library.Register(ItemTypeLayout, Layout{library: library})
	library.Register(ItemTypeHTML, HTML{})
	library.Register(ItemTypeOEmbed, OEmbed{})
	library.Register(ItemTypeTabs, Tabs{library: library})
	library.Register(ItemTypeText, Text{})
	library.Register(ItemTypeWYSIWYG, WYSIWYG{})
}
