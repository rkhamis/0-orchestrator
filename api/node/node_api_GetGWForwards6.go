package node

import (
	"encoding/json"
	"net/http"
)

// nodeidgwsgwnamefirewallforwards6Get is the handler for GET /nodes/{nodeid}/gws/{gwname}/firewall/forwards6
// Get list for IPv6 Forwards
func (api NodeAPI) nodeidgwsgwnamefirewallforwards6Get(w http.ResponseWriter, r *http.Request) {
	var respBody []PortForward6
	json.NewEncoder(w).Encode(&respBody)
}
