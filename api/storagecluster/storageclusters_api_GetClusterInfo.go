package storagecluster

import (
	"encoding/json"
	"net/http"
)

// GetClusterInfo is the handler for GET /storageclusters/{label}
// Get full Information about specific cluster
func (api StorageclustersAPI) GetClusterInfo(w http.ResponseWriter, r *http.Request) {
	var respBody Cluster
	json.NewEncoder(w).Encode(&respBody)
}
