package vocabulary

type ItemType struct {
	Code        string
	Label       string
	Description string
}

// ItemTypes returns all item types provided by this widget library
func ItemTypes() []ItemType {
	return []ItemType{
		{
			Code:  ItemTypeWYSIWYG,
			Label: "Rich Text Content",
		},
		{
			Code:  ItemTypeText,
			Label: "Plain Text Content",
		},
		{
			Code:  ItemTypeOEmbed,
			Label: "Embed Image or Video",
		},
		{
			Code:  ItemTypeTabs,
			Label: "Tabs",
		},
	}
}
