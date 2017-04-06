package node

import (
	"encoding/json"
	"net/http"
)

// ListZerotier is the handler for GET /nodes/{nodeid}/zerotiers
// List running Zerotier networks
func (api NodeAPI) ListZerotier(w http.ResponseWriter, r *http.Request) {
	var respBody []ZerotierListItem
	json.NewEncoder(w).Encode(&respBody)
}
