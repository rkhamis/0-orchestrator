package node

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// DeleteStoragePool is the handler for DELETE /node/{nodeid}/storagepool/{storagepoolname}
// Delete the storage pool
func (api NodeAPI) DeleteStoragePool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["storagepoolname"]

	_, err := api.AysAPI.Ays.DeleteServiceByName(name, "storagepool", api.AysRepo, nil, nil)
	if err != nil {
		log.Errorf("Error deleting storagepool services : %+v", err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
