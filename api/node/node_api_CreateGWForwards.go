package node

import (
	"encoding/json"
	"net/http"
)

// nodeidgwsgwnamefirewallforwardsPost is the handler for POST /nodes/{nodeid}/gws/{gwname}/firewall/forwards
// Create a new Portforwarding
func (api NodeAPI) nodeidgwsgwnamefirewallforwardsPost(w http.ResponseWriter, r *http.Request) {
	var reqBody PortForward

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

}
