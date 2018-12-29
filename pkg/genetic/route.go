package genetic

// Route represents the order of the locations to visit and its fitness (total time)
type Route struct {
	order   []int
	fitness float64
}
