package node

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// RollbackFilesystemSnapshot is the handler for POST /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshot/{snapshotname}/rollback
// Rollback the filesystem to the state at the moment the snapshot was taken
func (api NodeAPI) RollbackFilesystemSnapshot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["snapshotname"]

	// execute the delete action of the snapshot
	blueprint := map[string]interface{}{
		"actions": []map[string]string{{
			"action":  "rollback",
			"actor":   "fssnapshot",
			"service": name,
		}},
	}

	bpName := fmt.Sprintf("%s_rollback_%d", name, time.Now().Unix())

	if _, err := tools.ExecuteBlueprint(api.AysRepo, bpName, blueprint); err != nil {
		httpErr := err.(tools.HTTPError)
		log.Errorf("Error executing blueprint for fssnapshot rollback : %+v", err)
		tools.WriteError(w, httpErr.Resp.StatusCode, httpErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
