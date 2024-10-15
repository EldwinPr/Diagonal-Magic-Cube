package algorithms

import (
	"DiagonalMagicCube/cubeFuncs"
	"DiagonalMagicCube/objectiveFunction"
	"math"
)

func SimulatedAnnealing(cube [5][5][5]int) [5][5][5]int {
	T := 1000000000.0 // Initial temperature
	iteration := 1    // Iteration number

	// helper function to update temperature
	Schedule := func(T float64, iteration int) float64 {
		// Update temperature based on iteration number

		if iteration < 1000 {
			return T * 0.99
		} else if iteration < 10000 {
			return T * 0.9999
		} else if iteration < 100000 {
			return T * 0.99995
		} else {
			return T * 0.999993
		}
	}

	exit := false
	for !exit {
		T = Schedule(T, iteration) // Update temperature
		iteration++
		if T < 0.0005 { // Exit if temperature is too low
			exit = true
		}

		newcube := cubeFuncs.FindSuccessor(cube)
		deltaE := objectiveFunction.OF(newcube) - objectiveFunction.OF(cube)

		if deltaE < 0 { // New cube has lower objective function value -> update current
			cube = newcube
		} else { // New cube has higher objective function value
			prob := math.Exp(-float64(deltaE) / T)
			if 0.2 < prob { // Accept new cube with random probability
				cube = newcube
			}
		}
	}
	return cube
}
