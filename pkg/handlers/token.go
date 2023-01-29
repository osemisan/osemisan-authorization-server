package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/osemisan-authorization-server/pkg/clients"
	"github.com/osemisan-authorization-server/pkg/jwt"
	"github.com/osemisan-authorization-server/pkg/kvs"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"Bearer"`
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	id, sec, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idx := clients.OsemisanClients.Find(id)
	if idx == -1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	c := clients.OsemisanClients[idx]

	if c.Secret != sec {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	r.ParseForm()

	g := r.PostFormValue("grant_type")

	if g == "authorization_code" {
		code := r.PostFormValue("code")
		codeValues := kvs.CodesKVS[code]
		if codeValues != nil {
			delete(kvs.CodesKVS, code)
			expectedId := codeValues.Get("client_id")
			if (expectedId == c.Id) {
				token, err := jwt.BuildJWT(c.Scope)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				response := TokenResponse{
					AccessToken: token,
				}
				bytesRes, err := json.Marshal(response)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(bytesRes)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("Client mismatch, expected %s, actual: %s", expectedId, c.Id)))
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Unkown code: %s", code)))
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
