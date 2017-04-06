package node

import (
	"fmt"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// RebootNode is the handler for POST /nodes/{nodeid}/reboot
// Immediately reboot the machine.
func (api NodeAPI) RebootNode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nodeId := vars["nodeid"]

	blueprint := map[string]interface{}{
		"actions": []map[string]string{{
			"action":  "reboot",
			"actor":   "node",
			"service": nodeId,
		}},
	}
	blueprintName := fmt.Sprintf("node__%s_reboot_%d", nodeId, time.Now().Unix())

	run, err := tools.ExecuteBlueprint(api.AysRepo, blueprintName, blueprint)
	if err != nil {
		httpErr := err.(tools.HTTPError)
		tools.WriteError(w, httpErr.Resp.StatusCode, httpErr)
		return
	}

	if err = tools.WaitRunDone(run.Key, api.AysRepo); err != nil {
		httpErr, ok := err.(tools.HTTPError)
		if ok {
			tools.WriteError(w, httpErr.Resp.StatusCode, httpErr)
		} else {
			tools.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
