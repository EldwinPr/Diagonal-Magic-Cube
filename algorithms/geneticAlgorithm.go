package algorithms

import (
	"DiagonalMagicCube/cubeFuncs"
	"DiagonalMagicCube/objectiveFunction"
	"math/rand"
)

func GeneticAlgorithm(initialCube [5][5][5]int) [5][5][5]int {
	// Fixed parameters
	populationSize := 100
	maxGenerations := 1000

	// Initialize population
	population := make([][5][5][5]int, populationSize)
	population[0] = initialCube // Keep initial cube
	for i := 1; i < populationSize; i++ {
		population[i] = cubeFuncs.RandomizeCube(initialCube)
	}

	// Track best solution
	bestCube := initialCube
	bestFitness := objectiveFunction.OF(initialCube)

	// Main loop - generations
	for generation := 0; generation < maxGenerations; generation++ {
		// Calculate fitness for all cubes
		fitness := make([]int, populationSize)
		for i := 0; i < populationSize; i++ {
			fitness[i] = objectiveFunction.OF(population[i])
			if fitness[i] < bestFitness {
				bestFitness = fitness[i]
				bestCube = population[i]
			}
		}

		// Create new population
		newPopulation := make([][5][5][5]int, populationSize)
		newPopulation[0] = bestCube // Keep best cube (elitism)

		// Generate new individuals
		for i := 1; i < populationSize; i += 2 {
			// Select parents
			parent1 := tournamentSelect(population, fitness)
			parent2 := tournamentSelect(population, fitness)

			// Create children
			child1, child2 := crossover(parent1, parent2)

			// Mutate children (50% chance)
			if rand.Float64() < 0.5 {
				child1 = mutate(child1)
			}
			if rand.Float64() < 0.5 {
				child2 = mutate(child2)
			}

			// Add to new population
			newPopulation[i] = child1
			if i+1 < populationSize {
				newPopulation[i+1] = child2
			}
		}

		population = newPopulation
	}

	return bestCube
}

// Select best cube from random tournament
func tournamentSelect(population [][5][5][5]int, fitness []int) [5][5][5]int {
	best := rand.Intn(len(population))
	for i := 0; i < 3; i++ { // Tournament size of 3
		next := rand.Intn(len(population))
		if fitness[next] < fitness[best] {
			best = next
		}
	}
	return population[best]
}

// Swap some random positions between parents
func crossover(parent1, parent2 [5][5][5]int) ([5][5][5]int, [5][5][5]int) {
	child1 := parent1
	child2 := parent2

	numSwaps := rand.Intn(10) + 5 // 5-15 swaps
	for i := 0; i < numSwaps; i++ {
		x := rand.Intn(5)
		y := rand.Intn(5)
		z := rand.Intn(5)
		child1[x][y][z], child2[x][y][z] = child2[x][y][z], child1[x][y][z]
	}

	return child1, child2
}

// Swap some random positions in the cube
func mutate(cube [5][5][5]int) [5][5][5]int {
	numSwaps := rand.Intn(3) + 1 // 1-3 swaps
	for i := 0; i < numSwaps; i++ {
		x1, y1, z1 := rand.Intn(5), rand.Intn(5), rand.Intn(5)
		x2, y2, z2 := rand.Intn(5), rand.Intn(5), rand.Intn(5)
		cube[x1][y1][z1], cube[x2][y2][z2] = cube[x2][y2][z2], cube[x1][y1][z1]
	}
	return cube
}
