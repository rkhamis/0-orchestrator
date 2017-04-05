package node

import (
	"fmt"
	"net/http"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"

	log "github.com/Sirupsen/logrus"
)

// DeleteFilesystem is the handler for DELETE /node/{nodeid}/storagepool/{storagepoolname}/filesystem/{filesystemname}
// Delete filesystem
func (api NodeAPI) DeleteFilesystem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["filesystemname"]

	resp, err := api.AysAPI.Ays.DeleteServiceByName(name, "filesystem", api.AysRepo, nil, nil)
	if err != nil {
		log.Errorf("Error deleting filesystem services : %+v", err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if resp.StatusCode != http.StatusNoContent {
		log.Errorf("Error deleting filesystem services : %+v", resp.Status)
		tools.WriteError(w, resp.StatusCode, fmt.Errorf("bad response from AYS: %s", resp.Status))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
