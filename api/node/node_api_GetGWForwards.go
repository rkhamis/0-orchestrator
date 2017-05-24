package node

import (
	"encoding/json"
	"net/http"
)

// nodeidgwsgwnamefirewallforwardsGet is the handler for GET /nodes/{nodeid}/gws/{gwname}/firewall/forwards
// Get list for IPv4 Forwards
func (api NodeAPI) GetGWForwards(w http.ResponseWriter, r *http.Request) {
	var respBody []PortForward
	json.NewEncoder(w).Encode(&respBody)
}
