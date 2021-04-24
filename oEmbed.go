package content

import "github.com/benpate/html"

const ItemTypeOEmbed = "OEMBED"

func OEmbedViewer(library *Library, builder *html.Builder, item *Item) {

	// If the oEmbed data includes HTML, then just use that and be done.
	if html := item.GetString("html"); html != "" {
		builder.WriteString(html)
		return
	}

	// Special handling for known types
	switch item.GetString("type") {

	case "photo":
		builder.Empty("img").
			Attr("src", item.GetString("url")).
			Attr("width", item.GetString("width")).
			Attr("height", item.GetString("height")).
			Close()
	}
}
