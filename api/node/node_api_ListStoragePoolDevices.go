package node

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	log "github.com/Sirupsen/logrus"
	g8client "github.com/g8os/go-client"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// ListStoragePoolDevices is the handler for GET /nodes/{nodeid}/storagepools/{storagepoolname}/devices
// Lists the devices in the storage pool
func (api NodeAPI) ListStoragePoolDevices(w http.ResponseWriter, r *http.Request) {
	// TODO: Device Status is missing
	var respBody []StoragePoolDevice

	vars := mux.Vars(r)
	storagepoolname := vars["storagepoolname"]
	nodeid := vars["nodeid"]

	devices, err := api.getStoragePoolDevices(nodeid, storagepoolname)
	if err != nil {
		log.Errorf("Error Listing storage pool devices: %+v", err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	cl, err := tools.GetConnection(r, api)
	if err != nil {
		log.Errorf("Error Listing storage pool devices: %+v", err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	btrfsMgr := g8client.Btrfs(cl)
	fss, err := btrfsMgr.List()
	if err != nil {
		log.Errorf("Error Listing storage pool devices: %+v", err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	for _, fs := range fss {
		for _, device := range fs.Devices {
			if containsStrings(devices, device.Path) {
				respBody = append(respBody, StoragePoolDevice{Uuid: fs.UUID})
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}

// Get storagepool devices
func (api NodeAPI) getStoragePoolDevices(node, storagepool string) ([]string, error) {
	queryParams := map[string]interface{}{"parent": fmt.Sprintf("node.g8os!%s", node)}

	service, _, err := api.AysAPI.Ays.GetServiceByName(storagepool, "storagepool", api.AysRepo, nil, queryParams)
	if err != nil {
		return nil, err
	}

	var data struct {
		Devices []string `json:"devices"`
	}

	if err := json.Unmarshal(service.Data, &data); err != nil {
		log.Errorf("Error Unmarshal storagepool service '%s' data: %+v", storagepool, err)
		return nil, err
	}

	return data.Devices, nil
}

func containsStrings(slice []string, target string) bool {
	sort.Strings(slice)
	i := sort.SearchStrings(slice, target)
	if i < len(slice) && slice[i] == target {
		return true
	}
	return false
}
