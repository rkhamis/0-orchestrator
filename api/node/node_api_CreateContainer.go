package node

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// CreateContainer is the handler for POST /nodes/{nodeid}/container
// Create a new Container
func (api NodeAPI) CreateContainer(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateContainer

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		tools.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		tools.WriteError(w, http.StatusBadRequest, err)
		return
	}

	vars := mux.Vars(r)
	nodeId := vars["nodeid"]

	container := struct {
		Nics           []ContainerNIC            `json:"nics"`
		Filesystems    []string                  `json:"filesystems"`
		Flist          string                    `json:"flist"`
		HostNetworking bool                      `json:"hostNetworking"`
		Hostname       string                    `json:"hostname"`
		Node           string                    `json:"node"`
		InitProcesses  []CoreSystem              `json:"initProcesses"`
		Ports          []string                  `json:"ports"`
		Status         EnumCreateContainerStatus `json:"status"`
		Storage        string                    `json:"storage"`
	}{
		Nics:           reqBody.Nics,
		Filesystems:    reqBody.Filesystems,
		Flist:          reqBody.Flist,
		HostNetworking: reqBody.HostNetworking,
		Hostname:       reqBody.Hostname,
		InitProcesses:  reqBody.InitProcesses,
		Node:           nodeId,
		Ports:          reqBody.Ports,
		Status:         reqBody.Status,
		Storage:        reqBody.Storage,
	}

	obj := make(map[string]interface{})
	obj[fmt.Sprintf("container__%s", reqBody.Id)] = container
	obj["actions"] = []map[string]string{map[string]string{"action": "install"}}

	if _, err := tools.ExecuteBlueprint(api.AysRepo, "container", reqBody.Id, "install", obj); err != nil {
		log.Errorf("error executing blueprint for container %s creation : %+v", reqBody.Id, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/nodes/%s/containers/%s", nodeId, reqBody.Id))
	w.WriteHeader(http.StatusCreated)

}
