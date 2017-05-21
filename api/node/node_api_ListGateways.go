package node

import (
	"encoding/json"
	"net/http"
)

// ListGateways is the handler for GET /nodes/{nodeid}/gws
// List running gateways
func (api NodeAPI) ListGateways(w http.ResponseWriter, r *http.Request) {
	var respBody []GW
	json.NewEncoder(w).Encode(&respBody)
}
