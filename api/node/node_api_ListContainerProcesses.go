package node

import (
	"encoding/json"
	"net/http"

	client "github.com/g8os/go-client"
	"github.com/g8os/grid/api/tools"
)

// ListContainerProcesses is the handler for GET /nodes/{nodeid}/containers/{containerid}/process
// Get running processes in this container
func (api NodeAPI) ListContainerProcesses(w http.ResponseWriter, r *http.Request) {
	var respBody []Process

	conn, err := tools.GetContainerConnection(r, api)
	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	core := client.Core(conn)
	clprocesses, err := core.Processes()
	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	for _, clprocess := range clprocesses {
		pr := Process{
			Cmd: clprocess.Command,
			// Cpu : # TODO:,
			Pid:  uint64(clprocess.PID),
			Rss:  clprocess.RSS,
			Swap: clprocess.Swap,
			Vms:  clprocess.VMS,
		}
		respBody = append(respBody, pr)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
