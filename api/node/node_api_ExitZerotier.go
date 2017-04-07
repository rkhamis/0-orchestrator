package node

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// ExitZerotier is the handler for DELETE /node/{nodeid}/zerotiers/{zerotierid}
// Exit the Zerotier network
func (api NodeAPI) ExitZerotier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	zerotierID := vars["zerotierid"]

	// execute the exit action of the zerotier
	bp := map[string]interface{}{
		"actions": []map[string]string{{
			"action":  "leave",
			"actor":   "zerotier",
			"service": zerotierID,
		}},
	}

	// And execute
	if _, err := tools.ExecuteBlueprint(api.AysRepo, "zerotier", zerotierID, "exit", bp); err != nil {
		log.Errorf("error executing blueprint for zerotier %s exit : %+v", zerotierID, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
