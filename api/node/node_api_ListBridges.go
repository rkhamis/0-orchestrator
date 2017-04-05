package node

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"fmt"

	"github.com/g8os/grid/api/tools"
)

// ListBridges is the handler for GET /node/{nodeid}/bridge
// List bridges
func (api NodeAPI) ListBridges(w http.ResponseWriter, r *http.Request) {
	var respBody []Bridge
	json.NewEncoder(w).Encode(&respBody)
	vars := mux.Vars(r)
	nodeid := vars["nodeid"]
	services, resp, err := api.AysAPI.Ays.ListServicesByRole("bridge", api.AysRepo, nil, map[string]interface{}{"parent": fmt.Sprintf("node.g8os!%s", nodeid)})

	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		return
	}

	for _, service := range services {
		srv, resp, err := api.AysAPI.Ays.GetServiceByName(service.Name, service.Role, api.AysRepo, nil, nil)
		if err != nil {
			tools.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(resp.StatusCode)
			log.Errorf("Error in listing bridges: %+v\n", err)
			return
		}

		var bridge Bridge
		if err := json.Unmarshal(srv.Data, &bridge); err != nil {
			tools.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		bridge.Name = srv.Name
		respBody = append(respBody, bridge)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
