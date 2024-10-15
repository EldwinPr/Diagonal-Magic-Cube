package algorithms

// masih bau

import (
	"DiagonalMagicCube/cubeFuncs"
	"DiagonalMagicCube/objectiveFunction"
	"math/rand"
)

// GeneticAlgorithm runs a genetic algorithm on the given cube
func GeneticAlgorithm(initialCube [5][5][5]int) [5][5][5]int {

	populationSize := 200
	generations := 10000
	initialMutationRate := 0.4
	finalMutationRate := 0.1
	stagnationLimit := 500 // Limit to increase mutation rate if no improvement
	stagnationCount := 0
	bestFitness := int(^uint(0) >> 1) // Set to max int value initially

	// Initialize population
	population := make([][5][5][5]int, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = cubeFuncs.RandomizeCube(initialCube)
	}

	// Run generations
	for g := 0; g < generations; g++ {
		// Evaluate fitness of each individual
		fitness := make([]int, populationSize)
		fitnessSum := 0
		currentBestFitness := int(^uint(0) >> 1)
		for i, cube := range population {
			fitness[i] = objectiveFunction.OF(cube)
			fitnessSum += fitness[i]
			if fitness[i] == 0 {
				return cube // Found perfect solution
			}
			if fitness[i] < currentBestFitness {
				currentBestFitness = fitness[i]
			}
		}

		// Check for stagnation
		if currentBestFitness < bestFitness {
			bestFitness = currentBestFitness
			stagnationCount = 0
		} else {
			stagnationCount++
		}
		if stagnationCount > stagnationLimit {
			initialMutationRate += 0.1 // Increase mutation rate to introduce diversity
			if initialMutationRate > 0.8 {
				initialMutationRate = 0.8
			}
			stagnationCount = 0
		}

		// Preserve the best individual (elitism)
		bestIndex := 0
		bestFitness = fitness[0]
		for i := 1; i < populationSize; i++ {
			if fitness[i] < bestFitness {
				bestIndex = i
				bestFitness = fitness[i]
			}
		}
		bestIndividual := population[bestIndex]

		// Create new population
		newPopulation := make([][5][5][5]int, populationSize)
		newPopulation[0] = bestIndividual // Carry over the best individual

		// Adjust mutation rate dynamically
		mutationRate := initialMutationRate - (initialMutationRate-finalMutationRate)*float64(g)/float64(generations)

		// Create next generation using roulette wheel selection and crossover
		for i := 1; i < populationSize; i += 2 {
			parent1 := rouletteWheelSelection(population, fitness, fitnessSum)
			parent2 := rouletteWheelSelection(population, fitness, fitnessSum)

			child1, child2 := twoPointCrossover(parent1, parent2)

			// Apply mutation
			child1 = mutate(child1, mutationRate)
			child2 = mutate(child2, mutationRate)

			newPopulation[i] = child1
			if i+1 < populationSize {
				newPopulation[i+1] = child2
			}
		}

		// Introduce diversity by replacing some worst individuals with new random ones
		for i := populationSize - 10; i < populationSize; i++ {
			newPopulation[i] = cubeFuncs.RandomizeCube(initialCube)
		}

		population = newPopulation
	}

	// Return the best individual found
	best := population[0]
	bestFitness = objectiveFunction.OF(best)
	for _, cube := range population {
		currentFitness := objectiveFunction.OF(cube)
		if currentFitness < bestFitness {
			best = cube
			bestFitness = currentFitness
		}
	}

	return best
}

// Roulette wheel selection to choose a parent
func rouletteWheelSelection(population [][5][5][5]int, fitness []int, fitnessSum int) [5][5][5]int {
	selectionPoint := rand.Intn(fitnessSum)
	sum := 0
	for i, fit := range fitness {
		sum += fitnessSum - fit // Invert fitness to prioritize lower values
		if sum >= selectionPoint {
			return population[i]
		}
	}
	return population[len(population)-1]
}

// Two-point crossover combines two parents to create offspring
func twoPointCrossover(parent1, parent2 [5][5][5]int) ([5][5][5]int, [5][5][5]int) {
	var child1, child2 [5][5][5]int
	point1 := rand.Intn(5)
	point2 := rand.Intn(5)
	if point1 > point2 {
		point1, point2 = point2, point1
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				if i >= point1 && i <= point2 {
					child1[i][j][k] = parent2[i][j][k]
					child2[i][j][k] = parent1[i][j][k]
				} else {
					child1[i][j][k] = parent1[i][j][k]
					child2[i][j][k] = parent2[i][j][k]
				}
			}
		}
	}
	return child1, child2
}

// Mutate performs a random mutation on the given cube
func mutate(cube [5][5][5]int, mutationRate float64) [5][5][5]int {
	for i := range cube {
		for j := range cube[i] {
			for k := range cube[i][j] {
				if rand.Float64() < mutationRate {
					mutationAmount := rand.Intn(5) - 2 // Random change between -2 and 2
					cube[i][j][k] += mutationAmount
					if cube[i][j][k] < 1 {
						cube[i][j][k] = 1
					} else if cube[i][j][k] > 125 {
						cube[i][j][k] = 125
					}
				}
			}
		}
	}
	return cube
}
