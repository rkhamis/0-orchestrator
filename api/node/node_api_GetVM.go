package node

import (
	"encoding/json"
	"net/http"
)

// GetVM is the handler for GET /nodes/{nodeid}/vm/{vmid}
// Get detailed virtual machine object
func (api NodeAPI) GetVM(w http.ResponseWriter, r *http.Request) {
	var respBody VM
	json.NewEncoder(w).Encode(&respBody)
}
