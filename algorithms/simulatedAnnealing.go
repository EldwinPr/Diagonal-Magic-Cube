package algorithms

import (
	"DiagonalMagicCube/cubeFuncs"
	"DiagonalMagicCube/objectiveFunction"
	"DiagonalMagicCube/types"
	"math"
	"math/rand"
	"time"
)

func SimulatedAnnealing(cube [5][5][5]int) types.AlgorithmResult {

	// Schedule function for temperature
	Schedule := func(T float64, iteration int) float64 {
		if iteration < 1000 {
			return T * 0.995 // Faster cooling at start
		} else if iteration < 5000 {
			return T * 0.999
		} else if iteration < 10000 {
			return T * 0.9995
		} else {
			return T * 0.99993
		}
	}

	// initialize results
	results := types.AlgorithmResult{
		Algorithm:   "Simulated Annealing",
		InitialCube: cube,
		InitialOF:   objectiveFunction.OF(cube),
		States:      make([]types.IterationState, 0),
	}

	T := rand.Float64() * 2000000000.0 // initial temperature

	// record initial state
	results.States = append(results.States, types.IterationState{
		Iteration:   0,
		Cube:        cube,
		OF:          objectiveFunction.OF(cube),
		Action:      "Initial",
		Temperature: T,
	})

	// initialize variables
	var newcube [5][5][5]int
	exit := false
	i := 1
	starttime := time.Now()

	for !exit {
		T = Schedule(T, i) // Update temperature
		i++
		if T < 0.0005 { // Exit if temperature is too low
			exit = true
		}

		newcube = cubeFuncs.FindSuccessor(cube)
		deltaE := objectiveFunction.OF(newcube) - objectiveFunction.OF(cube)

		prob := math.Exp(-float64(deltaE) / T)

		// Check if new cube has lower objective function value
		if deltaE < 0 {
			cube = newcube

			// record state
			results.States = append(results.States, types.IterationState{
				Iteration:   i,
				Cube:        cube,
				OF:          objectiveFunction.OF(cube),
				Action:      "Move",
				Temperature: T,
				Prob:        prob,
			})
			results.CustomVar++

		} else {

			// Check if new cube has higher objective function value
			if rand.Float64() < prob {
				cube = newcube

				// record state
				results.States = append(results.States, types.IterationState{
					Iteration:   i,
					Cube:        cube,
					OF:          objectiveFunction.OF(cube),
					Action:      "Backward Move",
					Temperature: T,
					Prob:        prob,
				})
			}
		}

	}

	// record final state
	results.States = append(results.States, types.IterationState{
		Iteration:   i,
		Cube:        cube,
		OF:          objectiveFunction.OF(cube),
		Action:      "Final State",
		Temperature: T,
		Prob:        0,
	})

	results.FinalCube = cube
	results.FinalOF = objectiveFunction.OF(cube)
	results.Duration = time.Since(starttime)

	return results
}
