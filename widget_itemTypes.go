package nebula

import (
	"github.com/benpate/datatype"
)

type ItemType struct {
	Code        string
	Label       string
	Description string
	Icon        string
	Data        datatype.Map
}

// ItemTypes returns all item types provided by this widget library
func ItemTypes() []ItemType {
	return []ItemType{
		{
			Code:  ItemTypeWYSIWYG,
			Label: "HTML Content",
			Icon:  "fa-solid fa-align-left",
		},
		{
			Code:  ItemTypeOEmbed,
			Label: "Image/Video Upload",
			Icon:  "fa-regular fa-image",
		},
		{
			Code:  ItemTypeTabs,
			Label: "Tabs",
			Icon:  "fa-solid",
		},
		{
			Code:  ItemTypeText,
			Label: "Plain Text Content",
			Icon:  "fa-solid fa-align-left",
		},
		{
			Code:  ItemTypeHTML,
			Label: "HTML Code",
			Icon:  "fa-solid fa-code",
		},
	}
}
