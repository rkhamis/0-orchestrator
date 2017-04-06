package node

import (
	"net/http"
)

// PauseVM is the handler for POST /nodes/{nodeid}/vms/{vmid}/pause
// Pauses the VM
func (api NodeAPI) PauseVM(w http.ResponseWriter, r *http.Request) {
}
