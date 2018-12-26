package genetic

import (
	"math/rand"
)

type geneticAlgorithm struct {
	locations      []Location
	randGen        *rand.Rand
	populationSize int
}

// Run the genetic algorithm given the locations and the population size
func Run(locations []Location, size int, randGen *rand.Rand) []Location {
	// randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	algo := geneticAlgorithm{locations, randGen, size}
	algo.initialize()

	return []Location{}
}

func (algo *geneticAlgorithm) initialize() [][]int {
	size := algo.populationSize
	locSize := len(algo.locations)
	population := make([][]int, size) // Make a list of possible routes
	for i := 0; i < size; i++ {
		population[i] = make([]int, locSize)
		for j := 0; j < locSize; j++ {
			population[i][j] = j
		}

		// Populate with random routes
		tempSlice := population[i][1:]
		algo.randGen.Shuffle(len(tempSlice), func(i, j int) {
			tempSlice[i], tempSlice[j] = tempSlice[j], tempSlice[i]
		})
	}
	return population
}
