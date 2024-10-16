package objectiveFunction

// placeholder for objective function

func OF(cube [5][5][5]int) int {
	// Absolute difference of sums of diagonals
	// return AbsDOF(cube)
	// Sum of squares of differences of sums of diagonals
	return SSDOF(cube)
}

func CheckForErrors(cube [5][5][5]int) (int, int, int, int, int) {
	row, column, pillar, diagface, diag := 0, 0, 0, 0, 0

	// Check for errors in rows, columns, and pillars
	for z := 0; z < 5; z++ {
		for y := 0; y < 5; y++ {
			rowSum, colSum, pillarSum := 0, 0, 0
			for x := 0; x < 5; x++ {
				rowSum += cube[z][y][x]
				colSum += cube[z][x][y]
				pillarSum += cube[x][y][z]
			}
			if rowSum != 315 {
				row++
			}
			if colSum != 315 {
				column++
			}
			if pillarSum != 315 {
				pillar++
			}
		}
	}

	// Check for errors in diagonals of the cube faces
	for z := 0; z < 5; z++ {
		// Diagonals for each layer (z)
		diagSum1, diagSum2 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum1 += cube[z][i][i]
			diagSum2 += cube[z][i][4-i]
		}
		if diagSum1 != 315 {
			diagface++
		}
		if diagSum2 != 315 {
			diagface++
		}
	}

	for y := 0; y < 5; y++ {
		// Diagonals in y-direction across layers
		diagSum3, diagSum4 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum3 += cube[i][y][i]
			diagSum4 += cube[4-i][y][i]
		}
		if diagSum3 != 315 {
			diagface++
		}
		if diagSum4 != 315 {
			diagface++
		}
	}

	for x := 0; x < 5; x++ {
		// Diagonals in x-direction across layers
		diagSum5, diagSum6 := 0, 0
		for i := 0; i < 5; i++ {
			diagSum5 += cube[i][i][x]
			diagSum6 += cube[4-i][i][x]
		}
		if diagSum5 != 315 {
			diagface++
		}
		if diagSum6 != 315 {
			diagface++
		}
	}

	// Check for errors in space diagonals
	spaceDiag1, spaceDiag2, spaceDiag3, spaceDiag4 := 0, 0, 0, 0
	for i := 0; i < 5; i++ {
		spaceDiag1 += cube[i][i][i]
		spaceDiag2 += cube[4-i][i][i]
		spaceDiag3 += cube[i][4-i][i]
		spaceDiag4 += cube[4-i][4-i][i]
	}
	if spaceDiag1 != 315 {
		diag++
	}
	if spaceDiag2 != 315 {
		diag++
	}
	if spaceDiag3 != 315 {
		diag++
	}
	if spaceDiag4 != 315 {
		diag++
	}

	return row, column, pillar, diagface, diag
}
