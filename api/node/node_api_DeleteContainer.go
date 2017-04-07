package node

import (
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
	"net/http"
)

// DeleteContainer is the handler for DELETE /nodes/{nodeid}/containers/{containerid}
// Delete Container instance
func (api NodeAPI) DeleteContainer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["containerid"]
	res, err := api.AysAPI.Ays.DeleteServiceByName(id, "container", api.AysRepo, nil, nil)
	if !tools.HandleAYSResponse(err, res, w, "deleting service") {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
