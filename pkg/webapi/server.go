package webapi

import (
	"github.com/gmarcial/travel-route/pkg/travel"
	webapi "github.com/gmarcial/travel-route/pkg/webapi/resources/travel"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func ListenAndServe(dependency *travel.Dependency) {
	router := chi.NewRouter()
	webapi.Routing(router, dependency)
	server := http.Server{
		Addr:              ":5000",
		Handler:           router,
	}

	log.Print(server.ListenAndServe())
}
