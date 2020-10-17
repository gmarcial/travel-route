package routes_map

import "errors"

var (
	ErrNameEmptyOrNull           = errors.New("the name is empty or null")
	ErrInvalidName               = errors.New("the size of the name should be abbreviated in three characters only")
	ErrExistentConnection        = errors.New("already exist a connection to this destination")
	ErrOriginAndDestinationEqual = errors.New("the origin and destination shouldn't be the same")
)

//Point represent a point with a connections in routes map
type Point struct {
	Name        string
	Connections []*Connection
}

//NewPoint constructor of point
func NewPoint(name string) (*Point, error) {
	if name == "" {
		return nil, ErrNameEmptyOrNull
	}

	if len(name) > 3 {
		return nil, ErrInvalidName
	}

	return &Point{
		Name:        name,
		Connections: make([]*Connection, 0),
	}, nil
}

//containsConnectionTo verify is already contains a connection with this destination
func (point *Point) containsConnectionTo(destiny *Point) bool {
	for i := 0; i < len(point.Connections); i++ {
		if point.Connections[i].Destination.Name == destiny.Name {
			return true
		}
	}

	return false
}

//NewConnection create a connection with a new destination
func (point *Point) NewConnection(destination *Point, price uint) error {
	if point.Name == destination.Name {
		return ErrOriginAndDestinationEqual
	}

	contains := point.containsConnectionTo(destination)
	if contains {
		return ErrExistentConnection
	}

	connection, err := newConnection(destination, price)
	if err != nil {
		return err
	}

	point.Connections = append(point.Connections, connection)

	return nil
}
