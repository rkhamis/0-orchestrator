package storagecluster

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/g8os/grid/api/tools"

	log "github.com/Sirupsen/logrus"
)

// DeployNewCluster is the handler for POST /storageclusters
// Deploy New Cluster
func (api StorageclustersAPI) DeployNewCluster(w http.ResponseWriter, r *http.Request) {
	var reqBody ClusterCreate

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

	blueprint := struct {
		Label    string   `json:"label"`
		NrServer int      `json:"nrServer"`
		HasSlave bool     `json:"hasSlave"`
		DiskType string   `json:"diskType"`
		Nodes    []string `json:"nodes"`
	}{
		Label:    reqBody.Label,
		NrServer: reqBody.Servers,
		HasSlave: reqBody.SlaveNodes,
		DiskType: string(reqBody.DriveType),
		Nodes:    reqBody.Nodes,
	}

	obj := make(map[string]interface{})
	obj[fmt.Sprintf("storage_cluster__%s", reqBody.Label)] = blueprint
	obj["actions"] = []map[string]string{map[string]string{"action": "install"}}

	if _, err := tools.ExecuteBlueprint(api.AysRepo, "storage_cluster", reqBody.Label, "install", obj); err != nil {
		log.Errorf("error executing blueprint for storage_cluster %s creation : %+v", reqBody.Label, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/storageclusters/%s", reqBody.Label))
	w.WriteHeader(http.StatusCreated)
}
