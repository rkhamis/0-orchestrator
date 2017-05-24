package node

import (
	"net/http"
)

// nodeidgwsgwnamefirewallforwardsforwardidDelete is the handler for DELETE /nodes/{nodeid}/gws/{gwname}/firewall/forwards/{forwardid}
// Delete portforward, forwardid = srcip:srcport
func (api NodeAPI) DeleteGWForward(w http.ResponseWriter, r *http.Request) {
}
