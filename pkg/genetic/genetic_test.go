package genetic

import (
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

	algo := geneticAlgorithm{
		locations: locations, populationSize: size, randGen: randGen,
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
	algo.initialize()
	res := algo.routes

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
