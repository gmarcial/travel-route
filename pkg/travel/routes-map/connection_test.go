package routes_map

import "testing"

func TestUnitCreateAConnection(t *testing.T) {
	//Arrange
	destination, _ := NewPoint("DTN")
	price := uint(5)

	//Action
	_, err := newConnection(destination, price)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to create a connection: %v", err.Error())
	}
}

func TestUnitTryCreateAConnectionWithoutDestination(t *testing.T) {
	//Arrange
	price := uint(5)

	//Action
	_, err := newConnection(nil, price)

	//Assert
	if err == nil {
		t.Errorf("Was created a connection without destination")

		return
	}

	if err != ErrNullDestination {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryCreateAConnectionWithPriceZero(t *testing.T) {
	//Arrange
	destination, _ := NewPoint("DTN")
	price := uint(0)

	//Action
	_, err := newConnection(destination, price)

	//Assert
	if err == nil {
		t.Errorf("Was created a connection without destination")

		return
	}

	if err != ErrPriceIsZero {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}