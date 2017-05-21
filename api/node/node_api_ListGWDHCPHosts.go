package node

import (
	"encoding/json"
	"net/http"
)

// ListGWDHCPHosts is the handler for GET /nodes/{nodeid}/gws/{gwname}/dhcp/{interface}/hosts
// List DHCPHosts for specified interface
func (api NodeAPI) ListGWDHCPHosts(w http.ResponseWriter, r *http.Request) {
	var respBody []GWHost
	json.NewEncoder(w).Encode(&respBody)
}
