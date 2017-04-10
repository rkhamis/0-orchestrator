package node

import (
	"encoding/json"
	"net/http"
)

// GetNodeProcess is the handler for GET /nodes/{nodeid}/process/{proccessid}
// Get process details
func (api NodeAPI) GetNodeProcess(w http.ResponseWriter, r *http.Request) {
	var respBody Process
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
