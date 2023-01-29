package handlers

import (
	"net/http"

	"github.com/osemisan-authorization-server/pkg/kvs"
	"github.com/osemisan-authorization-server/pkg/random"
	"github.com/osemisan-authorization-server/pkg/templates"
	"github.com/osemisan-authorization-server/pkg/url"
)

func ApproveHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqid := r.PostFormValue("reqid")
	query := kvs.RequestsKVS[reqid]
	delete(kvs.RequestsKVS, reqid)

	if query == nil {
		w.WriteHeader(http.StatusBadRequest)
		templates.Render("error", w, map[string]string{
			"message": "No matching authorization request",
		})
		return
	}

	approve := r.PostFormValue("approve")

	if approve != "" {
		resType := query.Get("response_type")
		if resType == "code" {
			code := random.GenStr(8)
			kvs.CodesKVS[code] = query

			redirectURI := query.Get("redirect_uri")
			state := query.Get("state")

			u, err := url.BuildURL(redirectURI, map[string]string{
				"code": code,
				"state": state,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, u, http.StatusPermanentRedirect)
		}
	} else {
		redirectURI := query.Get("redirect_uri")
		u, err := url.BuildURL(redirectURI, map[string]string{
			"error": "unsupported_response_type",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, u, http.StatusPermanentRedirect)
	}
}
