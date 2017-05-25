package node

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/g8os/resourcepool/api/tools"
	"github.com/gorilla/mux"
)

// DeleteDHCPHost is the handler for DELETE /nodes/{nodeid}/gws/{gwname}/dhcp/{interface}/hosts/{macaddress}
// Delete dhcp host
func (api NodeAPI) DeleteDHCPHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gateway := vars["gwname"]
	nodeId := vars["nodeid"]
	nicInterface := vars["interface"]
	macaddress := vars["macaddress"]

	queryParams := map[string]interface{}{
		"parent": fmt.Sprintf("node.g8os!%s", nodeId),
	}

	service, res, err := api.AysAPI.Ays.GetServiceByName(gateway, "gateway", api.AysRepo, nil, queryParams)
	if !tools.HandleAYSResponse(err, res, w, "Getting gateway service") {
		return
	}

	var data CreateGWBP
	if err := json.Unmarshal(service.Data, &data); err != nil {
		errMessage := fmt.Errorf("Error Unmarshal gateway service '%s' data: %+v", gateway, err)
		log.Error(errMessage)
		tools.WriteError(w, http.StatusInternalServerError, errMessage)
		return
	}

	var exists bool
	for i, nic := range data.Nics {
		if nic.Name == nicInterface {
			if nic.Dhcpserver == nil {
				break
			}
			for j, host := range nic.Dhcpserver.Hosts {
				if host.Macaddress == macaddress {
					exists = true
					data.Nics[i].Dhcpserver.Hosts = append(data.Nics[i].Dhcpserver.Hosts[:j],
						data.Nics[i].Dhcpserver.Hosts[j+1:]...)
					break
				}
			}
		}
	}

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	obj := make(map[string]interface{})
	obj[fmt.Sprintf("gateway__%s", gateway)] = data

	if _, err := tools.ExecuteBlueprint(api.AysRepo, "gateway", gateway, "update", obj); err != nil {
		fmt.Errorf("error executing blueprint for gateway %s update : %+v", gateway, err)
		tools.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
