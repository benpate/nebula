package viewer

import (
	"encoding/json"
	"testing"

	"github.com/benpate/content"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestTabs(t *testing.T) {

	text := []byte(`[
		{
			"type":"TABS",
			"refs":[1,2,3],
			"data": {
				"labels": ["First Tab", "Second Tab", "Third Tab"]
			}
		},{
			"type":"HTML",
			"data": {
				"html":"This is the HTML for the first tab."
			}
		},{
			"type":"TEXT",
			"data": {
				"text":"This is the text for the second tab."
			}
		},{
			"type":"HTML",
			"data": {
				"html":"This is the text for the third tab."
			}
		}]`)

	var tabs content.Content

	err := json.Unmarshal(text, &tabs)
	require.Nil(t, err)

	viewer := New()
	result := viewer.Draw(tabs)

	spew.Dump(result)

}
