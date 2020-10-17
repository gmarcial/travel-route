package routes_map

import "testing"

func TestUnitCreateAPoint(t *testing.T)  {
	//Arrange
	name := "BRC"

	//Action
	_, err := NewPoint(name)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to create a point: %v", err.Error())
	}
}

func TestUnitTryCreateAPointWithoutName(t *testing.T)  {
	//Arrange
	name := ""

	//Action
	_, err := NewPoint(name)

	//Assert
	if err == nil {
		t.Errorf("Was created a connection without name")

		return
	}

	if err != ErrNameEmptyOrNull {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryCreateAPointWithInvalidName(t *testing.T)  {
	//Arrange
	name := "BRCAAD"

	//Action
	_, err := NewPoint(name)

	//Assert
	if err == nil {
		t.Errorf("Was created a connection with invalid name")

		return
	}

	if err != ErrInvalidName {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitContainsConnectionToDestination(t *testing.T)  {
	//Arrange
	name := "BRC"
	origin, _ := NewPoint(name)

	destination, _ := NewPoint("GRU")
	_ = origin.NewConnection(destination, 1)

	//Action
	result := origin.containsConnectionTo(destination)

	//Assert
	if !result {
		t.Errorf("Was not found the existing destination")
	}
}

func TestUnitDontContainsConnectionToDestination(t *testing.T)  {
	//Arrange
	name := "BRC"
	origin, _ := NewPoint(name)
	destination, _ := NewPoint("GRU")

	//Action
	result := origin.containsConnectionTo(destination)

	//Assert
	if result {
		t.Errorf("Found the unexisting destination")
	}
}

func TestUnitCreateAConnectionWithADestination(t *testing.T)  {
	//Arrange
	origin, _ := NewPoint("ABC")
	destination, _ := NewPoint("CBA")

	//Action
	err := origin.NewConnection(destination, uint(5))

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to create a connection with a destination: %v", err.Error())
	}
}

func TestUnitTryCreateAConnectionWithTheOrigenAndDestinationEquals(t *testing.T)  {
	//Arrange
	origin, _ := NewPoint("ABC")
	destination, _ := NewPoint("ABC")

	//Action
	err := origin.NewConnection(destination, uint(5))

	//Assert
	if err == nil {
		t.Errorf("Was created a connection with the origin and destinations equals")

		return
	}

	if err != ErrOriginAndDestinationEqual {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryCreateAConnectionRepeated(t *testing.T)  {
	//Arrange
	origin, _ := NewPoint("ABC")
	destination, _ := NewPoint("CBA")

	//Action
	_ = origin.NewConnection(destination, uint(5))
	err := origin.NewConnection(destination, uint(5))

	//Assert
	if err == nil {
		t.Errorf("Was created a connection repeated")

		return
	}

	if err != ErrExistentConnection {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}
