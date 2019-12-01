package views

import (
	"Domainf/logic"
	"net/http"
)

// View that return the servers information
func GetServers(w  http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	response := logic.GetServerInformation(host)
	w.Write(response)
}

// View that return the history of all servers
func GetHistory(w http.ResponseWriter, r *http.Request) {
	response := logic.GetServersHistory()
	w.Write(response)
}