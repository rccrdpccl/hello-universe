package main

import (
	"fmt"

	planet_types "github.com/rccrdpccl/hello-universe/planets/types"
	"github.com/rccrdpccl/hello-universe/types"
)

func main() {
	milkyWay := &types.Galaxy{}
	milkyWay.Name = "MilkyWay"

	fmt.Println(milkyWay)

	earth := planet_types.PlanetEarth()
	fmt.Println(earth)
}
