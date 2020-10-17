package console

import (
	"bufio"
	"fmt"
	routes_map "github.com/gmarcial/travel-route/pkg/travel/routes-map"
	consult_better_itinerary "github.com/gmarcial/travel-route/pkg/travel/usecase/consult-better-itinerary"
	"os"
	"strings"
)

func Start(routesMap *routes_map.RoutesMap) {
	usecase := consult_better_itinerary.ConsultBetterItineraryUseCase{RoutesMap: routesMap}

	for {
		fmt.Println()
		fmt.Println("Please enter the route, example: GRU-CDG")
		command, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		args := strings.Split(string(command), "-")
		if len(args) > 2 {
			fmt.Println("Only two arguments are accepted, example: GRU-CDG")
			continue
		}

		if len(args[0]) > 3{
			fmt.Println("Invalid argument, only three characters example: GRU-CGD")
			continue
		}

		if len(args[1]) > 3{
			fmt.Println("Invalid argument, only three characters example: GRU-CGD")
			continue
		}

		queryRoute := &consult_better_itinerary.QueryRoute{
			Origin:      args[0],
			Destination: args[1],
		}
		route, err := usecase.Handle(queryRoute)
		if err != nil {
			fmt.Printf("Message: %v \n", err.Error())
		} else {
			fmt.Printf("Best route: %v \n", route)
		}
	}

}
