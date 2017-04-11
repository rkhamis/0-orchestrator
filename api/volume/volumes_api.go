package volume

import (
	ays "github.com/g8os/grid/api/ays-client"
)

// VolumesAPI is API implementation of /volumes root endpoint
type VolumesAPI struct {
	AysRepo string
	AysAPI  *ays.AtYourServiceAPI
}

func NewVolumeAPI(repo string, client *ays.AtYourServiceAPI) VolumesAPI {
	return VolumesAPI{
		AysRepo: repo,
		AysAPI:  client,
	}
}
