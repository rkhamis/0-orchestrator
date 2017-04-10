package node

import (
	"encoding/json"
	"net/http"
)

// GetVMInfo is the handler for GET /nodes/{nodeid}/vm/{vmid}/info
// Get statistical information about the virtual machine.
func (api NodeAPI) GetVMInfo(w http.ResponseWriter, r *http.Request) {
	var respBody VMInfo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
