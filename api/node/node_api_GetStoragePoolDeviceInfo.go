package node

import (
	"encoding/json"
	"net/http"
)

// GetStoragePoolDeviceInfo is the handler for GET /nodes/{nodeid}/storagepools/{storagepoolname}/device/{deviceuuid}
// Get information of the device
func (api NodeAPI) GetStoragePoolDeviceInfo(w http.ResponseWriter, r *http.Request) {
	var respBody StoragePoolDevice
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
