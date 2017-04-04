package node

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// CreateStoragePool is the handler for POST /node/{nodeid}/storagepool
// Create a new storage pool
func (api NodeAPI) CreateStoragePool(w http.ResponseWriter, r *http.Request) {
	var reqBody StoragePoolCreate
	node := mux.Vars(r)["nodeid"]

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		tools.WriteError(w, http.StatusBadRequest, err)
		return
	}

	bpContent := struct {
		StoragePoolCreate
		node string
	}{
		StoragePoolCreate: reqBody,
		node:              node,
	}

	blueprint := map[string]interface{}{
		fmt.Sprintf("storagepool__%s", node): bpContent,
	}

	blueprintName := fmt.Sprintf("storagepool_%s_create_%d", node, time.Now().Unix())
	tools.ExecuteBlueprint(w, api.AysRepo, blueprintName, blueprint)
}
