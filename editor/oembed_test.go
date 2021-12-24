package editor

import (
	"encoding/json"
	"testing"

	"github.com/benpate/content"
	"github.com/stretchr/testify/require"
)

func TestOEmbedImage(t *testing.T) {

	// Unmarshal JSON into a content object
	var c content.Content
	text := []byte(`[{
	   		"type": "OEMBED",
	   		"data": {
	   			"type": "photo",
	   			"url": "/image.png",
	   			"height":90,
	   			"width":160
	   		}}]`)

	require.Nil(t, json.Unmarshal(text, &c))

	// Make a new editor widget
	editor := New("/endpoint")
	html := editor.Draw(c)

	// Test results
	require.Equal(t, `<img src="/image.png" width="160" height="90">`, html)
}

func TestOEmbedVideo(t *testing.T) {

	// Unmarshal JSON into a content object
	var c content.Content
	text := []byte(`[{
	   		"type": "OEMBED",
	   		"data": {
	   			"type": "video",
	   			"html": "Here's where the video html should go"
	   		}}]`)

	require.Nil(t, json.Unmarshal(text, &c))

	// Make a new editor widget
	editor := New("/endpoint")
	html := editor.Draw(c)

	// Test results
	require.Equal(t, "Here's where the video html should go", html)
}
