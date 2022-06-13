package nebula

import (
	"github.com/benpate/html"
)

// NilWidget is a default/empty widget that does nothing.
type NilWidget struct{}

func (w NilWidget) View(b *html.Builder, container *Container, id int) {
	// Nothing to see here
}

func (w NilWidget) Edit(b *html.Builder, container *Container, id int, endpoint string) {
	// Nothing to see here
}

// Validate cleans the container for invalid content
func (w NilWidget) Validate(container *Container, index int) {
	// Nothing to see here
}
