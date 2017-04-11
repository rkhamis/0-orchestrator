package storagecluster

import (
	"encoding/json"
	"net/http"
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

}
