package volume

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	ays "github.com/g8os/grid/api/ays-client"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
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

// CreateNewVolume is the handler for POST /volumes
// Create a new volume, can be a copy from an existing volume
func (api VolumesAPI) CreateNewVolume(w http.ResponseWriter, r *http.Request) {
	var reqBody VolumeCreate

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

	// Create the blueprint
	bp := struct {
		Size           int    `json:"size"`
		BlockSize      int    `json:"blocksize"`
		TemplateVolume string `json:"templateVolume,omitempty"`
		ReadOnly       bool   `json:"readOnly"`
		Driver         string `json:"driver"`
		StorageCluster string `json:"storageCluster"`
	}{
		Size:           reqBody.Size,
		BlockSize:      reqBody.Blocksize,
		TemplateVolume: reqBody.Templatevolume,
		ReadOnly:       reqBody.ReadOnly,
		Driver:         string(reqBody.Volumetype),
		StorageCluster: reqBody.Storagecluster,
	}

	vName := fmt.Sprintf("v%v", time.Now().Unix())
	bpName := fmt.Sprintf("volume__%v", vName)

	obj := make(map[string]interface{})
	obj[bpName] = bp
	obj["actions"] = []map[string]string{map[string]string{"action": "install"}}

	// And Execute
	if _, err := tools.ExecuteBlueprint(api.AysRepo, vName, obj); err != nil {
		log.Errorf("error executing blueprint for volume %v creation : %+v", vName, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/volumes/%v", vName))
	w.WriteHeader(http.StatusCreated)
}

// GetVolumeInfo is the handler for GET /volumes/{volumeid}
// Get volume information
func (api VolumesAPI) GetVolumeInfo(w http.ResponseWriter, r *http.Request) {
	volumeId := mux.Vars(r)["volumeid"]

	serv, resp, err := api.AysAPI.Ays.GetServiceByName(volumeId, "volume", api.AysRepo, nil, nil)
	if !tools.HandleAYSResponse(err, resp, w, fmt.Sprintf("getting info about volume %s", volumeId)) {
		return
	}

	var respBody Volume
	if err := json.Unmarshal(serv.Data, &respBody); err != nil {
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	respBody.Id = serv.Name

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}

// DeleteVolume is the handler for DELETE /volumes/{volumeid}
// Delete Volume
func (api VolumesAPI) DeleteVolume(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	volumeId := vars["volumeid"]

	// execute the delete action of the snapshot
	blueprint := map[string]interface{}{
		"actions": []map[string]string{{
			"action":  "delete",
			"actor":   "volume",
			"service": volumeId,
		}},
	}

	blueprintName := fmt.Sprintf("%s_delete_%d", volumeId, time.Now().Unix())

	run, err := tools.ExecuteBlueprint(api.AysRepo, blueprintName, blueprint)
	if err != nil {
		httpErr := err.(tools.HTTPError)
		log.Errorf("Error executing blueprint for volume deletion : %+v", err.Error())
		tools.WriteError(w, httpErr.Resp.StatusCode, httpErr)
		return
	}

	// Wait for the delete job to be finshed before we delete the service
	if err = tools.WaitRunDone(run.Key, api.AysRepo); err != nil {
		httpErr, ok := err.(tools.HTTPError)
		if ok {
			tools.WriteError(w, httpErr.Resp.StatusCode, httpErr)
		} else {
			tools.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	_, err = api.AysAPI.Ays.DeleteServiceByName(volumeId, "volume", api.AysRepo, nil, nil)

	if err != nil {
		log.Errorf("Error in deleting volume %s : %+v", volumeId, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ResizeVolume is the handler for POST /volumes/{volumeid}/resize
// Resize Volume
func (api VolumesAPI) ResizeVolume(w http.ResponseWriter, r *http.Request) {
	var reqBody VolumeResize

	volumeId := mux.Vars(r)["volumeid"]

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

	// Create resize blueprint
	bp := struct {
		Size int `json:"size"`
	}{
		Size: reqBody.NewSize,
	}

	bpName := fmt.Sprintf("volumeresize%sto%vat%v", volumeId, reqBody.NewSize, time.Now().Unix())
	decl := fmt.Sprintf("volume__%v", volumeId)

	obj := make(map[string]interface{})
	obj[decl] = bp
	obj["actions"] = []map[string]string{map[string]string{"action": "resize"}}

	// And execute
	if _, err := tools.ExecuteBlueprint(api.AysRepo, bpName, obj); err != nil {
		log.Errorf("error executing blueprint for volume %s resize : %+v", volumeId, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RollbackVolume is the handler for POST /volumes/{volumeid}/rollback
// Rollback a volume to a previous state
func (api VolumesAPI) RollbackVolume(w http.ResponseWriter, r *http.Request) {
	var reqBody VolumeRollback

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

	// Create rollback blueprint
	// TODO: define rollback
	// volumeId := mux.Vars(r)["volumeid"]
	// bp := struct {
	// 	Epoch int `json:"epoch"`
	// }{
	// 	Epoch: reqBody.Epoch,
	// }
	//
	// bpName := fmt.Sprintf("volumerollback%sfrom%vto%v", volumeId, time.Now().Unix(), reqBody.Epoch)
	// decl := fmt.Sprintf("volume__%v", volumeId)
	//
	// obj := make(map[string]interface{})
	// obj[decl] = bp
	// obj["actions"] = []map[string]string{map[string]string{"action": "rollback"}}
	//
	// // And execute
	// if _, err := tools.ExecuteBlueprint(api.AysRepo, bpName, obj); err != nil {
	// 	log.Errorf("error executing blueprint for volume %s rollback : %+v", volumeId, err)
	// 	tools.WriteError(w, http.StatusInternalServerError, err)
	// 	return
	// }
	//
	// w.WriteHeader(http.StatusNoContent)
}
