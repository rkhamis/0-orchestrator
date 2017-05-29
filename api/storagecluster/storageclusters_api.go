package storagecluster

import (
	ays "github.com/zero-os/0-rest-api/api/ays-client"
	_ "github.com/zero-os/0-rest-api/api/validators"
)

// StorageclusterAPI is API implementation of /storagecluster root endpoint
type StorageclustersAPI struct {
	AysRepo string
	AysAPI  *ays.AtYourServiceAPI
}

func NewStorageClusterAPI(repo string, client *ays.AtYourServiceAPI) StorageclustersAPI {
	return StorageclustersAPI{
		AysRepo: repo,
		AysAPI:  client,
	}
}
