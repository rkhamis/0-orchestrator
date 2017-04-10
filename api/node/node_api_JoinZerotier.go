package node

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/g8os/grid/api/tools"
	"github.com/gorilla/mux"
)

// JoinZerotier is the handler for POST /nodes/{nodeid}/zerotiers
// Join Zerotier network
func (api NodeAPI) JoinZerotier(w http.ResponseWriter, r *http.Request) {
	var reqBody ZerotierJoin

	nodeID := mux.Vars(r)["nodeid"]

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

	// Create join blueprint
	bp := struct {
		Nwid string `json:"nwid"`
	}{
		Nwid: reqBody.Nwid,
	}

	obj := make(map[string]interface{})
	obj[fmt.Sprintf("zerotier__%s_%s", nodeID, reqBody.Nwid)] = bp
	obj["actions"] = []map[string]string{map[string]string{"action": "join"}}

	if _, err := tools.ExecuteBlueprint(api.AysRepo, "zerotier", reqBody.Nwid, "join", obj); err != nil {
		log.Errorf("error executing blueprint for zerotiers %s join : %+v", reqBody.Nwid, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/nodes/%s/zerotiers/%s", nodeID, reqBody.Nwid))
	w.WriteHeader(http.StatusCreated)
}
