package viewer

import (
	"github.com/benpate/content"
	"github.com/benpate/html"
)

// Nil is a default/empty widget that does nothing.
func (v Viewer) Nil(b *html.Builder, c content.Content, id int) {
	// Nothing to see here
}
