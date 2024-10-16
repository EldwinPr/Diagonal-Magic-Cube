package objectiveFunction

// objective function that calculates the sum of squares of differences of sums of diagonals of a cube

func SSDOF(cube [5][5][5]int) int {
	totalSum := 0
	target := 315
	RCPMult := 2
	FDMult := 3
	SDMult := 1
	ACMult := 4
	auxiliaryTarget := 189

	// Helper function to calculate squared difference
	squaredDifference := func(sum, target int) int {
		diff := sum - target
		return diff * diff
	}

	// Calculate the sum of squares of differences for rows, columns, and pillars in a combined loop
	for z := 0; z < 5; z++ {
		for y := 0; y < 5; y++ {
			rowSum, colSum, pillarSum := 0, 0, 0
			for x := 0; x < 5; x++ {
				rowSum += cube[z][y][x]
				colSum += cube[z][x][y]
				pillarSum += cube[x][y][z]
			}
			totalSum += squaredDifference(rowSum, target) * RCPMult
			totalSum += squaredDifference(colSum, target) * RCPMult
			totalSum += squaredDifference(pillarSum, target) * RCPMult
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
		totalSum += squaredDifference(diagSum1, target) * FDMult
		totalSum += squaredDifference(diagSum2, target) * FDMult
	}

	for y := 0; y < 5; y++ {
		// Diagonals in y-direction across layers
		diagSum3, diagSum4 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum3 += cube[i][y][i]
			diagSum4 += cube[4-i][y][i]
		}
		totalSum += squaredDifference(diagSum3, target) * FDMult
		totalSum += squaredDifference(diagSum4, target) * FDMult
	}

	for x := 0; x < 5; x++ {
		// Diagonals in x-direction across layers
		diagSum5, diagSum6 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum5 += cube[i][i][x]
			diagSum6 += cube[4-i][i][x]
		}
		totalSum += squaredDifference(diagSum5, target) * FDMult
		totalSum += squaredDifference(diagSum6, target) * FDMult
	}

	// Calculate the sum of squares of differences for space diagonals
	spaceDiag1, spaceDiag2, spaceDiag3, spaceDiag4 := 0, 0, 0, 0
	for i := 0; i < 5; i++ {
		spaceDiag1 += cube[i][i][i]
		spaceDiag2 += cube[4-i][i][i]
		spaceDiag3 += cube[i][4-i][i]
		spaceDiag4 += cube[4-i][4-i][i]
	}
	totalSum += squaredDifference(spaceDiag1, target) * SDMult
	totalSum += squaredDifference(spaceDiag2, target) * SDMult
	totalSum += squaredDifference(spaceDiag3, target) * SDMult
	totalSum += squaredDifference(spaceDiag4, target) * SDMult

	// Calculate the sum of squares of differences for auxiliary 3x3 cube (centered at [2][2][2])
	// Rows, columns, and pillars passing through the center
	totalSum += squaredDifference(cube[1][2][2]+cube[2][2][2]+cube[3][2][2], auxiliaryTarget) * ACMult
	totalSum += squaredDifference(cube[2][1][2]+cube[2][2][2]+cube[2][3][2], auxiliaryTarget) * ACMult
	totalSum += squaredDifference(cube[2][2][1]+cube[2][2][2]+cube[2][2][3], auxiliaryTarget) * ACMult
	// Diagonals passing through the center
	totalSum += squaredDifference(cube[1][1][1]+cube[2][2][2]+cube[3][3][3], auxiliaryTarget) * ACMult
	totalSum += squaredDifference(cube[1][3][1]+cube[2][2][2]+cube[3][1][3], auxiliaryTarget) * ACMult
	totalSum += squaredDifference(cube[3][1][1]+cube[2][2][2]+cube[1][3][3], auxiliaryTarget) * ACMult
	totalSum += squaredDifference(cube[1][1][3]+cube[2][2][2]+cube[3][3][1], auxiliaryTarget) * ACMult

	return totalSum
}
