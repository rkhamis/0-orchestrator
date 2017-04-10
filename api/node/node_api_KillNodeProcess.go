package node

import (
	"net/http"
)

// KillNodeProcess is the handler for DELETE /nodes/{nodeid}/process/{proccessid}
// Kill Process
func (api NodeAPI) KillNodeProcess(w http.ResponseWriter, r *http.Request) {
}
