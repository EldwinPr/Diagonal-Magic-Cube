package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Stores results of a run
type AlgorithmResult struct {
	Algorithm      string           `json:"algorithm"`
	InitialCube    [5][5][5]int     `json:"initialCube"`
	FinalCube      [5][5][5]int     `json:"finalCube"`
	FinalOF        int              `json:"finalOF"`
	InitialOF      int              `json:"initialOF"`
	Duration       time.Duration    `json:"duration"`
	CustomVar      int              `json:"customVar"` // custom variable for HCWSM, RRHC, and SHC
	IterPerRestart []int            `json:"forRRHC"`
	States         []IterationState `json:"states"`
}

// Stores values of each iteration
type IterationState struct {
	Iteration   int          `json:"iteration"`
	Cube        [5][5][5]int `json:"cube"`
	OF          int          `json:"OF"`
	AvgOF       float64      `json:"avgOF"`
	Temperature float64      `json:"temperature"`
	Prob        float32      `json:"prob"`
	Population  int          `json:"population"`
	Action      string       `json:"action"`
}

func SaveEcxperimentResult(results AlgorithmResult, run int) error {

	// directory to store results
	var dirname string
	switch results.Algorithm {
	case "Steepest Ascent Hill Climb":
		dirname = "SAHC"
	case "Hill Climb with Sideways Moves":
		dirname = "HCWSM"
	case "Random Restart Hill Climb":
		dirname = "RRHC"
	case "Stochastic Hill Climb":
		dirname = "SHC"
	case "Simulated Annealing":
		dirname = "SA"
	}

	// filename and path
	filename := fmt.Sprintf("display/cubes/%s/%s_%d.json", dirname, results.Algorithm, run)

	// Marshal results to JSON
	file, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		return err
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write to file
	return ioutil.WriteFile(filename, file, 0644)
}
