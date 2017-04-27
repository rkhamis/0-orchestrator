package node

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// ListZerotier is the handler for GET /nodes/{nodeid}/zerotiers
// List running Zerotier networks
func (api NodeAPI) ListZerotier(w http.ResponseWriter, r *http.Request) {
	nodeID := mux.Vars(r)["nodeid"]
	// Only zerotiers with the node from the request as parent
	queryParams := map[string]interface{}{
		"fields": "nwid,status,type",
		"parent": fmt.Sprintf("node.g8os!%s", nodeID),
	}

	services, res, err := api.AysAPI.Ays.ListServicesByRole("zerotier", api.AysRepo, nil, queryParams)
	if !tools.HandleAYSResponse(err, res, w, "listing zerotiers") {
		return
	}

	respBody := make([]ZerotierListItem, len(services))
	for i, serv := range services {
		var data ZerotierListItem
		if err := json.Unmarshal(serv.Data, &data); err != nil {
			tools.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		data.Name = serv.Name
		respBody[i] = data
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
