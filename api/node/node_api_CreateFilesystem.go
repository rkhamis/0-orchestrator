package node

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// CreateFilesystem is the handler for POST /node/{nodeid}/storagepool/{storagepoolname}/filesystem
// Create a new filesystem
func (api NodeAPI) CreateFilesystem(w http.ResponseWriter, r *http.Request) {
	var reqBody FilesystemCreate
	storagepool := mux.Vars(r)["storagepoolname"]

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

	bpContent := struct {
		StoragePool string `json:"storagePool"`
		Name        string `json:"name"`
		Quota       uint32 `json:"quota"`
	}{
		StoragePool: storagepool,
		Name:        reqBody.Name,
		Quota:       reqBody.Quota,
	}

	blueprint := map[string]interface{}{
		fmt.Sprintf("filesystem__%s", reqBody.Name): bpContent,
		"actions": []map[string]string{{"action": "install"}},
	}
	blueprintName := fmt.Sprintf("filesystem__%s_create_%d", storagepool, time.Now().Unix())
	if err := tools.ExecuteBlueprint(api.AysRepo, blueprintName, blueprint); err != nil {
		httpErr := err.(tools.HTTPError)
		log.Errorf("Error executing blueprint for filesystem creation : %+v", err.Error())
		tools.WriteError(w, httpErr.Resp.StatusCode, httpErr)
	}
}
