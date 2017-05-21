package node

import (
	"net/http"
)

// StopGateway is the handler for POST /nodes/{nodeid}/gws/{gwname}/stop
// Stop gateway instance
func (api NodeAPI) StopGateway(w http.ResponseWriter, r *http.Request) {
}
