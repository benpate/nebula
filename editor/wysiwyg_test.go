package editor

import (
	"encoding/json"
	"testing"

	"github.com/benpate/content"
	"github.com/stretchr/testify/require"
)

func TestWYSIWYG(t *testing.T) {

	// Unmarshal JSON into a content object
	var c content.Content
	s := []byte(`[{
		"type":"WYSIWYG",
		"check": "123456789101112",
		"data":{
			"html":"This is some <i>HTML</i>"
		}}]`)

	require.Nil(t, json.Unmarshal(s, &c))

	// Make a  new editor widget
	editor := New("/endpoint")
	html := editor.Draw(c)

	// Tes results
	expected := `<form method="post" action="/my-url" data-script="install wysiwyg"><input type="hidden" name="type" value="update-item"><input type="hidden" name="itemId" value="0"><input type="hidden" name="check" value="123456789101112"><input type="hidden" name="html"><div class="ck-editor">This is some <i>HTML</i></div></form>`
	require.Equal(t, expected, html)

}
