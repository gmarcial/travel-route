package travel

import (
	"github.com/gmarcial/travel-route/pkg/travel"
	"github.com/gmarcial/travel-route/pkg/webapi/resources/travel/handlers"
	"github.com/go-chi/chi"
)

func Routing(router *chi.Mux, dependency *travel.Dependency) {
	router.Route("/travel", func(router chi.Router) {
		router.Get("/consult-better-itinerary", handlers.HandleConsultBetterItinerary(dependency.RoutesMap))
		router.Post("/register-new-route", handlers.HandleRegisterNewRoute(dependency.SourcePath, dependency.RoutesMap))
	})
}

