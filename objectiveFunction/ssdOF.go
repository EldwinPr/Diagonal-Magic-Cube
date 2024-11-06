package objectiveFunction

// SSDOF calculates the sum of squares of differences of sums in a diagonal magic cube
func SSDOF(cube [5][5][5]int) int {
	totalSum := 0
	target := 315

	// Helper function to calculate squared difference
	squaredDifference := func(sum, target int) int {
		diff := sum - target
		return diff * diff
	}

	// Calculate sums for rows and columns within each layer
	for z := 0; z < 5; z++ {
		for y := 0; y < 5; y++ {
			rowSum, colSum := 0, 0
			for x := 0; x < 5; x++ {
				rowSum += cube[z][y][x]
				colSum += cube[z][x][y]
			}
			totalSum += squaredDifference(rowSum, target)
			totalSum += squaredDifference(colSum, target)
		}
	}

	// Calculate sums for pillars (summing over z)
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			pillarSum := 0
			for z := 0; z < 5; z++ {
				pillarSum += cube[z][y][x]
			}
			totalSum += squaredDifference(pillarSum, target)
		}
	}

	// Calculate the sum of squares of differences for face diagonals on each layer
	for z := 0; z < 5; z++ {
		diagSum1, diagSum2 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum1 += cube[z][i][i]
			diagSum2 += cube[z][i][4-i]
		}
		totalSum += squaredDifference(diagSum1, target)
		totalSum += squaredDifference(diagSum2, target)
	}

	// Diagonals across layers in the y-direction
	for y := 0; y < 5; y++ {
		diagSum3, diagSum4 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum3 += cube[i][y][i]
			diagSum4 += cube[4-i][y][i]
		}
		totalSum += squaredDifference(diagSum3, target)
		totalSum += squaredDifference(diagSum4, target)
	}

	// Diagonals across layers in the x-direction
	for x := 0; x < 5; x++ {
		diagSum5, diagSum6 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum5 += cube[i][i][x]
			diagSum6 += cube[4-i][i][x]
		}
		totalSum += squaredDifference(diagSum5, target)
		totalSum += squaredDifference(diagSum6, target)
	}

	// Main space diagonals
	diagSum7, diagSum8, diagSum9, diagSum10 := 0, 0, 0, 0
	for i := 0; i < 5; i++ {
		diagSum7 += cube[i][i][i]      // From (0,0,0) to (4,4,4)
		diagSum8 += cube[4-i][i][i]    // From (4,0,0) to (0,4,4)
		diagSum9 += cube[i][4-i][i]    // From (0,4,0) to (4,0,4)
		diagSum10 += cube[4-i][4-i][i] // From (4,4,0) to (0,0,4)
	}
	totalSum += squaredDifference(diagSum7, target)
	totalSum += squaredDifference(diagSum8, target)
	totalSum += squaredDifference(diagSum9, target)
	totalSum += squaredDifference(diagSum10, target)

	// Check central value (should be 63)
	centralValue := cube[2][2][2]
	totalSum += squaredDifference(centralValue, 63)

	// inner cube
	for i := 0; i < 5; i++ {
		if i != 2 { // Skip center position
			// X direction through center
			sum := cube[i][2][2] + cube[4-i][2][2] + centralValue
			totalSum += squaredDifference(sum, 189)

			// Y direction through center
			sum = cube[2][i][2] + cube[2][4-i][2] + centralValue
			totalSum += squaredDifference(sum, 189)

			// Z direction through center
			sum = cube[2][2][i] + cube[2][2][4-i] + centralValue
			totalSum += squaredDifference(sum, 189)
		}
	}

	// Diagonal lines through center
	sum1 := cube[0][0][0] + cube[4][4][4] + centralValue
	sum2 := cube[0][0][4] + cube[4][4][0] + centralValue
	sum3 := cube[0][4][0] + cube[4][0][4] + centralValue
	sum4 := cube[0][4][4] + cube[4][0][0] + centralValue

	totalSum += squaredDifference(sum1, 189)
	totalSum += squaredDifference(sum2, 189)
	totalSum += squaredDifference(sum3, 189)
	totalSum += squaredDifference(sum4, 189)

	return totalSum
}
