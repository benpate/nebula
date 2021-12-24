package viewer

import (
	"encoding/json"
	"testing"

	"github.com/benpate/content"
	"github.com/stretchr/testify/require"
)

func TestOEmbedImage(t *testing.T) {
	var c content.Content

	text := []byte(`[{
	   		"type": "OEMBED",
	   		"data": {
	   			"type": "photo",
	   			"url": "/image.png",
	   			"height":90,
	   			"width":160
	   		}}]`)

	err := json.Unmarshal(text, &c)

	require.Nil(t, err)

	viewer := New()

	html := viewer.Draw(c)

	require.Equal(t, `<img src="/image.png" width="160" height="90">`, html)
}

func TestOEmbedVideo(t *testing.T) {

	var c content.Content

	text := []byte(`[{
	   		"type": "OEMBED",
	   		"data": {
	   			"type": "video",
	   			"html": "Here's where the video html should go"
	   		}}]`)

	err := json.Unmarshal(text, &c)

	require.Nil(t, err)

	viewer := New()

	html := viewer.Draw(c)

	require.Equal(t, "Here's where the video html should go", html)
}
