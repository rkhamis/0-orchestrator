package node

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// DeleteContainer is the handler for DELETE /nodes/{nodeid}/containers/{containerid}
// Delete Container instance
func (api NodeAPI) DeleteContainer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerID := vars["containerid"]

	// execute the delete action of the snapshot
	bp := map[string]interface{}{
		"actions": []map[string]string{{
			"action":  "delete",
			"actor":   "container",
			"service": containerID,
		}},
	}

	run, err := tools.ExecuteBlueprint(api.AysRepo, "container", containerID, "delete", bp)
	if err != nil {
		httpErr := err.(tools.HTTPError)
		log.Errorf("Error executing blueprint for container deletion : %+v", err.Error())
		tools.WriteError(w, httpErr.Resp.StatusCode, httpErr)
		return
	}

	// Wait for the delete job to be finshed before we delete the service
	if err = tools.WaitRunDone(run.Key, api.AysRepo); err != nil {
		httpErr, ok := err.(tools.HTTPError)
		if ok {
			tools.WriteError(w, httpErr.Resp.StatusCode, httpErr)
		} else {
			tools.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	res, err := api.AysAPI.Ays.DeleteServiceByName(containerID, "container", api.AysRepo, nil, nil)
	if !tools.HandleAYSResponse(err, res, w, "deleting service") {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
