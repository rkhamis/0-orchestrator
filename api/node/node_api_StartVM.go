package node

import (
	"net/http"
)

// StartVM is the handler for POST /nodes/{nodeid}/vms/{vmid}/start
// Starts the VM
func (api NodeAPI) StartVM(w http.ResponseWriter, r *http.Request) {
}
