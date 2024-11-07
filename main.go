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
	fmt.Println("Generate new results? (y/n)")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "y" {
		cube := cubeFuncs.MakeCube() // Create a 5x5x5 cube with values from 1 to 125
		timestart := time.Now()

		maxmoves, SHCiters, RRHCiters := 100, 2300000, 10

		fmt.Print("max amount of sideways moves: ")
		fmt.Scan(&maxmoves)
		fmt.Print("amount of iterations for SHC: ")
		fmt.Scan(&SHCiters)
		fmt.Print("amount of iterations for RRHC: ")
		fmt.Scan(&RRHCiters)

		// main loop
		for i := 0; i < 3; i++ {

			//randomize cube
			cube = cubeFuncs.RandomizeCube(cube)

			// SAHC
			results := algorithms.SteepestAscentHillClimb(cube)
			types.SaveEcxperimentResult(results, i+1)

			// HCWSM
			results = algorithms.HillClimbWithSidewaysMoves(cube, maxmoves)
			types.SaveEcxperimentResult(results, i+1)

			// RRHC
			results = algorithms.RandomRestartHillClimb(cube, RRHCiters)
			types.SaveEcxperimentResult(results, i+1)

			// SHC
			results = algorithms.StochasticHillClimb(cube, SHCiters)
			types.SaveEcxperimentResult(results, i+1)

			// SA
			results = algorithms.SimulatedAnnealing(cube)
			types.SaveEcxperimentResult(results, i+1)
		}
		fmt.Printf("Execution time: %v\n", time.Since(timestart))
		fmt.Println("All algorithms has been executed")
	}

	fs := http.FileServer(http.Dir("docs"))
	http.Handle("/", fs)

	fmt.Println("Server starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
