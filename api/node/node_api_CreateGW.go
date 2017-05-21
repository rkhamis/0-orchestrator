package node

import (
	"encoding/json"
	"net/http"
)

// CreateGW is the handler for POST /nodes/{nodeid}/gws
// Create a new gateway
func (api NodeAPI) CreateGW(w http.ResponseWriter, r *http.Request) {
	var reqBody GW

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

}
