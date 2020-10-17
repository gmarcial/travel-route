package main

import (
	"github.com/gmarcial/travel-route/pkg/console"
	"github.com/gmarcial/travel-route/pkg/data/source"
	"github.com/gmarcial/travel-route/pkg/travel"
	"github.com/gmarcial/travel-route/pkg/webapi"
	"os"
)

func main() {
	sourcePath := os.Args[1]
	routesMap, err := source.RoutesMapFromCsv(sourcePath)
	if err != nil {
		panic(err.Error())
	}

	dependency := &travel.Dependency{
		SourcePath: sourcePath,
		RoutesMap: &routesMap,
	}

	go webapi.ListenAndServe(dependency)
	console.Start(&routesMap)
}

