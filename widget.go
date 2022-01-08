package nebula

import (
	"github.com/benpate/html"
)

type Widget interface {

	// View writes the widget HTML into the provided HTML builder
	View(*html.Builder, *Container, int)

	// Edit writes an editable widget into the provided HTML builder
	Edit(*html.Builder, *Container, int, string)
}

type WidgetIniter interface {

	// Init is called on every new Item in the container data.
	// It allows widgets to customize their internal data before use.
	Init(*Container, int)
}
