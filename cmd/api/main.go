// package main starts the Cerberus api services
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/Monster-yk/cerberus/internal/httpx"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	r := chi.NewRouter()

	r.Use(httpx.RequestLogger(logger))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/repositories", listRepositories)
		r.Route("/repositories/{id}", func(r chi.Router) {
			r.Get("/", getRepository)
		})
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func listRepositories(w http.ResponseWriter, _ *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, []string{})
}

func getRepository(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	httpx.WriteJSON(w, http.StatusOK, map[string]string{"id": id})
}
