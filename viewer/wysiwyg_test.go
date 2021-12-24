package viewer

import (
	"encoding/json"
	"testing"

	"github.com/benpate/content"
	"github.com/stretchr/testify/require"
)

func TestWYSIWYG(t *testing.T) {

	var c content.Content

	s := []byte(`[{
		"type":"WYSIWYG",
		"check": "123456789101112",
		"data":{
			"html":"This is some <i>HTML</i>"
		}}]`)

	err := json.Unmarshal(s, &c)

	require.Nil(t, err)

	viewer := New()

	html := viewer.Draw(c)
	require.Equal(t, "This is some <i>HTML</i>", html)
}
