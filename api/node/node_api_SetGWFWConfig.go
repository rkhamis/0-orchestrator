package node

import (
	"net/http"
)

// SetGWFWConfig is the handler for POST /nodes/{nodeid}/gws/{gwname}/advanced/firewall
// Set FW config
func (api NodeAPI) SetGWFWConfig(w http.ResponseWriter, r *http.Request) {
}
