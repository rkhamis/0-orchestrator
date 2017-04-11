package volume

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// ListVolumes is the handler for GET /volumes
// Get volume information
func (api VolumesAPI) ListVolumes(w http.ResponseWriter, r *http.Request) {
	volumeID := mux.Vars(r)["volumeid"]

	services, resp, err := api.AysAPI.Ays.ListServicesByRole("volume", api.AysRepo, nil, nil)
	if !tools.HandleAYSResponse(err, resp, w, fmt.Sprintf("Listing volume services %s", volumeID)) {
		return
	}

	var respBody = make([]VolumeListItem, len(services))
	for _, service := range services {
		var volumeInfo VolumeListItem
		if err := json.Unmarshal(service.Data, &volumeInfo); err != nil {
			tools.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		volumeInfo.ID = service.Name
		respBody = append(respBody, volumeInfo)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
