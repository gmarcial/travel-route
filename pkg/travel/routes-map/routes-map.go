package routes_map

import (
	"errors"
	"fmt"
	"strings"
)

var ErrPointsEmptyOrNull = errors.New("the points are necessary to compose the routes map")

//RoutesMap represent the available routes to travel, where is formed per points linked per connections.
type RoutesMap map[string]*Point

//New construct RoutesMap
func New(points map[string]*Point) (RoutesMap, error) {
	if points == nil || len(points) == 0 {
		return nil, ErrPointsEmptyOrNull
	}

	return points, nil
}

//state representing the actual state of search
type state struct {
	actualPath []string
	totalPrice uint
}

//storeState store the state how a valid route, that represent the itinerary
//between the origin and destination
func storeState(foundRoutes *[]*Route, target string, state state) error {
	state.actualPath = append(state.actualPath, target)
	itinerary := strings.Join(state.actualPath, " - ")
	route, err := newRoute(itinerary, state.totalPrice)
	if err != nil {
		return err
	}

	*foundRoutes = append(*foundRoutes, route)
	return nil
}

//search finds recursively all routes that correspond the origin to target
func search(origin *Point, target string, state state) ([]*Route, error) {
	foundRoutes := make([]*Route, 0)
	state.actualPath = append(state.actualPath, origin.Name)

	if origin.Name == target {
		itinerary := strings.Join(state.actualPath, " - ")
		route, err := newRoute(itinerary, state.totalPrice)
		if err != nil {
			return nil, err
		}
		foundRoutes = append(foundRoutes, route)
	}

	for i := 0; i < len(origin.Connections); i++ {
		connection := origin.Connections[i]
		state.totalPrice += connection.Price

		if len(connection.Destination.Connections) > 0 {
			newFoundRoutes, err := search(connection.Destination, target, state)
			if err != nil {
				return nil, err
			}

			foundRoutes = append(foundRoutes, newFoundRoutes...)
		} else if connection.Destination.Name == target {
			err := storeState(&foundRoutes, target, state)
			if err != nil {
				return nil, err
			}
		}

		state.totalPrice -= connection.Price
	}

	return foundRoutes, nil
}

//ComputeBetterRoute compute and return the cheapest route of a origin(point A) to a destination(point B)
func (routesMap RoutesMap) ComputeBetterRoute(origin, destination string) (*Route, error) {
	if _, exist := routesMap[origin]; !exist {
		return nil, fmt.Errorf("not found the informed %v %v in routes map", "origin", origin)
	}

	if _, exist := routesMap[destination]; !exist {
		return nil, fmt.Errorf("not found the informed %v %v in routes map", "destination", destination)
	}

	originPoint := routesMap[origin]
	state := state{make([]string, 0), uint(0)}

	routes, err := search(originPoint, destination, state)
	if err != nil {
		return nil, err
	}

	bestRoute := new(Route)
	for i := 0; i < len(routes); i++ {
		if routes[i].Price < bestRoute.Price || bestRoute.Price == 0 {
			bestRoute = routes[i]
		}
	}

	return bestRoute, nil
}

//ConnectNewRoute connect two points to create a new route in map
func (routesMap RoutesMap) ConnectNewRoute(origin, destination string, price uint) error {
	var originPoint *Point
	originPoint, exist := routesMap[origin]
	if exist {
		for _, connection := range originPoint.Connections{
			if connection.Destination.Name == destination  {
				return ErrExistentConnection
			}
		}
	} else {
		op, err := NewPoint(origin)
		if err != nil {
			return err
		}

		originPoint = op
		routesMap[origin] = originPoint
	}

	var destinationPoint *Point
	destinationPoint, exist = routesMap[destination]
	if !exist {
		dp, err := NewPoint(destination)
		if err != nil {
			return err
		}

		destinationPoint = dp
		routesMap[destination] = destinationPoint
	}

	err := originPoint.NewConnection(destinationPoint, price)
	if err != nil {
		return err
	}

	return nil
}
