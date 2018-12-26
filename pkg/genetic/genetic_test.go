package genetic

import (
	"math/rand"
	"testing"
)

func TestInitialization(t *testing.T) {
	randGen := rand.New(rand.NewSource(0))
	size := 5
	locations := []Location{
		Location{"0", 0, 0},
		Location{"1", 1, 1},
		Location{"2", 2, 2},
		Location{"3", 3, 3},
		Location{"4", 4, 4},
		Location{"5", 5, 5},
		Location{"6", 6, 6},
		Location{"7", 7, 7},
		Location{"8", 8, 8},
		Location{"9", 9, 9},
	}

	algo := geneticAlgorithm{
		locations: locations, populationSize: size, randGen: randGen,
	}
	expected := [][]int{
		[]int{0, 3, 4, 6, 7, 8, 1, 5, 2, 9},
		[]int{0, 7, 1, 4, 8, 5, 6, 3, 2, 9},
		[]int{0, 3, 8, 4, 9, 2, 5, 6, 1, 7},
		[]int{0, 5, 7, 1, 8, 4, 9, 3, 6, 2},
		[]int{0, 8, 1, 5, 9, 4, 7, 3, 6, 2},
	}
	res := algo.initialize()

	if len(expected) != len(res) {
		t.Errorf("Length Error: expected length = %d, actual length = %d\n",
			len(expected), len(res))
	}
	for i := range res {
		if len(res[i]) != len(expected[i]) {
			t.Errorf("Route Length Error at %d\n", i)
		}
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] != expected[i][j] {
				t.Errorf("Route order Error at %dth route: %d\n", i, j)
				break
			}
		}
	}
}
