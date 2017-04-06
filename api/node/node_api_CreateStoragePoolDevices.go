package node

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// CreateStoragePoolDevices is the handler for POST /nodes/{nodeid}/storagepools/{storagepoolname}/device
// Add extra devices to this storage pool
func (api NodeAPI) CreateStoragePoolDevices(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node := vars["nodeid"]
	storagepool := vars["storagepoolname"]

	devices, err := api.getStoragePoolDevices(node, storagepool)
	if err != nil {
		log.Errorf("Error Listing storage pool devices: %+v", err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// decode request
	var newDevices []string
	if err := json.NewDecoder(r.Body).Decode(&devices); err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	for _, device := range newDevices {
		if !containsStrings(devices, device) {
			devices = append(devices, device)
		}
	}

	bpContent := struct {
		Devices []string `json:"devices"`
	}{
		Devices: devices,
	}
	blueprint := map[string]interface{}{
		fmt.Sprintf("storagepool__%s", storagepool): bpContent,
	}

	blueprintName := fmt.Sprintf("storagepooldevice__%s_create_%d", node, time.Now().Unix())
	if _, err := tools.ExecuteBlueprint(api.AysRepo, blueprintName, blueprint); err != nil {
		httpErr := err.(tools.HTTPError)
		log.Errorf("Error executing blueprint for storagepool device creation : %+v", err.Error())
		tools.WriteError(w, httpErr.Resp.StatusCode, httpErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
