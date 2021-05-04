package https

import "net/http"

var (
	URL               = "http://api.waditu.com"
	hdrContentTypeKey = http.CanonicalHeaderKey("Content-Type")
)
