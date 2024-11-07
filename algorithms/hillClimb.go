package algorithms

import (
	"DiagonalMagicCube/cubeFuncs"
	"DiagonalMagicCube/objectiveFunction"
	"DiagonalMagicCube/types"
	"time"
)

func SteepestAscentHillClimb(cube [5][5][5]int) types.AlgorithmResult {

	// initialize results
	results := types.AlgorithmResult{
		Algorithm:   "Steepest Ascent Hill Climb",
		InitialCube: cube,
		InitialOF:   objectiveFunction.OF(cube),
		States:      make([]types.IterationState, 0),
	}

	// record initial state
	results.States = append(results.States, types.IterationState{
		Iteration: 0,
		Cube:      cube,
		OF:        objectiveFunction.OF(cube),
		Action:    "Initial",
	})

	// initialize variables
	var newcube [5][5][5]int
	exit := false
	i := 1
	starttime := time.Now()

	// main loop
	for !exit {
		// find best successor
		newcube = cubeFuncs.FindBestSuccessor(cube)

		// check if new successor is better than current
		if objectiveFunction.OF(newcube) >= objectiveFunction.OF(cube) {
			exit = true // exit if not

		} else { // continue if yes
			cube = newcube

			// record state
			results.States = append(results.States, types.IterationState{
				Iteration: i,
				Cube:      cube,
				OF:        objectiveFunction.OF(cube),
				Action:    "Move",
			})

			i++
		}
	}

	// record final state
	results.FinalCube = cube
	results.FinalOF = objectiveFunction.OF(cube)
	results.Duration = time.Since(starttime)

	return results
}

func HillClimbWithSidewaysMoves(cube [5][5][5]int, maxSidewaysMoves int) types.AlgorithmResult {

	// initialize results
	results := types.AlgorithmResult{
		Algorithm:   "Hill Climb with Sideways Moves",
		InitialCube: cube,
		InitialOF:   objectiveFunction.OF(cube),
		CustomVar:   maxSidewaysMoves,
		States:      make([]types.IterationState, 0),
	}

	// record initial state
	results.States = append(results.States, types.IterationState{
		Iteration: 0,
		Cube:      cube,
		OF:        objectiveFunction.OF(cube),
		Action:    "Initial",
	})

	// initialize variables
	var newcube [5][5][5]int
	sidewaysCount := 0
	exit := false
	i := 1
	starttime := time.Now()

	// main loop
	for !exit {
		// find best successor
		newcube = cubeFuncs.FindBestSuccessor(cube)

		// check if new successor is better than current
		if objectiveFunction.OF(newcube) > objectiveFunction.OF(cube) {
			exit = true // exit if not

		} else if objectiveFunction.OF(newcube) == objectiveFunction.OF(cube) && sidewaysCount < maxSidewaysMoves {
			sidewaysCount++
			cube = newcube

			// record state
			results.States = append(results.States, types.IterationState{
				Iteration: i,
				Cube:      cube,
				OF:        objectiveFunction.OF(cube),
				Action:    "Sideways Move",
			})

			i++

		} else if objectiveFunction.OF(newcube) < objectiveFunction.OF(cube) {
			sidewaysCount = 0
			cube = newcube

			// record state
			results.States = append(results.States, types.IterationState{
				Iteration: i,
				Cube:      cube,
				OF:        objectiveFunction.OF(cube),
				Action:    "Move",
			})

			i++

		} else {
			exit = true
		}
	}

	// record final state
	results.FinalCube = cube
	results.FinalOF = objectiveFunction.OF(cube)
	results.Duration = time.Since(starttime)

	return results
}

func StochasticHillClimb(cube [5][5][5]int, amount int) types.AlgorithmResult {

	// initialize results
	results := types.AlgorithmResult{
		Algorithm:   "Stochastic Hill Climb",
		InitialCube: cube,
		InitialOF:   objectiveFunction.OF(cube),
		CustomVar:   amount,
		States:      make([]types.IterationState, 0),
	}

	// record initial state
	results.States = append(results.States, types.IterationState{
		Iteration: 0,
		Cube:      cube,
		OF:        objectiveFunction.OF(cube),
		Action:    "Initial",
	})

	// initialize variables
	var newcube [5][5][5]int
	starttime := time.Now()

	// main loop
	for i := 0; i < amount; i++ { // repeat n times
		newcube = cubeFuncs.FindSuccessor(cube)
		if objectiveFunction.OF(newcube) < objectiveFunction.OF(cube) {
			cube = newcube

			// record state
			results.States = append(results.States, types.IterationState{
				Iteration: i,
				Cube:      cube,
				OF:        objectiveFunction.OF(cube),
				Action:    "Move",
			})
		}
	}

	// record final state
	results.FinalCube = cube
	results.FinalOF = objectiveFunction.OF(cube)
	results.Duration = time.Since(starttime)

	return results
}

func RandomRestartHillClimb(cube [5][5][5]int, amount int) types.AlgorithmResult {

	// initialize results
	results := types.AlgorithmResult{
		Algorithm:      "Random Restart Hill Climb",
		InitialCube:    cube,
		InitialOF:      objectiveFunction.OF(cube),
		CustomVar:      amount,
		States:         make([]types.IterationState, 0),
		IterPerRestart: make([]int, amount),
	}

	// record initial state
	results.States = append(results.States, types.IterationState{
		Iteration: 0,
		Cube:      cube,
		OF:        objectiveFunction.OF(cube),
		Action:    "Initial",
	})

	// initialize variables
	var newcube [5][5][5]int
	starttime := time.Now()
	bestcube := cube
	bestcubeOF := objectiveFunction.OF(cube)

	// random restart loop
	for i := 0; i < amount; i++ {
		if i > 0 {
			cube = cubeFuncs.RandomizeCube(cube)
			// Record restart state
			results.States = append(results.States, types.IterationState{
				Iteration: 0,
				Cube:      cube,
				OF:        objectiveFunction.OF(cube),
				Action:    "restart",
			})
		}

		// initialize variables
		exit := false
		j := 1

		// main loop (modified steepest ascent hill climb)
		for !exit {
			// find best successor
			newcube = cubeFuncs.FindBestSuccessor(cube)

			// check if new successor is better than current
			if objectiveFunction.OF(newcube) >= objectiveFunction.OF(cube) {
				exit = true // exit if not
				results.IterPerRestart[i] = j

			} else { // continue if yes
				cube = newcube

				// record state
				results.States = append(results.States, types.IterationState{
					Iteration: j,
					Cube:      cube,
					OF:        objectiveFunction.OF(cube),
					Action:    "Move",
				})

				j++
			}
		}

		if objectiveFunction.OF(cube) < bestcubeOF { // finds best cube
			bestcube = cube
			bestcubeOF = objectiveFunction.OF(cube)
		}
	}

	// record final state
	results.FinalCube = bestcube // returns best cubeq
	results.FinalOF = bestcubeOF
	results.Duration = time.Since(starttime)

	return results
}
