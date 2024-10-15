package objectiveFunction

// objective function that calculates the sum of squares of differences of sums of diagonals of a cube

func SSDOF(cube [5][5][5]int) int {
	totalSum := 0
	target := 315

	// Helper function to calculate squared difference
	squaredDifference := func(sum, target int) int {
		diff := sum - target
		return diff * diff
	}

	// Calculate the sum of squares of differences for rows, columns, and pillars
	for z := 0; z < 5; z++ {
		// Rows and Columns in each layer
		for y := 0; y < 5; y++ {
			// Row sum
			rowSum := 0
			for x := 0; x < 5; x++ {
				rowSum += cube[z][y][x]
			}
			totalSum += squaredDifference(rowSum, target)
		}

		for x := 0; x < 5; x++ {
			// Column sum
			colSum := 0
			for y := 0; y < 5; y++ {
				colSum += cube[z][y][x]
			}
			totalSum += squaredDifference(colSum, target)
		}
	}

	// Pillars
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			pillarSum := 0
			for z := 0; z < 5; z++ {
				pillarSum += cube[z][y][x]
			}
			totalSum += squaredDifference(pillarSum, target)
		}
	}

	// Calculate the sum of squares of differences for diagonals of the cube faces
	for z := 0; z < 5; z++ {
		// Diagonals for each layer (z)
		diagSum1, diagSum2 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum1 += cube[z][i][i]
			diagSum2 += cube[z][i][4-i]
		}
		totalSum += squaredDifference(diagSum1, target)
		totalSum += squaredDifference(diagSum2, target)
	}

	for y := 0; y < 5; y++ {
		// Diagonals in y-direction across layers
		diagSum3, diagSum4 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum3 += cube[i][y][i]
			diagSum4 += cube[4-i][y][i]
		}
		totalSum += squaredDifference(diagSum3, target)
		totalSum += squaredDifference(diagSum4, target)
	}

	for x := 0; x < 5; x++ {
		// Diagonals in x-direction across layers
		diagSum5, diagSum6 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum5 += cube[i][i][x]
			diagSum6 += cube[4-i][i][x]
		}
		totalSum += squaredDifference(diagSum5, target)
		totalSum += squaredDifference(diagSum6, target)
	}

	// Calculate the sum of squares of differences for space diagonals
	spaceDiag1, spaceDiag2, spaceDiag3, spaceDiag4 := 0, 0, 0, 0
	for i := 0; i < 5; i++ {
		spaceDiag1 += cube[i][i][i]
		spaceDiag2 += cube[4-i][i][i]
		spaceDiag3 += cube[i][4-i][i]
		spaceDiag4 += cube[4-i][4-i][i]
	}
	totalSum += squaredDifference(spaceDiag1, target)
	totalSum += squaredDifference(spaceDiag2, target)
	totalSum += squaredDifference(spaceDiag3, target)
	totalSum += squaredDifference(spaceDiag4, target)

	return totalSum
}
