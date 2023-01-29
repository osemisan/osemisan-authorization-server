package handlers

import (
	"net/http"

	"github.com/osemisan-authorization-server/pkg/clients"
	"github.com/osemisan-authorization-server/pkg/kvs"
	"github.com/osemisan-authorization-server/pkg/random"
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

	c := clients.OsemisanClients[idx]

	uri := r.FormValue("redirect_uri")
	if !c.ContainsURI(uri) {
		err := templates.Render("error", w, map[string]string{"message": "Invalid redirect URI"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqid := random.GenStr(8)
	kvs.RequestsKVS[reqid] = r.Form

	err := templates.Render("approve", w, map[string]string{
		"reqid":       reqid,
		"clientId":    c.Id,
		"clientScope": c.Scope,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
