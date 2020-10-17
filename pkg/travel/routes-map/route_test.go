package routes_map

import "testing"

func TestUnitCreateARoute(t *testing.T)  {
	//Arrange
	itinerary := "GRU - BRA - FUU - JAA"
	price := uint(100)

	//Action
	_, err := newRoute(itinerary, price)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to create a route: %v", err.Error())
	}
}

func TestUnitTryCreateARouteWithAUndefinedItinerary(t *testing.T)  {
	//Arrange
	itinerary := ""
	price := uint(100)

	//Action
	_, err := newRoute(itinerary, price)

	//Assert
	if err == nil {
		t.Errorf("Was created a route with a undefined itinerary")

		return
	}

	if err != ErrItineraryUndefined {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryCreateARouteWithoutPrice(t *testing.T)  {
	//Arrange
	itinerary := "GRU - BRA - FUU - JAA"
	price := uint(0)

	//Action
	_, err := newRoute(itinerary, price)

	//Assert
	if err == nil {
		t.Errorf("Was created a route without price")

		return
	}

	if err != ErrPriceIsZero {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}
