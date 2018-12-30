package genetic

import (
	"math"
	"math/rand"
	"sort"
	"time"
	
)

type parentsStruct struct {
	parent1 int
	parent2 int
}

// Run the genetic algorithm given the locations and the population size
func Run(locations []Location, popSize, iteration int, isTransit bool) []Location {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	population := initialize(locations, popSize, randGen)

	for i := 0; i < iteration; i++ {
		// Fitness calculation
		calcFitness(locations, population, isTransit)

		// Selection
		selectedParents := rouletteWheelSelection(population, randGen)

		// Crossover
		population = crossover(population, selectedParents, randGen)
	}

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

func calcFitness(locations []Location, routes []Route, isTransit bool) {
	fitnessFunc := getDistFitness
	if isTransit {
		fitnessFunc = getTransitFitness
	}

	for i := 0; i < len(routes); i++ {
		routes[i].fitness = fitnessFunc(locations, routes[i].order)
	}
}

func getDistFitness(locations []Location, order []int) float64 {
	fitness := float64(0)
	// Use Haversine
	for i := 0; i < len(order)-1; i++ {
		from := locations[order[i]]
		to := locations[order[i+1]]
		fitness += calcHaversine(
			from.latitude, from.longitude, to.latitude, to.longitude,
		)
	}
	return 1 / fitness // Inverse of the cost
}

func getTransitFitness(locations []Location, order []int) float64 {
	// Use Google Maps API to calculate the total travel time
	return -1
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

	return distance // Distance in km
}

func rouletteWheelSelection(routes []Route, randGen *rand.Rand) []parentsStruct {
	// Sort in a descending order
	sort.SliceStable(routes, func(i, j int) bool {
		return routes[i].fitness > routes[j].fitness
	})

	var totalFitness float64
	for _, v := range routes {
		totalFitness += v.fitness
	}

	selectedParents := make([]parentsStruct, len(routes))
	for i := 0; i < len(selectedParents); i++ {
		selectedParents[i] = parentsStruct{
			getParent(randGen, routes, totalFitness),
			getParent(randGen, routes, totalFitness),
		}
	}

	return selectedParents
}

func getParent(randGen *rand.Rand, routes []Route, totFit float64) int {
	randProb := randGen.Float64()
	runningTotal := float64(0)
	for i, v := range routes {
		runningTotal += v.fitness / totFit
		if randProb < runningTotal {
			return i
		}
	}
	return len(routes) - 1 // Return the last element
}

func crossover(routes []Route, selectedParents []parentsStruct, randGen *rand.Rand) []Route {
	newRoutes := make([]Route, len(routes))
	for i := 0; i < len(selectedParents); i++ {
		newRoutes[i] = getChild(routes, selectedParents[i], randGen)
	}

	return newRoutes
}

func getChild(routes []Route, parents parentsStruct, randGen *rand.Rand) Route {
	coin := randGen.Intn(2)	// Determine which parent to go first
	firstParent, secondParent := parents.parent1, parents.parent2
	if coin == 1 {
		firstParent, secondParent = secondParent, firstParent
	}

	newRoute := Route{make([]int, len(routes[0].order)), -1}
	crossInd := randGen.Intn(len(newRoute.order)-2)
	crossInd++
	used := make(map[int]bool)

	for i := 1; i <= crossInd; i++ {
		newRoute.order[i] = routes[firstParent].order[i]
		used[newRoute.order[i]] = true
	}
	for i := 1; i < len(newRoute.order); i++ {
		if !used[routes[secondParent].order[i]] {
			crossInd++
			newRoute.order[crossInd] = routes[secondParent].order[i]
		}
	}

	return newRoute
}
