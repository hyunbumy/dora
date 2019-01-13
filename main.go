package main

import (
	"github.com/hyunbumy/dora/pkg/genetic"
	"fmt"
)

func main() {
	fmt.Println("Test")

	size := 20
	iteration := 1000
	locations := []genetic.Location{
		genetic.Location{Name: "LAX Airport", Latitude: 33.941845, Longitude: -118.408635},
		genetic.Location{Name: "Tommy Trojan", Latitude: 34.020547, Longitude: -118.285397},
		genetic.Location{Name: "Coliseum", Latitude: 34.014156, Longitude: -118.287923},
		genetic.Location{Name: "Chinese Theatre", Latitude: 34.102021, Longitude: -118.340946},
		genetic.Location{Name: "Whiskey a Go Go", Latitude: 34.090839, Longitude: -118.385725},
		genetic.Location{Name: "Getty Center", Latitude: 34.078062, Longitude: -118.473892},
		genetic.Location{Name: "Getty Villa", Latitude: 34.045868, Longitude: -118.564850},
		genetic.Location{Name: "Disneyland", Latitude: 33.812110, Longitude: -117.918921},
		genetic.Location{Name: "The Huntington Library", Latitude: 34.129178, Longitude: -118.114556},
		genetic.Location{Name: "Rose Bowl", Latitude: 34.161373, Longitude: -118.167646},
		genetic.Location{Name: "Griffith Observatory", Latitude: 34.118509, Longitude: -118.300414},
		genetic.Location{Name: "Hollywood Sign", Latitude: 34.134124, Longitude: -118.321548},
		genetic.Location{Name: "Magic Mountain", Latitude: 34.425392, Longitude: -118.597230},
		genetic.Location{Name: "Third Street Promenade", Latitude: 34.016297, Longitude: -118.496838},
		genetic.Location{Name: "Venice Beach", Latitude: 33.985857, Longitude: -118.473167},
		genetic.Location{Name: "Catalina Island", Latitude: 33.394698, Longitude: -118.415119},
		genetic.Location{Name: "Staples Center", Latitude: 34.043097, Longitude: -118.267351},
		genetic.Location{Name: "Dodger Stadium", Latitude: 34.072744, Longitude: -118.240594},
		genetic.Location{Name: "La Brea Tar Pits", Latitude: 34.063814, Longitude: -118.355466},
		genetic.Location{Name: "Zuma Beach", Latitude: 34.015489, Longitude: -118.822160},
	}
	mutation := 0.3

	res := genetic.Run(locations, size, iteration, false, mutation)

	fmt.Println(res)
}
