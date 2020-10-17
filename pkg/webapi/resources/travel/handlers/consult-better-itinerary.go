package handlers

import (
	"encoding/json"
	routes_map "github.com/gmarcial/travel-route/pkg/travel/routes-map"
	consult_better_itinerary "github.com/gmarcial/travel-route/pkg/travel/usecase/consult-better-itinerary"
	"github.com/gmarcial/travel-route/pkg/webapi/response"
	"net/http"
)

func HandleConsultBetterItinerary(routesMap *routes_map.RoutesMap) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		usecase := consult_better_itinerary.ConsultBetterItineraryUseCase{RoutesMap: routesMap}

		origin := request.URL.Query().Get("origin")
		destination := request.URL.Query().Get("destination")

		queryRoute := &consult_better_itinerary.QueryRoute{
			Origin:      origin,
			Destination: destination,
		}

		result, err := usecase.Handle(queryRoute)
		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			err := encoder.Encode(response.Error{Value: err.Error()})
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		err = encoder.Encode(result)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
}
