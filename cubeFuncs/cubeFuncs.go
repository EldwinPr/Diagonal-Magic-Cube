package cubeFuncs

import (
	"math/rand"
)

func MakeCube() [5][5][5]int {
	// Create a 5x5x5 cube with values from 1 to 125

	var cube [5][5][5]int
	num := 1
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				cube[i][j][k] = num
				num++
			}
		}
	}
	return cube
}

func RandomizeCube(cube [5][5][5]int) [5][5][5]int {
	values := rand.Perm(125) // Generate a permutation of numbers from 0 to 124 and puts it randomly in values[125]
	index := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				cube[i][j][k] = values[index] + 1 // Use values from 1 to 125
				index++
			}
		}
	}
	return cube
}
