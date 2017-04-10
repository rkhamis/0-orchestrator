package node

import (
	"encoding/json"
	"net/http"
)

// GetContainerProcess is the handler for GET /nodes/{nodeid}/container/{containerid}/process/{proccessid}
// Get process details
func (api NodeAPI) GetContainerProcess(w http.ResponseWriter, r *http.Request) {
	var respBody Process
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
