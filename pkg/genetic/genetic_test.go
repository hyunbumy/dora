package genetic

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInitialization(t *testing.T) {
	randGen := rand.New(rand.NewSource(0))
	size := 8
	locations := []Location{
		Location{"JFK Airport", 40.642523, -73.778216},
		Location{"Empire State Building", 40.748522, -73.985505},
		Location{"Statue of Liberty", 40.689282, -74.044533},
		Location{"Times Square", 40.758871, -73.985056},
		Location{"9/11 Memorial", 40.711573, -74.013283},
		Location{"Yankee Stadium", 40.829610, -73.926164},
		Location{"Long Beach Boardwalk", 40.583162, -73.660001},
		Location{"Madison Square Garden", 40.750610, -73.993492},
		Location{"Central Park", 40.771222, -73.973430},
		Location{"Columbia University", 40.806814, -73.962496},
	}

	expected := []Route{
		Route{[]int{0, 3, 4, 6, 7, 8, 1, 5, 2, 9}, -1},
		Route{[]int{0, 7, 1, 4, 8, 5, 6, 3, 2, 9}, -1},
		Route{[]int{0, 3, 8, 4, 9, 2, 5, 6, 1, 7}, -1},
		Route{[]int{0, 5, 7, 1, 8, 4, 9, 3, 6, 2}, -1},
		Route{[]int{0, 8, 1, 5, 9, 4, 7, 3, 6, 2}, -1},
		Route{[]int{0, 7, 2, 3, 8, 9, 4, 6, 1, 5}, -1},
		Route{[]int{0, 6, 2, 8, 4, 3, 9, 1, 7, 5}, -1},
		Route{[]int{0, 2, 5, 8, 3, 9, 4, 1, 7, 6}, -1},
	}
	res := initialize(locations, size, randGen)

	if len(expected) != len(res) {
		t.Errorf("Length Error: expected length = %d, actual length = %d\n",
			len(expected), len(res))
	}
	for i := range res {
		if len(res[i].order) != len(expected[i].order) {
			t.Errorf("Route Length Error at %d\n", i)
		}
		for j := 0; j < len(res[i].order); j++ {
			if res[i].order[j] != expected[i].order[j] {
				t.Errorf("Route order Error at %dth route: %d\n", i, j)
				break
			}
		}
	}
}

func TestHaversineFitness(t *testing.T) {
	size := 8
	locations := []Location{
		Location{"LAX Airport", 33.941845, -118.408635},
		Location{"Tommy Trojan", 34.020547, -118.285397},
		Location{"Coliseum", 34.014156, -118.287923},
		Location{"Chinese Theatre", 34.102021, -118.340946},
		Location{"Whiskey a Go Go", 34.090839, -118.385725},
		Location{"Getty Center", 34.078062, -118.473892},
		Location{"Getty Villa", 34.045868, -118.564850},
		Location{"Disneyland", 33.812110, -117.918921},
		Location{"The Huntington Library", 34.129178, -118.114556},
		Location{"Rose Bowl", 34.161373, -118.167646},
		Location{"Griffith Observatory", 34.118509, -118.300414},
		Location{"Hollywood Sign", 34.134124, -118.321548},
		Location{"Magic Mountain", 34.425392, -118.597230},
		Location{"Third Street Promenade", 34.016297, -118.496838},
		Location{"Venice Beach", 33.985857, -118.473167},
		Location{"Catalina Island", 33.394698, -118.415119},
		Location{"Staples Center", 34.043097, -118.267351},
		Location{"Dodger Stadium", 34.072744, -118.240594},
		Location{"La Brea Tar Pits", 34.063814, -118.355466},
		Location{"Zuma Beach", 34.015489, -118.822160},
	}

	routes := [][]int{
		{0, 9, 6, 11, 14, 17, 19, 12, 13, 4, 8, 1, 2, 7, 18, 5, 16, 3, 10, 15},
		{0, 17, 11, 7, 6, 2, 12, 18, 4, 19, 10, 8, 15, 16, 13, 14, 1, 5, 9, 3},
		{0, 2, 8, 17, 1, 9, 13, 15, 18, 12, 11, 7, 6, 4, 5, 14, 19, 16, 3, 10},
		{0, 18, 8, 4, 3, 7, 12, 6, 14, 16, 11, 15, 17, 10, 1, 13, 19, 5, 2, 9},
		{0, 14, 18, 2, 10, 4, 16, 11, 3, 12, 9, 15, 7, 5, 6, 19, 1, 13, 8, 17},
		{0, 5, 4, 3, 15, 6, 1, 9, 13, 14, 16, 12, 2, 7, 19, 17, 10, 8, 18, 11},
		{0, 18, 5, 17, 13, 14, 12, 1, 15, 3, 11, 10, 16, 2, 19, 8, 4, 9, 6, 7},
		{0, 1, 12, 5, 3, 9, 16, 2, 13, 4, 10, 7, 15, 18, 11, 17, 8, 19, 14, 6},
	}

	fitness := make([]float64, size)
	for i := 0; i < size; i++ {
		fitness[i] = calcFitness(locations, routes[i], false)
	}

	fmt.Println(fitness)
}
