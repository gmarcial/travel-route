package source

import (
	"errors"
	routes_map "github.com/gmarcial/travel-route/pkg/travel/routes-map"
	"os"
	"strconv"
	"strings"
	"testing"
)

func makeRoutesMapTest() map[string]*routes_map.Point {

	gru, _ := routes_map.NewPoint("GRU")
	brc, _ := routes_map.NewPoint("BRC")
	scl, _ := routes_map.NewPoint("SCL")
	cdg, _ := routes_map.NewPoint("CDG")
	orl, _ := routes_map.NewPoint("ORL")

	_ = gru.NewConnection(brc, 10)
	_ = gru.NewConnection(cdg, 75)
	_ = gru.NewConnection(scl, 20)
	_ = gru.NewConnection(orl, 56)

	_ = brc.NewConnection(scl, 5)
	_ = orl.NewConnection(cdg, 5)
	_ = scl.NewConnection(orl, 20)

	return map[string]*routes_map.Point{
		"GRU": gru,
		"BRC": brc,
		"SCL": scl,
		"CDG": cdg,
		"ORL": orl,
	}
}

func TestIntegrationMountTheRoutesMapFromSourceInCsv(t *testing.T) {
	//Arrange
	path := "../../../test/data/routes-map.csv"

	//Action
	routesMap, err := RoutesMapFromCsv(path)

	//Assert
	if err != nil {
		t.Errorf("the mounting of routes map from source csv failed: %v", err.Error())
	}

	routesMapTest := makeRoutesMapTest()

	if len(routesMap) != len(routesMapTest) {
		t.Error("the quantity of routes is different from expected")
	}

	pointsNames := []string{"GRU", "BRC", "SCL", "CDG", "ORL"}

	for _, name := range pointsNames {
		point, exist := routesMap[name]
		if !exist {
			t.Errorf("don't was found the point expecte: %v", name)
		}

		pointTest, _ := routesMapTest[name]

		if point.Name != pointTest.Name {
			t.Error("a point with the different name of expected")
		}

		for i := 0; i < len(point.Connections); i++ {
			if point.Connections[i].Price != pointTest.Connections[i].Price {
				t.Error("a point with the different price of expected")
			}

			if point.Connections[i].Destination.Name != pointTest.Connections[i].Destination.Name {
				t.Error("a point with the different destination of expected")
			}
		}
	}
}

func TestIntegrationTryMountTheRoutesMapFromSourceInCsvWithEmptyPath(t *testing.T) {
	//Arrange
	path := ""

	//Action
	_, err := RoutesMapFromCsv(path)

	//Assert
	if err == nil {
		t.Error("the routes map was created the same with an empty path")

		return
	}

	if err != ErrEmptyPath {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestIntegrationTryMountTheRoutesMapFromSourceInCsvNotExisting(t *testing.T) {
	//Arrange
	path := "../../../test/data/routes-map-sos.csv"

	//Action
	_, err := RoutesMapFromCsv(path)

	//Assert
	if err == nil {
		t.Error("the routes map was created the same with a source not existing")

		return
	}

	if errors.Is(err, os.PathError{}.Err) {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryMountTheRoutesMapFromSourceInCsvWithEmptyName(t *testing.T) {
	//Arrange
	source := `
GRU,BRC,10
,SCL,5
GRU,CDG,75`

	//Action
	_, err := pointsFromCsv(strings.NewReader(source))

	//Assert
	if err == nil {
		t.Error("the routes map was created the same with a empty name")

		return
	}

	if err != routes_map.ErrNameEmptyOrNull {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryMountTheRoutesMapFromSourceInCsvWithInvalidName(t *testing.T) {
	//Arrange
	source := `
GRU,BRC,10
BRC,SCLLLL,5
GRU,CDG,75`

	//Action
	_, err := pointsFromCsv(strings.NewReader(source))

	//Assert
	if err == nil {
		t.Error("the routes map was created the same with a invalid name")

		return
	}

	if err != routes_map.ErrInvalidName {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryMountTheRoutesMapFromSourceInCsvWithPriceNegative(t *testing.T) {
	//Arrange
	source := `
GRU,BRC,10
BRC,SCL,5
GRU,CDG,-75`

	//Action
	_, err := pointsFromCsv(strings.NewReader(source))

	//Assert
	if err == nil {
		t.Error("the routes map was created the same with a price negative")

		return
	}

	if errors.Is(err, strconv.NumError{}.Err) {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryMountTheRoutesMapFromSourceInCsvWithAOrigenEqualTheADestination(t *testing.T) {
	//Arrange
	source := `
GRU,BRC,10
BRC,BRC,5
GRU,CDG,75`

	//Action
	_, err := pointsFromCsv(strings.NewReader(source))

	//Assert
	if err == nil {
		t.Error("the routes map was created the same with a origen equal the a destination")

		return
	}

	if err != routes_map.ErrOriginAndDestinationEqual {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryMountTheRoutesMapFromSourceInCsvWithRepeatedConnection(t *testing.T) {
	//Arrange
	source := `
GRU,BRC,10
BRC,SCL,5
BRC,SCL,5
GRU,CDG,75`

	//Action
	_, err := pointsFromCsv(strings.NewReader(source))

	//Assert
	if err == nil {
		t.Error("the routes map was created the same with a repeated connection")

		return
	}

	if err != routes_map.ErrExistentConnection {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryMountTheRoutesMapFromSourceInCsvWithANullDestination(t *testing.T) {
	//Arrange
	source := `
GRU,BRC,10
BRC,,5
GRU,CDG,75`

	//Action
	_, err := pointsFromCsv(strings.NewReader(source))

	//Assert
	if err == nil {
		t.Error("the routes map was created the same with a null destination")

		return
	}

	if err != routes_map.ErrNameEmptyOrNull {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestUnitTryMountTheRoutesMapFromSourceInCsvWithAPriceZero(t *testing.T) {
	//Arrange
	source := `
GRU,BRC,10
BRC,SCL,5
GRU,CDG,0`

	//Action
	_, err := pointsFromCsv(strings.NewReader(source))

	//Assert
	if err == nil {
		t.Error("the routes map was created the same with a price zero")

		return
	}

	if err != routes_map.ErrPriceIsZero {
		t.Errorf("the error returned is different from expected: %v", err.Error())
	}
}

func TestIntegrationWriteNewRoute(t *testing.T) {
	//Arrange
	path := "../../../test/data/routes-map.csv"
	origin := "ABC"
	destination := "CBA"

	//Action
	err := WriteNewRoute(path, origin, destination, 1000)

	//Assert
	if err != nil {
		t.Errorf("the write of a new route fail: %v", err.Error())
	}

	routesMap, _ := RoutesMapFromCsv(path)
	originPoint, exist := routesMap[origin]
	if !exist {
		t.Error("don't found in routes map the origin of new route")
	}

	_, exist = routesMap[destination]
	if !exist {
		t.Error("don't found in routes map the destination of new route")
	}

	result := false
	for _, connection := range originPoint.Connections {
		if connection.Destination.Name == destination {
			result = true
		}
	}
	if !result {
		t.Error("don't was connected the origin and destination in new route")
	}
}
