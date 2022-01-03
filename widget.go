package nebula

import (
	"net/url"

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

// PropertyEditor wraps the Prop method, which allows widgets to
// define custom modals that perform container-specific actions
type PropertyEditor interface {

	// Prop returns HTML for a property editor
	Prop(*html.Builder, *Container, int, string, url.Values) error
}
