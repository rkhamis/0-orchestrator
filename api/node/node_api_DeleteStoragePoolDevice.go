package node

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// DeleteStoragePoolDevice is the handler for DELETE /nodes/{nodeid}/storagepools/{storagepoolname}/device/{deviceuuid}
// Removes the device from the storagepool
func (api NodeAPI) DeleteStoragePoolDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	node := vars["nodeid"]
	storagepool := vars["storagepoolname"]
	toDeleteDevice := vars["deviceuuid"]

	devices, err := api.getStoragePoolDevices(node, storagepool)
	if err != nil {
		log.Errorf("Error Listing storage pool devices: %+v", err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&devices); err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// remove device from list of current devices
	for i, device := range devices {
		if device == toDeleteDevice {
			devices = append(devices[:i], devices[i+1:]...)
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

	if _, err := tools.ExecuteBlueprint(api.AysRepo, "storagepool", storagepool, "removeDevices", blueprint); err != nil {
		httpErr := err.(tools.HTTPError)
		log.Errorf("Error executing blueprint for storagepool device deletion : %+v", err.Error())
		tools.WriteError(w, httpErr.Resp.StatusCode, httpErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
