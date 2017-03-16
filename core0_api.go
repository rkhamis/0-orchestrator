package main

import (
	"encoding/json"
	"net/http"
)

// Core0API is API implementation of /core0 root endpoint
type Core0API struct {
}

// Get is the handler for GET /core0
// List Core0s
func (api Core0API) Get(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")

}

// idGet is the handler for GET /core0/{id}
func (api Core0API) idGet(w http.ResponseWriter, r *http.Request) {
	var respBody Core0
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcorerebootPost is the handler for POST /core0/{id}/core/reboot
// Immediately reboot the machine.
func (api Core0API) idcorerebootPost(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcorepingPost is the handler for POST /core0/{id}/core/ping
// Execute a ping command to this Core0
func (api Core0API) idcorepingPost(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcoresystemPost is the handler for POST /core0/{id}/core/system
// Execute a new process  on this Core0
func (api Core0API) idcoresystemPost(w http.ResponseWriter, r *http.Request) {
	var reqBody CoreSystem

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcorekillPost is the handler for POST /core0/{id}/core/kill
// Kill a process / command
func (api Core0API) idcorekillPost(w http.ResponseWriter, r *http.Request) {
	var reqBody CoreKill

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcorekillallPost is the handler for POST /core0/{id}/core/killall
// Kills all running commands
func (api Core0API) idcorekillallPost(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcorestateGet is the handler for GET /core0/{id}/core/state
// The aggregated consumption of core0 + all processes (cpu, memory, etc...)
func (api Core0API) idcorestateGet(w http.ResponseWriter, r *http.Request) {
	var respBody CoreStateResult
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idkvmdomainGet is the handler for GET /core0/{id}/kvmdomain
// List kvmdomain
func (api Core0API) idkvmdomainGet(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idkvmdomainPost is the handler for POST /core0/{id}/kvmdomain
// Create a new kvmdomain
func (api Core0API) idkvmdomainPost(w http.ResponseWriter, r *http.Request) {
	var reqBody KVMCreate

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idkvmdomaindomainidGet is the handler for GET /core0/{id}/kvmdomain/{domainid}
// Get detailed domain object
func (api Core0API) idkvmdomaindomainidGet(w http.ResponseWriter, r *http.Request) {
	var respBody KVMDomain
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idkvmdomaindomainidDelete is the handler for DELETE /core0/{id}/kvmdomain/{domainid}
// Delete Domain
func (api Core0API) idkvmdomaindomainidDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idkvmdomaindomainidstopPost is the handler for POST /core0/{id}/kvmdomain/{domainid}/stop
// Gracefully stop the kvmdomain
func (api Core0API) idkvmdomaindomainidstopPost(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idkvmdomaindomainiddestroyPost is the handler for POST /core0/{id}/kvmdomain/{domainid}/destroy
// Destroy the kvmdomain
func (api Core0API) idkvmdomaindomainiddestroyPost(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idkvmdomaindomainidpausePost is the handler for POST /core0/{id}/kvmdomain/{domainid}/pause
// Pause the kvmdomain
func (api Core0API) idkvmdomaindomainidpausePost(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idkvmdomaindomainidstartPost is the handler for POST /core0/{id}/kvmdomain/{domainid}/start
// Start the kvmdomain
func (api Core0API) idkvmdomaindomainidstartPost(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idinfonicGet is the handler for GET /core0/{id}/info/nic
func (api Core0API) idinfonicGet(w http.ResponseWriter, r *http.Request) {
	var respBody []NicInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idinfoosGet is the handler for GET /core0/{id}/info/os
func (api Core0API) idinfoosGet(w http.ResponseWriter, r *http.Request) {
	var respBody OSInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idinfocpuGet is the handler for GET /core0/{id}/info/cpu
func (api Core0API) idinfocpuGet(w http.ResponseWriter, r *http.Request) {
	var respBody []CPUInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idinfodiskGet is the handler for GET /core0/{id}/info/disk
func (api Core0API) idinfodiskGet(w http.ResponseWriter, r *http.Request) {
	var respBody []DiskInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idinfomemGet is the handler for GET /core0/{id}/info/mem
func (api Core0API) idinfomemGet(w http.ResponseWriter, r *http.Request) {
	var respBody MemInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idprocessGet is the handler for GET /core0/{id}/process
// Get Processes
func (api Core0API) idprocessGet(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idprocessproccessidGet is the handler for GET /core0/{id}/process/{proccessid}
// Get process details
func (api Core0API) idprocessproccessidGet(w http.ResponseWriter, r *http.Request) {
	var respBody Process
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idprocessproccessidDelete is the handler for DELETE /core0/{id}/process/{proccessid}
// Kill Process
func (api Core0API) idprocessproccessidDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idbridgeGet is the handler for GET /core0/{id}/bridge
// List bridges
func (api Core0API) idbridgeGet(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idbridgePost is the handler for POST /core0/{id}/bridge
// Creates a new bridge
func (api Core0API) idbridgePost(w http.ResponseWriter, r *http.Request) {
	var reqBody BridgeCreate

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idbridgebridgeidGet is the handler for GET /core0/{id}/bridge/{bridgeid}
// Get bridge details
func (api Core0API) idbridgebridgeidGet(w http.ResponseWriter, r *http.Request) {
	var respBody Bridge
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idbridgebridgeidDelete is the handler for DELETE /core0/{id}/bridge/{bridgeid}
// Remove bridge
func (api Core0API) idbridgebridgeidDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskGet is the handler for GET /core0/{id}/disk
// List blockdevices present in the Core0
func (api Core0API) iddiskGet(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskdevicenameOrdiskserialGet is the handler for GET /core0/{id}/disk/{devicenameOrdiskserial}
// Get disk details
func (api Core0API) iddiskdevicenameOrdiskserialGet(w http.ResponseWriter, r *http.Request) {
	var respBody DiskExtInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskdevicenameOrdiskserialmktablePost is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/mktable
// Creates a partition table on a blockdevice
func (api Core0API) iddiskdevicenameOrdiskserialmktablePost(w http.ResponseWriter, r *http.Request) {
	var reqBody DiskMKTable

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskdevicenameOrdiskserialmountPost is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/mount
// Mounts the disk
func (api Core0API) iddiskdevicenameOrdiskserialmountPost(w http.ResponseWriter, r *http.Request) {
	var reqBody DiskMount

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskdevicenameOrdiskserialumountPost is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/umount
// Unmount the disk
func (api Core0API) iddiskdevicenameOrdiskserialumountPost(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskdevicenameOrdiskserialpartitionsGet is the handler for GET /core0/{id}/disk/{devicenameOrdiskserial}/partitions
// Lists partitions
func (api Core0API) iddiskdevicenameOrdiskserialpartitionsGet(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskdevicenameOrdiskserialpartitionsPost is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/partitions
// Create a new partition
func (api Core0API) iddiskdevicenameOrdiskserialpartitionsPost(w http.ResponseWriter, r *http.Request) {
	var reqBody DiskCreatePartition

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskdevicenameOrdiskserialpartitionspartitionidGet is the handler for GET /core0/{id}/disk/{devicenameOrdiskserial}/partitions/{partitionid}
// Gets partition info. (TODO Needs further speccing)
func (api Core0API) iddiskdevicenameOrdiskserialpartitionspartitionidGet(w http.ResponseWriter, r *http.Request) {
	var respBody string
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskdevicenameOrdiskserialpartitionspartitionidDelete is the handler for DELETE /core0/{id}/disk/{devicenameOrdiskserial}/partitions/{partitionid}
// Removes a partition
func (api Core0API) iddiskdevicenameOrdiskserialpartitionspartitionidDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskdevicenameOrdiskserialpartitionspartitionidmountPost is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/partitions/{partitionid}/mount
// Mounts the partition
func (api Core0API) iddiskdevicenameOrdiskserialpartitionspartitionidmountPost(w http.ResponseWriter, r *http.Request) {
	var reqBody DiskMount

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// iddiskdevicenameOrdiskserialpartitionspartitionidumountPost is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/partitions/{partitionid}/umount
// Unmount the partition
func (api Core0API) iddiskdevicenameOrdiskserialpartitionspartitionidumountPost(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcoreXGet is the handler for GET /core0/{id}/coreX
// List running CoreXses
func (api Core0API) idcoreXGet(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcoreXPost is the handler for POST /core0/{id}/coreX
// Create a new CoreX
func (api Core0API) idcoreXPost(w http.ResponseWriter, r *http.Request) {
	var reqBody CoreXCreate

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcoreXcoreXidGet is the handler for GET /core0/{id}/coreX/{coreXid}
// Get CoreX
func (api Core0API) idcoreXcoreXidGet(w http.ResponseWriter, r *http.Request) {
	var respBody CoreX
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcoreXcoreXidDelete is the handler for DELETE /core0/{id}/coreX/{coreXid}
// Delete CoreX instance
func (api Core0API) idcoreXcoreXidDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

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
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcommandGet is the handler for GET /core0/{id}/command
// List running commands
func (api Core0API) idcommandGet(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// idcommandcommandidGet is the handler for GET /core0/{id}/command/{commandid}
func (api Core0API) idcommandcommandidGet(w http.ResponseWriter, r *http.Request) {
	var respBody CommandResult
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}
