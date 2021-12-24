package transaction

import (
	"testing"

	"github.com/benpate/content"
	"github.com/benpate/datatype"
)

func getTestContent() content.Content {

	return content.Content{
		{
			Type:  "CONTAINER",
			Refs:  []int{1, 2, 3, 4},
			Check: "home",
			Data: datatype.Map{
				"style": "ROWS",
			},
		},
		{
			Type:  "WYSIWYG",
			Check: "secret1",
			Data: datatype.Map{
				"html": "This is the <b>html</b>",
			},
		},
		{
			Type:  "WYSIWYG",
			Check: "secret2",
			Data: datatype.Map{
				"html": "This is the second WYSIWYG section",
			},
		},
		{
			Type:  "WYSIWYG",
			Check: "secret3",
			Data: datatype.Map{
				"html": "This is the third.",
			},
		},
		{
			Type:  "WYSIWYG",
			Check: "secret4",
			Data: datatype.Map{
				"html": "You guessed it.  Fourth section here.",
			},
		},
	}
}

func TestCompact(t *testing.T) {
	c := getTestContent()
	c.Compact()
}