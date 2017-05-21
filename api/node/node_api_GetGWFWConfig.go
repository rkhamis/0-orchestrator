package node

import (
	"net/http"
)

// GetGWFWConfig is the handler for GET /nodes/{nodeid}/gws/{gwname}/advanced/firewall
// Get current FW config
// Once used you can not use gw.portforwards any longer
func (api NodeAPI) GetGWFWConfig(w http.ResponseWriter, r *http.Request) {
}
