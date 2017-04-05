package node

import (
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// DeleteBridge is the handler for DELETE /node/{nodeid}/bridge/{bridgeid}
// Remove bridge
func (api NodeAPI) DeleteBridge(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bridge := vars["bridgeid"]
	_, err := api.AysAPI.Ays.DeleteServiceByName(bridge, "bridge", api.AysRepo, nil, nil)

	if err != nil {
		log.Errorf("Error in deleting bridge %s : %+v", bridge, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
