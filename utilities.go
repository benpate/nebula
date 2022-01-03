package vocabulary

import "strings"

// makeURL assembles a URL string from a path and a set of parameters
func makeURL(path string, queryParams ...string) string {
	if len(queryParams) == 0 {
		return path
	}
	return path + "?" + strings.Join(queryParams, "&")
}
