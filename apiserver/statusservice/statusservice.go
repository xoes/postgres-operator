package statusservice

/*
Copyright 2017-2018 Crunchy Data Solutions, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/crunchydata/postgres-operator/apiserver"
	//msgs "github.com/crunchydata/postgres-operator/apiservermsgs"
	"github.com/gorilla/mux"
	"net/http"
	//"strconv"
)

// StatusHandler ...
// pgo status mycluster
// pgo status --selector=env=research
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Debug("statusservice.StatusHandler %v\n", vars)
	//clustername := vars["name"]

	//selector := r.URL.Query().Get("selector")
	//if selector != "" {
	//log.Debug("selector param was [" + selector + "]")
	//}

	err := apiserver.Authn(apiserver.STATUS_PERM, w, r)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

	//resp := StatusCluster(clustername, selector)
	resp := Status()

	json.NewEncoder(w).Encode(resp)
}
