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
	starttime := time.Now()

	// main loop
	for !exit {
		// find best successor
		newcube = cubeFuncs.FindBestSuccessor(cube)

		// exit if new successor is not better than current
		if objectiveFunction.OF(newcube) >= objectiveFunction.OF(cube) {
			exit = true
		}

		// move to new successor
		cube = newcube

		// record state
		results.States = append(results.States, types.IterationState{
			Iteration: len(results.States),
			Cube:      cube,
			OF:        objectiveFunction.OF(cube),
			Action:    "Move",
		})

	}

	// record final state
	results.States = append(results.States, types.IterationState{
		Iteration: len(results.States),
		Cube:      cube,
		OF:        objectiveFunction.OF(cube),
		Action:    "Final state",
	})

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
	starttime := time.Now()

	// main loop
	for !exit {

		// find best successor
		newcube = cubeFuncs.FindBestSuccessor(cube)

		// check if new successor is better than current
		if objectiveFunction.OF(newcube) > objectiveFunction.OF(cube) {
			exit = true // exit if not

		} else if objectiveFunction.OF(newcube) == objectiveFunction.OF(cube) && sidewaysCount < maxSidewaysMoves { // sideways move
			sidewaysCount++
			cube = newcube

			// record state
			results.States = append(results.States, types.IterationState{
				Iteration: len(results.States),
				Cube:      cube,
				OF:        objectiveFunction.OF(cube),
				Action:    "Sideways Move",
			})

		} else if objectiveFunction.OF(newcube) < objectiveFunction.OF(cube) { // move to new successor
			sidewaysCount = 0
			cube = newcube

			// record state
			results.States = append(results.States, types.IterationState{
				Iteration: len(results.States),
				Cube:      cube,
				OF:        objectiveFunction.OF(cube),
				Action:    "Move",
			})

		} else {
			exit = true
		}
	}

	// record final state
	results.States = append(results.States, types.IterationState{
		Iteration: len(results.States),
		Cube:      cube,
		OF:        objectiveFunction.OF(cube),
		Action:    "Final state",
	})

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
	nomoves := 0
	maxnomoves := 2000
	lastIter := 20000
	starttime := time.Now()

	// main loop
	for i := 1; i < amount+1; i++ { // repeat n times
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

			nomoves = 0
		} else {
			nomoves++
			if nomoves == maxnomoves {
				break
			}
		}
		lastIter = i
	}

	// record final state
	results.States = append(results.States, types.IterationState{
		Iteration: lastIter,
		Cube:      cube,
		OF:        objectiveFunction.OF(cube),
		Action:    "Final state",
	})

	results.FinalCube = cube
	results.FinalOF = objectiveFunction.OF(cube)
	results.Duration = time.Since(starttime)

	return results
}

func RandomRestartHillClimb(cube [5][5][5]int, amount int) types.AlgorithmResult {

	// initialize results
	results := types.AlgorithmResult{
		Algorithm:   "Random Restart Hill Climb",
		InitialCube: cube,
		InitialOF:   objectiveFunction.OF(cube),
		CustomVar:   amount,
		States:      make([]types.IterationState, 0),
		CustomArr:   make([]int, amount),
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
		// initialize variables
		exit := false

		if i > 0 {
			cube = cubeFuncs.RandomizeCube(cube)
			// Record restart state
			results.States = append(results.States, types.IterationState{
				Iteration: len(results.States),
				Cube:      cube,
				OF:        objectiveFunction.OF(cube),
				Action:    "restart",
			})
		}

		// main loop (modified steepest ascent hill climb)
		for !exit {

			// find best successor
			newcube = cubeFuncs.FindBestSuccessor(cube)

			// exit if new successor is not better than current
			if objectiveFunction.OF(newcube) >= objectiveFunction.OF(cube) {
				exit = true

				// record iteration count for each restart
				results.CustomArr[i] = len(results.States)
			}

			cube = newcube // move to new successor

			// record state
			results.States = append(results.States, types.IterationState{
				Iteration: len(results.States),
				Cube:      cube,
				OF:        objectiveFunction.OF(cube),
				Action:    "Move",
			})
		}

		if objectiveFunction.OF(cube) < bestcubeOF { // finds best cube
			bestcube = cube
			bestcubeOF = objectiveFunction.OF(cube)
		}

		// record final restart state
		results.States = append(results.States, types.IterationState{
			Iteration: len(results.States),
			Cube:      cube,
			OF:        objectiveFunction.OF(cube),
			Action:    "Final state",
		})
	}

	// record final state
	results.FinalCube = bestcube // returns best cubeq
	results.FinalOF = bestcubeOF
	results.Duration = time.Since(starttime)

	return results
}
