package types

func PlanetEarth() *Planet {
	return &Planet{
		Name: "PlanetEarth",
	}
}

type Planet struct {
	Name string
}
