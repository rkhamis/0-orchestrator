package node

import (
	"net/http"
)

// ShutdownVM is the handler for POST /nodes/{nodeid}/vms/{vmid}/shutdown
// Gracefully shutdown the VM
func (api NodeAPI) ShutdownVM(w http.ResponseWriter, r *http.Request) {
}
