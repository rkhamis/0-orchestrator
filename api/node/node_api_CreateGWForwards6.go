package node

import (
	"encoding/json"
	"net/http"
)

// nodeidgwsgwnamefirewallforwards6Post is the handler for POST /nodes/{nodeid}/gws/{gwname}/firewall/forwards6
// Create a new Portforwarding
func (api NodeAPI) nodeidgwsgwnamefirewallforwards6Post(w http.ResponseWriter, r *http.Request) {
	var reqBody PortForward6

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

}
