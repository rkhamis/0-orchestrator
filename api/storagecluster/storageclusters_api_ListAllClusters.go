package storagecluster

import (
	"encoding/json"
	"net/http"
)

// ListAllClusters is the handler for GET /storageclusters
// List all running clusters
func (api StorageclustersAPI) ListAllClusters(w http.ResponseWriter, r *http.Request) {
	var respBody []string
	json.NewEncoder(w).Encode(&respBody)
}
