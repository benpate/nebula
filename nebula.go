package nebula

import (
	"math/rand"
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
	builder.Div().Script("install NebulaLayout")
	library.Edit(builder, container, 0, endpoint)
	return builder.String()
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
