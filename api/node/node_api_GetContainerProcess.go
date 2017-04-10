package node

import (
	"encoding/json"
	"net/http"
	"strconv"

	client "github.com/g8os/go-client"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// GetContainerProcess is the handler for GET /nodes/{nodeid}/container/{containerid}/process/{proccessid}
// Get process details
func (api NodeAPI) GetContainerProcess(w http.ResponseWriter, r *http.Request) {
	var respBody Process
	vars := mux.Vars(r)
	conn, err := tools.GetContainerConnection(r, api)
	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	tmp, err := strconv.ParseUint(vars["processid"], 10, 64)
	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	processID := client.ProcessId(tmp)
	core := client.Core(conn)
	clprocess, err := core.Process(processID)

	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	respBody.Cmd = clprocess.Command
	// respBody.Cpu = #TODO:
	respBody.Pid = uint64(clprocess.PID)
	respBody.Rss = clprocess.RSS
	respBody.Swap = clprocess.Swap
	respBody.Vms = clprocess.VMS

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
