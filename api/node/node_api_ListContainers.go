package node

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// ListContainers is the handler for GET /nodes/{nodeid}/containers
// List running Containers
func (api NodeAPI) ListContainers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nodeID := vars["nodeid"]

	query := make(map[string]interface{})
	query["parent"] = strings.Join([]string{"node.g8os", nodeID}, "!")

	services, res, err := api.AysAPI.Ays.ListServicesByRole("container", api.AysRepo, nil, query)
	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if res.StatusCode != http.StatusOK {
		w.WriteHeader(res.StatusCode)
		return
	}

	type containerItem struct {
		Flist    string                      `json:"flist" validate:"nonzero"`
		Hostname string                      `json:"hostname" validate:"nonzero"`
		Status   EnumContainerListItemStatus `json:"status" validate:"nonzero"`
	}

	var respBody []ContainerListItem
	for _, service := range services {
		srv, res, err := api.AysAPI.Ays.GetServiceByName(service.Name, "container", api.AysRepo, nil, nil)

		if err != nil {
			tools.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		if res.StatusCode != http.StatusOK {
			w.WriteHeader(res.StatusCode)
			return
		}

		var containerItem containerItem
		if err := json.Unmarshal(srv.Data, &containerItem); err != nil {
			tools.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		container := ContainerListItem{
			Id:       service.Name,
			Flist:    containerItem.Flist,
			Hostname: containerItem.Hostname,
			Status:   containerItem.Status,
		}
		respBody = append(respBody, container)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
