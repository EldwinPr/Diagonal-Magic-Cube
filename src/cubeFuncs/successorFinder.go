package cubeFuncs

import (
	"DiagonalMagicCube/objectiveFunction"
	"math/rand"
)

func FindSuccessor(cube [5][5][5]int) [5][5][5]int {
	// Finds a successor of the cube by swapping two random positions

	// Select two random positions in the cube to swap
	i1, j1, k1 := rand.Intn(5), rand.Intn(5), rand.Intn(5)
	i2, j2, k2 := rand.Intn(5), rand.Intn(5), rand.Intn(5)

	// Ensure the two positions are different
	for i1 == i2 && j1 == j2 && k1 == k2 {
		i2, j2, k2 = rand.Intn(5), rand.Intn(5), rand.Intn(5)
	}

	// Swap the values at the two positions
	cube[i1][j1][k1], cube[i2][j2][k2] = cube[i2][j2][k2], cube[i1][j1][k1]

	return cube
}

func FindBestSuccessor(cube [5][5][5]int) [5][5][5]int {
	// Finds the best successor of the cube by checking all possible swaps (125*124)

	// Initialize the best cube and its objective function value
	bestCube := cube
	bestValue := objectiveFunction.OF(cube)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					for m := 0; m < 5; m++ {
						for n := 0; n < 5; n++ {
							if i != l || j != m || k != n {
								// Swap the values at the two positions
								cube[i][j][k], cube[l][m][n] = cube[l][m][n], cube[i][j][k]
								// Calculate the objective function value of the new cube
								value := objectiveFunction.OF(cube)
								// If the new cube has a better objective function value, update the best cube and its value
								if value < bestValue {
									bestCube = cube
									bestValue = value
								}
								// Swap the values back to restore the original cube
								cube[i][j][k], cube[l][m][n] = cube[l][m][n], cube[i][j][k]
							}
						}
					}
				}
			}
		}
	}

	return bestCube
}
