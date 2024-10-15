package main

import (
	"DiagonalMagicCube/algorithms"
	"DiagonalMagicCube/cubeFuncs"
	"DiagonalMagicCube/objectiveFunction"
	"fmt"
	"time"
)

func main() {
	cube := cubeFuncs.MakeCube()
	cube = cubeFuncs.RandomizeCube(cube)
	initial := cube
	fmt.Println("Initial objective function value:", objectiveFunction.OF(cube))

	// Steepest Ascent Hill Climb
	start := time.Now()
	cube = algorithms.SteepestAscentHillClimb(cube)
	fmt.Println("Time taken:", time.Since(start))
	fmt.Println("Final objective function value using Steepest Ascent HC:", objectiveFunction.OF(cube))

	// Hill Climb with Sideways Move
	cube = initial
	start = time.Now()
	cube = algorithms.HillClimbWithSidewaysMove(cube)
	fmt.Println("Time taken:", time.Since(start))
	fmt.Println("Final objective function value using HC with sideways move:", objectiveFunction.OF(cube))

	// Stochastic Hill Climb
	cube = initial
	start = time.Now()
	cube = algorithms.StochasticHillClimb(cube)
	fmt.Println("Time taken:", time.Since(start))
	fmt.Println("Final objective function value using Stochastic HC:", objectiveFunction.OF(cube))

	// Simulated Annealing
	cube = initial
	start = time.Now()
	cube = algorithms.SimulatedAnnealing(cube)
	fmt.Println("Time taken:", time.Since(start))
	fmt.Println("Final objective function value using Simulated Annealing:", objectiveFunction.OF(cube))

	// Genetic Algorithm
	cube = initial
	start = time.Now()
	cube = algorithms.GeneticAlgorithm(cube)
	fmt.Println("Time taken:", time.Since(start))
	fmt.Println("Final objective function value using Genetic Algorithm:", objectiveFunction.OF(cube))
}

// func main() {
// 	var initialCube = [5][5][5]int{
// 		{{25, 16, 80, 104, 90}, {115, 98, 4, 1, 97}, {42, 111, 85, 2, 75}, {66, 72, 27, 102, 48}, {67, 18, 119, 106, 5}},
// 		{{91, 77, 71, 6, 70}, {52, 64, 117, 69, 13}, {30, 118, 21, 123, 23}, {26, 39, 92, 44, 114}, {116, 17, 14, 73, 95}},
// 		{{47, 61, 45, 76, 86}, {107, 43, 38, 33, 94}, {89, 68, 63, 58, 37}, {32, 93, 88, 83, 19}, {40, 50, 81, 65, 79}},
// 		{{31, 53, 112, 109, 10}, {12, 82, 34, 87, 100}, {103, 3, 105, 8, 96}, {113, 57, 9, 62, 74}, {56, 120, 55, 49, 35}},
// 		{{121, 108, 7, 20, 59}, {29, 28, 122, 125, 11}, {51, 15, 41, 124, 84}, {78, 54, 99, 24, 60}, {36, 110, 46, 22, 101}},
// 	}
// 	fmt.Println("Initial objective function value:", objectiveFunction.OF(initialCube))
// }

// func main() {
// 	meanSAHC, meanHCWSM, meanSHC, meanSA := 0.0, 0.0, 0.0, 0.0
// 	mvSAHC, mvHCWSM, mvSHC, mvSA := 0, 0, 0, 0

// 	for i := 0; i < 10; i++ {
// 		// Randomize a new cube for each iteration
// 		initial := deepCopyCube(cubeFuncs.RandomizeCube(cubeFuncs.MakeCube()))

// 		// Steepest Ascent Hill Climb
// 		cube := deepCopyCube(initial)
// 		start := time.Now()
// 		cube = algorithms.SteepestAscentHillClimb(cube)
// 		meanSAHC += time.Since(start).Seconds()
// 		mvSAHC += objectiveFunction.OF(cube)

// 		// Hill Climb with Sideways Move
// 		cube = deepCopyCube(initial)
// 		start = time.Now()
// 		cube = algorithms.HillClimbWithSidewaysMove(cube)
// 		meanHCWSM += time.Since(start).Seconds()
// 		mvHCWSM += objectiveFunction.OF(cube)

// 		// Stochastic Hill Climb
// 		cube = deepCopyCube(initial)
// 		start = time.Now()
// 		cube = algorithms.StochasticHillClimb(cube)
// 		meanSHC += time.Since(start).Seconds()
// 		mvSHC += objectiveFunction.OF(cube)

// 		// Simulated Annealing
// 		cube = deepCopyCube(initial)
// 		start = time.Now()
// 		cube = algorithms.SimulatedAnnealing(cube)
// 		meanSA += time.Since(start).Seconds()
// 		mvSA += objectiveFunction.OF(cube)
// 	}

// 	fmt.Println("Mean time taken for Steepest Ascent Hill Climb:", meanSAHC/10)
// 	fmt.Println("Mean value of objective function for Steepest Ascent Hill Climb:", mvSAHC/10)
// 	fmt.Println("Mean time taken for Hill Climb with Sideways Move:", meanHCWSM/10)
// 	fmt.Println("Mean value of objective function for Hill Climb with Sideways Move:", mvHCWSM/10)
// 	fmt.Println("Mean time taken for Stochastic Hill Climb:", meanSHC/10)
// 	fmt.Println("Mean value of objective function for Stochastic Hill Climb:", mvSHC/10)
// 	fmt.Println("Mean time taken for Simulated Annealing:", meanSA/10)
// 	fmt.Println("Mean value of objective function for Simulated Annealing:", mvSA/10)
// }

// Assuming deepCopyCube is a utility function that creates a deep copy of the cube
// func deepCopyCube(cube [5][5][5]int) [5][5][5]int {
// 	newCube := cube
// 	return newCube
// }
