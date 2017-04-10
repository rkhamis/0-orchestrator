package node

import (
	"encoding/json"
	"net/http"
)

// ListContainerProcesses is the handler for GET /nodes/{nodeid}/containers/{containerid}/process
// Get running processes in this container
func (api NodeAPI) ListContainerProcesses(w http.ResponseWriter, r *http.Request) {
	var respBody []Process
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
