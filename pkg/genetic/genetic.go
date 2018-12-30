package genetic

import (
	"math"
	"math/rand"
)

type parents struct {
	parent1 int
	parent2 int
}

// Run the genetic algorithm given the locations and the population size
func Run(locations []Location, popSize int, isTransit bool) []Location {
	// randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	return []Location{}
}

func initialize(locations []Location, popSize int, randGen *rand.Rand) []Route {
	locSize := len(locations)
	population := make([]Route, popSize) // Make a list of possible routes
	for i := 0; i < popSize; i++ {
		population[i] = Route{make([]int, locSize), -1}
		for j := 0; j < locSize; j++ {
			population[i].order[j] = j
		}

		// Populate with random routes
		tempSlice := population[i].order[1:]
		randGen.Shuffle(len(tempSlice), func(i, j int) {
			tempSlice[i], tempSlice[j] = tempSlice[j], tempSlice[i]
		})
	}

	return population
}

func calcFitness(locations []Location, order []int, isTransit bool) float64 {
	fitness := float64(0)
	if isTransit {
		// Use Google Maps API to calculate the total travel time
	} else {
		// Use Haversine
		for i := 0; i < len(order)-1; i++ {
			from := locations[order[i]]
			to := locations[order[i+1]]
			fitness += calcHaversine(
				from.latitude, from.longitude, to.latitude, to.longitude,
			)
		}
	}
	return fitness
}

func calcHaversine(latFrom, lonFrom, latTo, lonTo float64) float64 {
	const earthRadius = float64(6371)
	var distance float64

	var deltaLat = (latTo - latFrom) * (math.Pi / 180)
	var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)

	var a = math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
			math.Cos(latFrom*(math.Pi/180))*math.Cos(latTo*(math.Pi/180))*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance = earthRadius * c

	return distance
}
