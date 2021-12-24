package editor

import (
	"github.com/benpate/content"
	"github.com/benpate/html"
)

func (e Editor) OEmbed(builder *html.Builder, c content.Content, id int) {
	builder.Div().InnerHTML("-- placeholder for oEmbed editor --")
}
