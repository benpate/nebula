package nebula

import "github.com/benpate/rosetta/maps"

type ItemType struct {
	Code        string
	Label       string
	Description string
	Icon        string
	Data        maps.Map
}

// ItemTypes returns all item types provided by this widget library
func ItemTypes() []ItemType {
	return []ItemType{
		{
			Code:  ItemTypeWYSIWYG,
			Label: "HTML Content",
			Icon:  "ti ti-align-left",
		},
		{
			Code:  ItemTypeOEmbed,
			Label: "Image/Video Upload",
			Icon:  "ti ti-photo",
		},
		{
			Code:  ItemTypeTabs,
			Label: "Tabs",
			Icon:  "ti",
		},
		{
			Code:  ItemTypeText,
			Label: "Plain Text Content",
			Icon:  "ti ti-align-left",
		},
		{
			Code:  ItemTypeHTML,
			Label: "HTML Code",
			Icon:  "ti ti-code",
		},
	}
}
