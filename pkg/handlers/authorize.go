package handlers

import (
	"net/http"

	"github.com/osemisan-authorization-server/pkg/clients"
)

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("client_id")
	idx := clients.OsemisanClients.Find(id)
	if idx == -1  {
	}
}
