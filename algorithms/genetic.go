// Package algorithms implements algorithms solving the TSP/VRP problem using Genetic Algorithm
package algorithms

import (
	"math/rand"
	"sort"
	"time"

	"pwr.lab0/model"
)

type Genetic struct {
	Problem model.Problem
}

// Run executes the Genetic Algorithm and returns the best CVRP solution.
func (g Genetic) Run() model.FinalSolution {
	// Step 1: Initialize population
	population := g.initializePopulation()

	// Step 2: Evolve population for max generations
	for generation := 0; generation < g.Problem.MaxGenerations; generation++ {
		population = g.evolvePopulation(population)
	}

	// Return the best solution
	routeCalculator := model.RouteCalculator{}
	bestRoute := population[0].Route

	return model.FinalSolution{
		SubRoutes: routeCalculator.SplitIntoSubRoutes(g.Problem.Instance, bestRoute),
		Cost:      population[0].Cost,
		Route:     bestRoute,
	}
}

// initializePopulation generates the initial population with one greedy solution and the rest random.
func (g Genetic) initializePopulation() []model.Solution {
	calculator := model.RouteCalculator{}
	population := make([]model.Solution, g.Problem.PopulationSize)

	// Add one greedy solution (starting from a random non-depot node)
	greedySolver := Greedy{Instance: g.Problem.Instance}
	startNode := rand.Intn(len(g.Problem.Instance.NodesMatrix)-1) + 1 // Random node from 1 to n-1
	greedySolution := greedySolver.Run(startNode)

	population[0] = model.Solution{
		Route: greedySolution.Route,
		Cost:  greedySolution.Cost,
	}

	// Add random solutions
	for i := 1; i < g.Problem.PopulationSize; i++ {
		randomRoute := g.generateRandomRoute()
		population[i] = model.Solution{
			Route: randomRoute,
			Cost:  calculator.CalculateCost(g.Problem.Instance, randomRoute),
		}
	}

	// Sort population by cost (ascending)
	sort.Slice(population, func(i, j int) bool {
		return population[i].Cost < population[j].Cost
	})

	return population
}

// evolvePopulation performs selection and mutation to create the next generation.
func (g Genetic) evolvePopulation(population []model.Solution) []model.Solution {
	nextGeneration := make([]model.Solution, 0, g.Problem.PopulationSize)

	// Step 1: Elitism - retain the best individuals
	nextGeneration = append(nextGeneration, population[:g.Problem.ElitismCount]...)

	// Step 2: Selection, crossover, and mutation
	for len(nextGeneration) < g.Problem.PopulationSize {
		parent1 := g.selectParent(population)
		parent2 := g.selectParent(population)

		var child model.Route
		if rand.Float32() < g.Problem.CrossoverRate {
			child = g.crossoverOX(parent1.Route, parent2.Route)
		} else {
			child = parent1.Route
		}

		// Mutation
		if rand.Float32() < g.Problem.MutationRate {
			child = g.mutate(child)
		}

		calculator := model.RouteCalculator{}
		childCost := calculator.CalculateCost(g.Problem.Instance, child)
		nextGeneration = append(nextGeneration, model.Solution{Route: child, Cost: childCost})
	}

	// Sort the new generation by cost
	sort.Slice(nextGeneration, func(i, j int) bool {
		return nextGeneration[i].Cost < nextGeneration[j].Cost
	})

	return nextGeneration
}

// generateRandomRoute creates a random route visiting all nodes (excluding depot at index 0).
func (g Genetic) generateRandomRoute() model.Route {
	rand.Seed(time.Now().UnixNano())
	nodeIDs := make([]int, len(g.Problem.Instance.NodesMatrix)-1)
	for i := 1; i < len(g.Problem.Instance.NodesMatrix); i++ {
		nodeIDs[i-1] = i
	}
	rand.Shuffle(len(nodeIDs), func(i, j int) { nodeIDs[i], nodeIDs[j] = nodeIDs[j], nodeIDs[i] })
	return model.Route{NodeIDs: nodeIDs}
}

// selectParent selects a parent using tournament selection.
func (g Genetic) selectParent(population []model.Solution) model.Solution {
	best := population[rand.Intn(len(population))]
	for i := 1; i < g.Problem.TournamentSize; i++ {
		competitor := population[rand.Intn(len(population))]
		if competitor.Cost < best.Cost {
			best = competitor
		}
	}
	return best
}

// mutate performs swap mutation on a route.
func (g Genetic) mutate(route model.Route) model.Route {
	size := len(route.NodeIDs)
	if size < 2 {
		return route
	}
	i := rand.Intn(size)
	j := rand.Intn(size)
	route.NodeIDs[i], route.NodeIDs[j] = route.NodeIDs[j], route.NodeIDs[i]
	return route
}

// crossoverOX performs the Order Crossover (OX) between two parent routes.
func (g Genetic) crossoverOX(parent1, parent2 model.Route) model.Route {
	size := len(parent1.NodeIDs)

	if size <= 1 {
		return parent1
	}

	child := make([]int, size)
	for i := range child {
		child[i] = -1
	}

	// Choose two random crossover points
	start := rand.Intn(size)
	end := rand.Intn(size)
	if start > end {
		start, end = end, start
	}
	if start == end {
		end = (end + 1) % size
		if end < start {
			start, end = end, start
		}
	}

	// Copy the slice from Parent 1
	used := make(map[int]bool)
	for i := start; i <= end; i++ {
		child[i] = parent1.NodeIDs[i]
		used[parent1.NodeIDs[i]] = true
	}

	// Fill remaining slots from Parent 2 in order
	childIndex := (end + 1) % size
	for _, gene := range parent2.NodeIDs {
		if !used[gene] {
			child[childIndex] = gene
			used[gene] = true
			childIndex = (childIndex + 1) % size
		}
	}

	return model.Route{NodeIDs: child}
}
