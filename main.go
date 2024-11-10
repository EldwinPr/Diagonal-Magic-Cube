package main

import (
	"DiagonalMagicCube/algorithms"
	"DiagonalMagicCube/cubeFuncs"
	"DiagonalMagicCube/types"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	// Clear results
	if len(os.Args) > 1 && os.Args[1] == "clear" {
		os.RemoveAll("display/cubes")
	}

	// Ask user if they want to generate new results
	fmt.Println("Generate new results? (y/n)")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "y" {
		cube := cubeFuncs.MakeCube() // Create a 5x5x5 cube with values from 1 to 125
		timestart := time.Now()

		maxmoves, SHCiters, RRHCiters := 100, 20000, 3 // default values

		// GA
		population1, population2, population3 := 100, 150, 200
		maxGenerations1, maxGenerations2, maxGenerations3 := 5000, 6000, 7000

		// fmt.Print("max amount of sideways moves: ")
		// fmt.Scan(&maxmoves)
		// fmt.Print("amount of iterations for SHC: ")
		// fmt.Scan(&SHCiters)
		// fmt.Print("amount of iterations for RRHC: ")
		// fmt.Scan(&RRHCiters)
		// fmt.Print("population 1 size for GA: ")
		// fmt.Scan(&population1)
		// fmt.Print("population 2 size for GA: ")
		// fmt.Scan(&population2)
		// fmt.Print("population  3size for GA: ")
		// fmt.Scan(&population3)
		// fmt.Print("max amount of generations for GA: ")
		// fmt.Scan(&maxGenerations)

		for i := 1; i < 4; i++ {
			cube := cubeFuncs.RandomizeCube(cube)

			// SAHC
			results := algorithms.SteepestAscentHillClimb(cube)
			types.SaveEcxperimentResult(results, i)

			// HCWSM
			results = algorithms.HillClimbWithSidewaysMoves(cube, maxmoves)
			types.SaveEcxperimentResult(results, i)

			// RRHC
			results = algorithms.RandomRestartHillClimb(cube, RRHCiters)
			types.SaveEcxperimentResult(results, i)

			// SHC
			results = algorithms.StochasticHillClimb(cube, SHCiters)
			types.SaveEcxperimentResult(results, i)

			// SA
			results = algorithms.SimulatedAnnealing(cube)
			types.SaveEcxperimentResult(results, i)
		}

		// GA
		results := algorithms.GeneticAlgorithm(cube, population1, maxGenerations1)
		types.SaveEcxperimentResult(results, 1)
		results = algorithms.GeneticAlgorithm(cube, population2, maxGenerations1)
		types.SaveEcxperimentResult(results, 2)
		results = algorithms.GeneticAlgorithm(cube, population3, maxGenerations1)
		types.SaveEcxperimentResult(results, 3)
		results = algorithms.GeneticAlgorithm(cube, population1, maxGenerations2)
		types.SaveEcxperimentResult(results, 4)
		results = algorithms.GeneticAlgorithm(cube, population1, maxGenerations3)
		types.SaveEcxperimentResult(results, 5)

		fmt.Printf("Execution time: %v\n", time.Since(timestart))
		fmt.Println("All algorithms has been executed")
	}

	// GA

	fs := http.FileServer(http.Dir("display"))
	http.Handle("/", fs)

	fmt.Println("Server starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
