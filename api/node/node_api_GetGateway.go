package node

import (
	"encoding/json"
	"net/http"
)

// GetGateway is the handler for GET /nodes/{nodeid}/gws/{gwname}
// Get gateway
func (api NodeAPI) GetGateway(w http.ResponseWriter, r *http.Request) {
	var respBody GW
	json.NewEncoder(w).Encode(&respBody)
}
