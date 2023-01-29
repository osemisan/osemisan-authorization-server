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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/authorize", handlers.AuthorizeHandler)

	http.ListenAndServe(":9001", r)
}
