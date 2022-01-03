package nebula

import (
	"math/rand"
	"net/url"
	"strconv"
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
	library.Edit(builder, container, 0, endpoint)
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

// NewChecksum generates a new random checksum for the content
func NewChecksum() string {
	seed := time.Now().Unix()
	source := rand.NewSource(seed)
	return strconv.FormatInt(source.Int63(), 36) + strconv.FormatInt(source.Int63(), 36)
}
