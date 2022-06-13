package nebula

import "github.com/microcosm-cc/bluemonday"

var htmlPolicy *bluemonday.Policy
var textPolicy *bluemonday.Policy

func init() {
	htmlPolicy = bluemonday.UGCPolicy()
	htmlPolicy.AddTargetBlankToFullyQualifiedLinks(true)
	htmlPolicy.RequireNoFollowOnFullyQualifiedLinks(true)
	htmlPolicy.RequireParseableURLs(true)

	textPolicy = bluemonday.StripTagsPolicy()
}
