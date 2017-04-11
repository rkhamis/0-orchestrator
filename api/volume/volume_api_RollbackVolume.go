package volume

import (
	"encoding/json"
	"net/http"
)

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
