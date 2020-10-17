package register_new_route

import (
	"github.com/gmarcial/travel-route/pkg/data/source"
	routes_map "github.com/gmarcial/travel-route/pkg/travel/routes-map"
)

type RegisterNewRouteUseCase struct {
	SourcePath string
	RoutesMap *routes_map.RoutesMap
}

func (usecase *RegisterNewRouteUseCase) Handle(newRoute NewRoute) error {
	origin := newRoute.Origin
	destination := newRoute.Destination
	price := newRoute.Price

	routesMap := usecase.RoutesMap

	err := routesMap.ConnectNewRoute(origin, destination, price)
	if err != nil {
		return err
	}

	err = source.WriteNewRoute(usecase.SourcePath, origin, destination, price)
	if err != nil {
		return err
	}

	return nil
}
