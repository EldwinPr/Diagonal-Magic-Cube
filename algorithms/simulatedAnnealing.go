package algorithms

import (
	"DiagonalMagicCube/cubeFuncs"
	"DiagonalMagicCube/objectiveFunction"
	"math"
	"math/rand"
)

func SimulatedAnnealing(cube [5][5][5]int) [5][5][5]int {
	T := 3000000000.0 // Initial temperature
	iteration := 1    // Iteration number

	// helper function to update temperature
	Schedule := func(T float64, iteration int) float64 {
		// Update temperature based on iteration number
		if iteration < 1000 {
			return T * 0.99999999
		} else if iteration < 10000 {
			return T * 0.999999
		} else if iteration < 100000 {
			return T * 0.9999
		} else if iteration < 1000000 {
			return T * 0.99999
		} else {
			return T * 0.999999
		}
	}

	exit := false
	for !exit {
		T = Schedule(T, iteration) // Update temperature
		iteration++
		if T < 0.0005 { // Exit if temperature is too low
			exit = true
		}

		if objectiveFunction.OF(cube) == 0 { // Exit if solution is found
			exit = true
		}

		newcube := cubeFuncs.FindSuccessor(cube)
		deltaE := objectiveFunction.OF(newcube) - objectiveFunction.OF(cube)

		if deltaE < 0 { // New cube has lower objective function value -> update current
			cube = newcube
		} else { // New cube has higher objective function value
			prob := math.Exp(-float64(deltaE) / T)
			if rand.Float64() < prob { // Accept new cube with random probability
				cube = newcube
			}
		}
	}
	return cube
}
