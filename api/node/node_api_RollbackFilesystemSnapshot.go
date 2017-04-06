package node

import (
	"net/http"
)

// RollbackFilesystemSnapshot is the handler for POST /nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshot/{snapshotname}/rollback
// Rollback the filesystem to the state at the moment the snapshot was taken
func (api NodeAPI) RollbackFilesystemSnapshot(w http.ResponseWriter, r *http.Request) {
}
