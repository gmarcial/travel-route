package handlers

import (
	"encoding/json"
	routes_map "github.com/gmarcial/travel-route/pkg/travel/routes-map"
	register_new_route "github.com/gmarcial/travel-route/pkg/travel/usecase/register-new-route"
	"github.com/gmarcial/travel-route/pkg/webapi/response"
	"net/http"
)

func HandleRegisterNewRoute(sourcePath string, routesMap *routes_map.RoutesMap) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		usecase := register_new_route.RegisterNewRouteUseCase{
			SourcePath: sourcePath,
			RoutesMap:  routesMap,
		}

		decoder := json.NewDecoder(request.Body)
		var newRoute register_new_route.NewRoute
		err := decoder.Decode(&newRoute)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}

		err = usecase.Handle(newRoute)
		encoder := json.NewEncoder(writer)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			err := encoder.Encode(response.Error{Value: err.Error()})
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
			}
		}

		writer.WriteHeader(http.StatusCreated)
	}
}
