package volume

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// GetVolumeInfo is the handler for GET /volumes/{volumeid}
// Get volume information
func (api VolumesAPI) GetVolumeInfo(w http.ResponseWriter, r *http.Request) {
	volumeID := mux.Vars(r)["volumeid"]

	serv, resp, err := api.AysAPI.Ays.GetServiceByName(volumeID, "volume", api.AysRepo, nil, nil)
	if !tools.HandleAYSResponse(err, resp, w, fmt.Sprintf("getting info about volume %s", volumeID)) {
		return
	}

	var respBody Volume
	if err := json.Unmarshal(serv.Data, &respBody); err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	respBody.ID = serv.Name

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
