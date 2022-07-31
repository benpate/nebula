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
			Code:        ItemTypeWYSIWYG,
			Label:       "Rich Text Editor",
			Description: "WYSIWYG Text editor for general purpose writing.",
			Icon:        "ti ti-align-left",
		},
		{
			Code:        ItemTypeOEmbed,
			Label:       "Image/Video Upload",
			Description: "Upload an image or video to be embedded in the page.",
			Icon:        "ti ti-photo",
		},
		{
			Code:        ItemTypeTabs,
			Label:       "Tabs",
			Description: "Tabs for organizing content.",
			Icon:        "ti ti-folder",
		},
		{
			Code:        ItemTypeText,
			Label:       "Plain Text Content",
			Description: "Plain text content without any additional formatting.",
			Icon:        "ti ti-align-left",
		},
		{
			Code:        ItemTypeMarkdown,
			Label:       "Markdown",
			Description: "Markdown formatted text",
			Icon:        "ti ti-markdown",
		},
		{
			Code:        ItemTypeHTML,
			Label:       "HTML Code",
			Description: "Raw HTML code to be placed in the page. Note: scripts and stylesheets will not be rendered.",
			Icon:        "ti ti-code",
		},
	}
}
