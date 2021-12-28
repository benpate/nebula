package content

import (
	"net/url"

	"github.com/benpate/html"
)

type Widget interface {

	// View writes the widget HTML into the provided HTML builder
	View(*html.Builder, Content, int)

	// Edit writes an editable widget into the provided HTML builder
	Edit(*html.Builder, Content, int, string)
}

type WidgetIniter interface {

	// Init is called on every new Item in the content data.
	// It allows widgets to customize their internal data before use.
	Init(Content, int)
}

// PropertyEditor wraps the Prop method, which allows widgets to
// define custom modals that perform content-specific actions
type PropertyEditor interface {

	// Prop returns HTML for a property editor
	Prop(*html.Builder, Content, int, url.Values, string) error
}
