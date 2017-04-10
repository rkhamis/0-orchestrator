package node

import (
	"net/http"
	"strconv"
	"syscall"

	client "github.com/g8os/go-client"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// KillNodeProcess is the handler for DELETE /nodes/{nodeid}/processes/{processid}
// Kill Process
func (api NodeAPI) KillNodeProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cl, err := tools.GetConnection(r, api)
	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	pID, err := strconv.ParseUint(vars["processid"], 10, 64)
	if err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	processID := client.ProcessId(pID)
	core := client.Core(cl)
	if err := core.KillProcess(processID, syscall.SIGTERM); err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
