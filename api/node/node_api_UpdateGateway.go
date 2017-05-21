package node

import (
	"encoding/json"
	"net/http"
)

// UpdateGateway is the handler for PUT /nodes/{nodeid}/gws/{gwname}
// Update Gateway
func (api NodeAPI) UpdateGateway(w http.ResponseWriter, r *http.Request) {
	var reqBody GW

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

}
