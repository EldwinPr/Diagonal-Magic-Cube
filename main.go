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

		maxmoves, SHCiters, RRHCiters := 100, 20000, 3

		// fmt.Print("max amount of sideways moves: ")
		// fmt.Scan(&maxmoves)
		// fmt.Print("amount of iterations for SHC: ")
		// fmt.Scan(&SHCiters)
		// fmt.Print("amount of iterations for RRHC: ")
		// fmt.Scan(&RRHCiters)
		cube1 := cubeFuncs.RandomizeCube(cube)
		cube2 := cubeFuncs.RandomizeCube(cube)
		cube3 := cubeFuncs.RandomizeCube(cube)

		// SAHC 1
		results := algorithms.SteepestAscentHillClimb(cube1)
		types.SaveEcxperimentResult(results, 1)
		// SAHC 2
		results = algorithms.SteepestAscentHillClimb(cube2)
		types.SaveEcxperimentResult(results, 2)
		// SAHC 3
		results = algorithms.SteepestAscentHillClimb(cube3)
		types.SaveEcxperimentResult(results, 3)

		// HCWSM 1
		results = algorithms.HillClimbWithSidewaysMoves(cube1, maxmoves)
		types.SaveEcxperimentResult(results, 1)
		// HCWSM 2
		results = algorithms.HillClimbWithSidewaysMoves(cube2, maxmoves)
		types.SaveEcxperimentResult(results, 2)
		// HCWSM 3
		results = algorithms.HillClimbWithSidewaysMoves(cube3, maxmoves)
		types.SaveEcxperimentResult(results, 3)

		// RRHC 1
		results = algorithms.RandomRestartHillClimb(cube1, RRHCiters)
		types.SaveEcxperimentResult(results, 1)
		// RRHC 2
		results = algorithms.RandomRestartHillClimb(cube2, RRHCiters)
		types.SaveEcxperimentResult(results, 2)
		// RRHC 3
		results = algorithms.RandomRestartHillClimb(cube3, RRHCiters)
		types.SaveEcxperimentResult(results, 3)

		// SHC 1
		results = algorithms.StochasticHillClimb(cube1, SHCiters)
		types.SaveEcxperimentResult(results, 1)
		// SHC 2
		results = algorithms.StochasticHillClimb(cube2, SHCiters)
		types.SaveEcxperimentResult(results, 2)
		// SHC 3
		results = algorithms.StochasticHillClimb(cube3, SHCiters)
		types.SaveEcxperimentResult(results, 3)

		// SA 1
		results = algorithms.SimulatedAnnealing(cube1)
		types.SaveEcxperimentResult(results, 1)
		// SA 2
		results = algorithms.SimulatedAnnealing(cube2)
		types.SaveEcxperimentResult(results, 2)
		// SA 3
		results = algorithms.SimulatedAnnealing(cube3)
		types.SaveEcxperimentResult(results, 3)

		fmt.Printf("Execution time: %v\n", time.Since(timestart))
		fmt.Println("All algorithms has been executed")
	}

	fs := http.FileServer(http.Dir("display"))
	http.Handle("/", fs)

	fmt.Println("Server starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
