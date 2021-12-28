package vocabulary

import (
	"github.com/benpate/content"
	"github.com/benpate/html"
)

const ItemTypeOEmbed = "OEMBED"

type OEmbed struct{}

func (w OEmbed) View(b *html.Builder, c content.Content, id int) {

	item := c.GetItem(id)

	// If the oEmbed data includes HTML, then just use that and be done.
	if html := item.GetString("html"); html != "" {
		b.WriteString(html)
		return
	}

	// Special handling for known types
	switch item.GetString("type") {

	case "photo":
		b.Empty("img").
			Attr("src", item.GetString("url")).
			Attr("width", item.GetString("width")).
			Attr("height", item.GetString("height"))
	}
}

func (w OEmbed) Edit(b *html.Builder, c content.Content, id int, endpoint string) {
	script := "install Uploader(endpoint:'" + endpoint + "')"
	b.Div().Script(script).EndBracket()
	b.H3().InnerHTML("Drag Images Here To Upload").Close()
	b.Close()
}
