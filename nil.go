package content

import (
	"github.com/benpate/html"
)

// NilWidget is a default/empty widget that does nothing.
type NilWidget struct{}

func (w NilWidget) View(b *html.Builder, c Content, id int) {
	// Nothing to see here
}

func (w NilWidget) Edit(b *html.Builder, c Content, id int, endpoint string) {
	// Nothing to see here
}
