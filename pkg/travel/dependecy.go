package travel

import routes_map "github.com/gmarcial/travel-route/pkg/travel/routes-map"

type Dependency struct {
	SourcePath string
	RoutesMap *routes_map.RoutesMap
}