package routes_map

import (
	"fmt"
	"testing"
)

func makeRoutesMapTest() RoutesMap {

	gru, _ := NewPoint("GRU")
	brc, _ := NewPoint("BRC")
	scl, _ := NewPoint("SCL")
	cdg, _ := NewPoint("CDG")
	orl, _ := NewPoint("ORL")

	_ = gru.NewConnection(brc, 10)
	_ = gru.NewConnection(cdg, 75)
	_ = gru.NewConnection(scl, 20)
	_ = gru.NewConnection(orl, 56)

	_ = brc.NewConnection(scl, 5)
	_ = orl.NewConnection(cdg, 5)
	_ = scl.NewConnection(orl, 20)


	return map[string]*Point{
		"GRU": gru,
		"BRC": brc,
		"SCL": scl,
		"CDG": cdg,
		"ORL": orl,
	}
}

func TestUnitCreateARoutesMap(t *testing.T)  {
	//Arrange
	points := map[string]*Point{
		"GRU": {
			Name: "GRU",
			Connections: []*Connection{
				{
					Destination: &Point{
						Name:        "ABC",
						Connections: nil,
					},
					Price: 10,
				},
			},
		},
	}

	//Action
	_, err := New(points)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to create the routes map: %v", err.Error())
	}
}

func TestUnitTryCreateARoutesMapWithoutPoints(t *testing.T)  {
	//Arrange
	//Action
	_, err := New(nil)

	//Assert
	if err == nil {
		t.Errorf("Was created a connection without points")

		return
	}

	if err != ErrPointsEmptyOrNull {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryCreateARoutesMapEmptyPoints(t *testing.T)  {
	//Arrange
	points := make(map[string]*Point, 0)

	//Action
	_, err := New(points)

	//Assert
	if err == nil {
		t.Errorf("Was created a connection without points")

		return
	}

	if err != ErrPointsEmptyOrNull {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryComputeBetterRouteWithAOriginNonexistent(t *testing.T){
	//Arrange
	routesMap := makeRoutesMapTest()
	nonexistentOrigin := "ABC"

	//Action
	_, err := routesMap.ComputeBetterRoute(nonexistentOrigin, "CDG")

	//Assert
	if err == nil {
		t.Errorf("Was computed the better route with a origin nonexistent")

		return
	}

	expectedError := fmt.Errorf("not found the informed %v %v in routes map", "origin", nonexistentOrigin)
	if err.Error() != expectedError.Error(){
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryComputeBetterRouteWithADestinationNonexistent(t *testing.T){
	//Arrange
	routesMap := makeRoutesMapTest()
	nonexistentDestination := "ABC"

	//Action
	_, err := routesMap.ComputeBetterRoute("GRU", nonexistentDestination)

	//Assert
	if err == nil {
		t.Errorf("Was computed the better route with a origin nonexistent")

		return
	}

	expectedError := fmt.Errorf("not found the informed %v %v in routes map", "destination", nonexistentDestination)
	if err.Error() != expectedError.Error(){
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitSearchAllRoutesThatCorrespondTheOriginToTarget(t *testing.T){
	//Arrange
	routesMap := makeRoutesMapTest()
	state := state{
		actualPath: make([]string, 0),
		totalPrice: uint(0),
	}

	point := routesMap["GRU"]

	//Action
	routes, err := search(point, "CDG", state)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to search the routes: %v", err.Error())
	}

	if len(routes) != 4 {
		t.Error("The quantity of routes found is different of expected")
	}

	expectedItineraries := []string{
		"GRU - BRC - SCL - ORL - CDG > 40",
		"GRU - CDG > 75",
		"GRU - SCL - ORL - CDG > 45",
		"GRU - ORL - CDG > 61",
	}

	for i := 0; i < len(routes); i++ {
		if routes[i].String() != expectedItineraries[i] {
			t.Errorf("A itinerary different of expected: Actual %v, Expected %v", routes[i].String(),  expectedItineraries[i])
		}
	}
}

func TestUnitStoreAStateHowAValidRoute(t *testing.T) {
	//Arrange
	foundedRoutes := make([]*Route,0)
	target := "CDG"
	state := state{
		[]string{"ABC", "DEF"},
		uint(100),
	}

	//Action
	err := storeState(&foundedRoutes, target, state)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to store the state: %v", err.Error())
	}

	if len(foundedRoutes) != 1 {
		t.Error("founded routes size not expected")
	}

	if foundedRoutes[0].String() != "ABC - DEF - CDG > 100" {
		t.Error("not expected route")
	}
}

func TestUnitConnectANewRouteWithoutNewPoints(t *testing.T) {
	//Arrange
	routesMap := makeRoutesMapTest()
	origin := "SLC"
	destination := "CDG"
	price := uint(100)

	//Action
	err := routesMap.ConnectNewRoute(origin, destination, price)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to connect a new route: %v", err.Error())
	}

	point := routesMap[origin]
	result := false
	for _, connection := range point.Connections {
		if connection.Destination.Name == destination {
			result = true
		}
	}

	if !result {
		t.Error("Don't  connected the new route")
	}
}

func TestUnitConnectANewRouteWithNewPointHowOrigin(t *testing.T) {
	//Arrange
	routesMap := makeRoutesMapTest()
	origin := "BRA"
	destination := "CDG"
	price := uint(100)

	//Action
	err := routesMap.ConnectNewRoute(origin, destination, price)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to connect a new route: %v", err.Error())
	}

	point, exist := routesMap[origin]
	if !exist {
		t.Error("The new point don't was added")
	}

	result := false
	for _, connection := range point.Connections {
		if connection.Destination.Name == destination {
			result = true
		}
	}

	if !result {
		t.Error("Don't  connected the new route")
	}
}

func TestUnitConnectANewRouteWithNewPointHowDestination(t *testing.T) {
	//Arrange
	routesMap := makeRoutesMapTest()
	origin := "GRU"
	destination := "BRA"
	price := uint(100)

	//Action
	err := routesMap.ConnectNewRoute(origin, destination, price)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to connect a new route: %v", err.Error())
	}

	_, exist := routesMap[destination]
	if !exist {
		t.Error("The new point don't was added")
	}

	point := routesMap[origin]
	result := false
	for _, connection := range point.Connections {
		if connection.Destination.Name == destination {
			result = true
		}
	}

	if !result {
		t.Error("Don't  connected the new route")
	}
}

func TestUnitConnectANewRouteWithNewPointsHowOriginAndDestination(t *testing.T) {
	//Arrange
	routesMap := makeRoutesMapTest()
	origin := "BRA"
	destination := "ARB"
	price := uint(100)

	//Action
	err := routesMap.ConnectNewRoute(origin, destination, price)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to connect a new route: %v", err.Error())
	}

	_, exist := routesMap[destination]
	if !exist {
		t.Error("The new origin point don't was added")
	}

	point, exist := routesMap[origin]
	if !exist {
		t.Error("The new destination point don't was added")
	}

	result := false
	for _, connection := range point.Connections {
		if connection.Destination.Name == destination {
			result = true
		}
	}

	if !result {
		t.Error("Don't  connected the new route")
	}
}

func TestUnitTryConnectANewRouteThatAlreadyExists(t *testing.T) {
	//Arrange
	routesMap := makeRoutesMapTest()
	origin := "GRU"
	destination := "CDG"
	price := uint(100)

	//Action
	err := routesMap.ConnectNewRoute(origin, destination, price)

	//Assert
	if err == nil {
		t.Errorf("Was connect a new route that already exists")

		return
	}

	if err != ErrExistentConnection {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}