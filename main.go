package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/osemisan-authorization-server/pkg/handlers"
)

func main() {
	l := httplog.NewLogger("osemisan-authorization-server", httplog.Options{
		JSON: true,
	})
	r := chi.NewRouter()

	r.Use(httplog.RequestLogger(l))

	r.Get("/", handlers.RootHandler)
	r.Get("/authorize", handlers.AuthorizeGetHandler)

	http.ListenAndServe(":9001", r)
}
