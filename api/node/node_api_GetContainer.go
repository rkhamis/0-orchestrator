package node

import (
	"encoding/json"
	"net/http"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// GetContainer is the handler for GET /nodes/{nodeid}/container/{containerid}
// Get Container
func (api NodeAPI) GetContainer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerID := vars["containerid"]
	service, res, err := api.AysAPI.Ays.GetServiceByName(containerID, "container", api.AysRepo, nil, nil)

	if !tools.HandleAYSResponse(err, res, w, "Getting container service") {
		return
	}

	containerItem := struct {
		Nics           []ContainerNIC
		Filesystems    []string
		Flist          string
		HostNetworking bool
		Hostname       string
		Id             int
		Initprocesses  []CoreSystem
		Ports          []string
		Status         EnumContainerStatus
		Zerotier       string
	}{}

	if err := json.Unmarshal(service.Data, &containerItem); err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	respBody := Container{
		Nics:           containerItem.Nics,
		Zerotier:       containerItem.Zerotier,
		Status:         containerItem.Status,
		Ports:          containerItem.Ports,
		Initprocesses:  containerItem.Initprocesses,
		ContainerId:    containerItem.Id,
		Hostname:       containerItem.Hostname,
		HostNetworking: containerItem.HostNetworking,
		Flist:          containerItem.Flist,
		Filesystems:    containerItem.Filesystems,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
