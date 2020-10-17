package routes_map

import (
	"errors"
	"fmt"
)

var (
	ErrItineraryUndefined = errors.New("a itinerary should be defined")
)

//Route represent a route found from search in routes map
type Route struct {
	Itinerary string
	Price     uint
}

//newRoute construct a new Route
func newRoute(itinerary string, price uint) (*Route, error) {
	if itinerary == "" {
		return nil, ErrItineraryUndefined
	}

	if price == 0 {
		return nil, ErrPriceIsZero
	}

	return &Route{
		Itinerary: itinerary,
		Price:     price,
	}, nil
}

func (route *Route) String() string {
	return fmt.Sprintf("%v > %v", route.Itinerary, route.Price)
}
