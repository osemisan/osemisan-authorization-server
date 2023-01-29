package handlers

import (
	"net/http"

	"github.com/osemisan-authorization-server/pkg/clients"
	"github.com/osemisan-authorization-server/pkg/templates"
)

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("client_id")
	idx := clients.OsemisanClients.Find(id)
	if idx == -1 {
		err := templates.Render("error", w, map[string]any{"message": "Unknown client"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
