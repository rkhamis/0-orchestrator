package storagecluster

import (
	"net/http"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"

	log "github.com/Sirupsen/logrus"
)

// KillCluster is the handler for DELETE /storageclusters/{label}
// Kill cluster
func (api StorageclustersAPI) KillCluster(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storage_cluster := vars["label"]

	// execute the delete action
	blueprint := map[string]interface{}{
		"actions": []map[string]string{{
			"action":  "delete",
			"actor":   "storage_cluster",
			"service": storage_cluster,
		}},
	}

	run, err := tools.ExecuteBlueprint(api.AysRepo, "storage_cluster", storage_cluster, "delete", blueprint)
	if err != nil {
		httpErr := err.(tools.HTTPError)
		log.Errorf("Error executing blueprint for storage_cluster deletion : %+v", err.Error())
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

	_, err = api.AysAPI.Ays.DeleteServiceByName(storage_cluster, "storage_cluster", api.AysRepo, nil, nil)

	if err != nil {
		log.Errorf("Error in deleting storage_cluster %s : %+v", storage_cluster, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
