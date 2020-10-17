package source

import (
	"encoding/csv"
	"errors"
	"github.com/gmarcial/travel-route/pkg/travel/routes-map"
	"io"
	"os"
	"strconv"
)

var (
	ErrEmptyPath = errors.New("the path of file is empty")
)

//retrieveOrCreatePoint retrieve a point in auxiliary map or case not exist, create, and store.
func retrieveOrCreatePoint(pointName string, points map[string]*routes_map.Point) (*routes_map.Point, error) {
	point, exist := points[pointName]
	if !exist {
		var err error
		point, err = routes_map.NewPoint(pointName)
		if err != nil {
			return nil, err
		}

		points[pointName] = point
	}

	return point, nil
}

//pointsFromCsv mount and construct the points from a definition in csv
func pointsFromCsv(reader io.Reader) (map[string]*routes_map.Point, error) {
	csvReader := csv.NewReader(reader)
	points := make(map[string]*routes_map.Point)

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		origin, err := retrieveOrCreatePoint(line[0], points)
		if err != nil {
			return nil, err
		}

		destination, err := retrieveOrCreatePoint(line[1], points)
		if err != nil {
			return nil, err
		}

		price, err := strconv.ParseUint(line[2], 10, 64)
		if err != nil {
			return nil, err
		}

		err = origin.NewConnection(destination, uint(price))
		if err != nil {
			return nil, err
		}
	}

	return points, nil
}

//RoutesMapFromCsv mount and construct the route map of travels from a definition in csv, deserializing in routes map
//struct that is returned.
func RoutesMapFromCsv(path string) (routes_map.RoutesMap, error) {
	if path == "" {
		return nil, ErrEmptyPath
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	points, err := pointsFromCsv(file)
	if err != nil {
		return nil, err
	}

	routes, err := routes_map.New(points)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

//WriteNewRoute write a new route between a origin and destination
func WriteNewRoute(path, origin, destination string, price uint) error {
	if path == "" {
		return ErrEmptyPath
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	priceInString := strconv.Itoa(int(price))
	newRecord := []string{origin, destination, priceInString}

	writer := csv.NewWriter(file)
	err = writer.Write(newRecord)
	if err != nil {
		return err
	}

	writer.Flush()
	if writer.Error() != nil {
		return err
	}

	return nil
}
