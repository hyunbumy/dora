package genetic

import (
	"math/rand"
)

// Run the genetic algorithm given the locations and the population size
func Run(locations []Location, size int, randGen *rand.Rand) [][]int {
	// randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	return initialize(len(locations), size, randGen)
}

func initialize(locSize, size int, randGen *rand.Rand) [][]int {
	population := make([][]int, size) // Make a list of possible routes
	for i := 0; i < size; i++ {
		population[i] = make([]int, locSize)
		for j := 0; j < locSize; j++ {
			population[i][j] = j
		}

		// Populate with random routes
		tempSlice := population[i][1:]
		randGen.Shuffle(len(tempSlice), func(i, j int) {
			tempSlice[i], tempSlice[j] = tempSlice[j], tempSlice[i]
		})
	}
	return population
}
