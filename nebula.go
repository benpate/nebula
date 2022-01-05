package nebula

import (
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/benpate/html"
)

/*****************************************
 * GLOBAL FUNCTIONS
 *****************************************/

// View returns an HTML string containing the VIEW version of the container
func View(library *Library, container *Container) string {
	builder := html.New()
	library.View(builder, container, 0)
	return builder.String()
}

// Edit returns an HTML string containing the EDIT version of the container
func Edit(library *Library, container *Container, endpoint string) string {
	builder := html.New()

	builder.Div().Class("content-editor").EndBracket()
	library.Edit(builder, container, 0, endpoint)
	builder.CloseAll()

	return builder.String()
}

// Prop returns an editable property form based on the URL params provided.
func Prop(library *Library, container *Container, itemID int, endpoint string, params url.Values) (string, error) {
	builder := html.New()
	err := library.Prop(builder, container, itemID, endpoint, params)
	return builder.String(), err
}

/*****************************************
 * OTHER UTILITY  FUNCTIONS
 *****************************************/

// newChecksum generates a new random checksum for the content
func newChecksum() string {
	seed := time.Now().Unix()
	source := rand.NewSource(seed)
	return strconv.FormatInt(source.Int63(), 36) + strconv.FormatInt(source.Int63(), 36)
}

// makeURL assembles a URL string from a path and a set of parameters
func makeURL(path string, queryParams ...string) string {
	if len(queryParams) == 0 {
		return path
	}
	return path + "?" + strings.Join(queryParams, "&")
}
