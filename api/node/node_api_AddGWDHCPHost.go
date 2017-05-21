package node

import (
	"encoding/json"
	"net/http"
)

// AddGWDHCPHost is the handler for POST /nodes/{nodeid}/gws/{gwname}/dhcp/{interface}/hosts
// Add a dhcp host to a specified interface
func (api NodeAPI) AddGWDHCPHost(w http.ResponseWriter, r *http.Request) {
	var reqBody GWHost

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

}
