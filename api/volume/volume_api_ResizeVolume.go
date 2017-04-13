package volume

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// ResizeVolume is the handler for POST /volumes/{volumeid}/resize
// Resize Volume
func (api VolumesAPI) ResizeVolume(w http.ResponseWriter, r *http.Request) {
	var reqBody VolumeResize

	volumeID := mux.Vars(r)["volumeid"]

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
		Size int `yaml:"size" json:"size"`
	}{
		Size: reqBody.NewSize,
	}

	decl := fmt.Sprintf("volume__%v", volumeID)

	obj := make(map[string]interface{})
	obj[decl] = bp
	obj["actions"] = []map[string]string{map[string]string{"action": "resize"}}

	// And execute
	if _, err := tools.ExecuteBlueprint(api.AysRepo, "volume", volumeID, "resize", obj); err != nil {
		log.Errorf("error executing blueprint for volume %s resize : %+v", volumeID, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
