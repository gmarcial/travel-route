package consult_better_itinerary

import routes_map "github.com/gmarcial/travel-route/pkg/travel/routes-map"

type ConsultBetterItineraryUseCase struct {
	RoutesMap *routes_map.RoutesMap
}

func (usecase *ConsultBetterItineraryUseCase) Handle(route *QueryRoute) (*routes_map.Route, error) {
	origin := route.Origin
	destination := route.Destination

	return usecase.RoutesMap.ComputeBetterRoute(origin, destination)
}
