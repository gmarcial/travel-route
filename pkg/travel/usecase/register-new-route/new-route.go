package register_new_route

type NewRoute struct {
	Origin string `json:"origin"`
	Destination string `json:"destination"`
	Price uint `json:"price"`
}
