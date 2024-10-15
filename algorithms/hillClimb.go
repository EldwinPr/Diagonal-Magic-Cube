package algorithms

import (
	"DiagonalMagicCube/cubeFuncs"
	"DiagonalMagicCube/objectiveFunction"
)

func SteepestAscentHillClimb(cube [5][5][5]int) [5][5][5]int {
	var newcube [5][5][5]int
	exit := false
	for !exit {
		newcube = cubeFuncs.FindBestSuccessor(cube)                      // find new successor
		if objectiveFunction.OF(newcube) >= objectiveFunction.OF(cube) { // no other successor higher than or equal to current -> exit
			exit = true
		} else { // new successor is better than current
			cube = newcube
		}
	}
	return cube
}

func HillClimbWithSidewaysMove(cube [5][5][5]int) [5][5][5]int {
	var newcube [5][5][5]int
	sidewaysCount := 0
	maxSidewaysMoves := 100
	exit := false
	for !exit {
		newcube = cubeFuncs.FindBestSuccessor(cube)
		if objectiveFunction.OF(newcube) > objectiveFunction.OF(cube) { // no other successor higher than current -> exit
			exit = true
		} else if objectiveFunction.OF(newcube) == objectiveFunction.OF(cube) && sidewaysCount < maxSidewaysMoves { // if new successor is equal to current and sideways moves are less than max
			sidewaysCount++ // increment sideways moves
			cube = newcube
		} else if objectiveFunction.OF(newcube) < objectiveFunction.OF(cube) { // new successor is better than current
			sidewaysCount = 0 // reset sideways moves
			cube = newcube
		} else {
			exit = true
		}
	}
	return cube
}

func StochasticHillClimb(cube [5][5][5]int) [5][5][5]int {
	for i := 0; i < 2300000; i++ { // repeat n times
		newcube := cubeFuncs.FindSuccessor(cube)
		if objectiveFunction.OF(newcube) < objectiveFunction.OF(cube) { // new successor is better than current -> update current
			cube = newcube
		}
	}
	return cube
}
