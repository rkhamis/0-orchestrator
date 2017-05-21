package node

import (
	"net/http"
)

// GetGWHTTPConfig is the handler for GET /nodes/{nodeid}/gws/{gwname}/advanced/http
// Get current HTTP config
// Once used you can not use gw.httpproxxies any longer
func (api NodeAPI) GetGWHTTPConfig(w http.ResponseWriter, r *http.Request) {
}
