package node

import (
	"net/http"
)

// KillContainerProcess is the handler for DELETE /nodes/{nodeid}/container/{containerid}/process/{proccessid}
// Kill Process
func (api NodeAPI) KillContainerProcess(w http.ResponseWriter, r *http.Request) {
}
