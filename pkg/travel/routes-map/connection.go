package routes_map

import "errors"

var (
	ErrNullDestination = errors.New("the destination is null")
	ErrPriceIsZero     = errors.New("the price should be more than zero")
)

//Connection represent a route that links a point of origin and other of destination, of which have a price to your
//execution.
type Connection struct {
	Destination *Point
	Price       uint
}

//newConnection constructor of connection
func newConnection(destination *Point, price uint) (*Connection, error) {
	if destination == nil {
		return nil, ErrNullDestination
	}

	if price == 0 {
		return nil, ErrPriceIsZero
	}

	return &Connection{
		Destination: destination,
		Price:       price,
	}, nil
}
