package nebula

import (
	"math/rand"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/list"
)

// ItemTypeOEmbed describes an oEmbed object (see https://oembed.com)
const ItemTypeOEmbed = "OEMBED"

type OEmbed struct{}

func (w OEmbed) View(b *html.Builder, container *Container, itemID int) {

	item := container.GetItem(itemID)

	// If the oEmbed data includes HTML, then just use that and be done.
	if html := item.GetString("html"); html != "" {
		b.WriteString(html)
		return
	}

	switch list.Head(item.GetString("mimeType"), "/") {

	case "video":
		b.Span().InnerHTML("video here...")
		b.Close()

	default:
		b.Empty("img").Class("pure-img").Attr("src", item.GetString("file"))
		b.Close()
	}
}

func (w OEmbed) Edit(b *html.Builder, container *Container, itemID int, endpoint string) {

	item := container.GetItem(itemID)
	idString := convert.String(itemID)
	idStringUnique := "file-" + idString + "-" + convert.String(rand.Int())

	b.Container("label").For(idStringUnique).Class("block")

	b.Form("", "").
		Data("id", idString).
		Data("hx-post", endpoint).
		Data("hx-trigger", "change").
		Data("hx-encoding", "multipart/form-data").
		Script("install DropToUpload").
		Class("uploader")

	b.Input("hidden", "action").Value("upload-file")
	b.Input("hidden", "itemId").Value(convert.String(itemID))
	b.Input("hidden", "check").Value(item.Check)
	b.Input("file", "file").ID(idStringUnique).Attr("accept", "image/*")

	if item.GetString("file") == "" {
		b.Div().Class("space-above", "space-below").InnerHTML("Drag Files Here, or Click To Select").Close()
	} else {
		w.View(b, container, itemID)
	}
	b.CloseAll()
}

// Validate cleans the container for invalid content
func (w OEmbed) Validate(container *Container, index int) {
}
