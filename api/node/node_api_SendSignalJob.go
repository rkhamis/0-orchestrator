package node

import (
	"encoding/json"
	"net/http"
	"syscall"

	client "github.com/g8os/go-client"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// SendSignalJob is the handler for POST /nodes/{nodeid}/containers/{containerid}/job
// Send signal to the job
func (api NodeAPI) SendSignalJob(w http.ResponseWriter, r *http.Request) {
	var reqBody ProcessSignal

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	containerId := mux.Vars(r)["containerid"]

	// Get a connection
	cl, err := tools.GetConnection(r, api)
	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Send signal to the container
	if err := client.Core(cl).Kill(client.JobId(containerId), syscall.Signal(reqBody.Signal)); err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
