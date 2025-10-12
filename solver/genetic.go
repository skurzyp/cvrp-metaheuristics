// Package solver implements algorithms solving the TSP/VRP problem using Genetic Algorithm
package solver

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"pwr.lab0/model"
)

type Genetic struct {
	Problem model.Problem
}

// Run executes the Genetic Algorithm and returns the best solution.
func (g Genetic) Run() model.FinalSolution {
	// Step 1: Initialize population
	population := g.initializePopulation()

	// Step 2: Evolve population for max generations
	for generation := 0; generation < g.Problem.MaxGenerations; generation++ {
		population = g.evolvePopulation(population)
	}

	// Return the best solution
	routeCalculator := model.RouteCalculator{}
	return model.FinalSolution{SubRoutes: routeCalculator.SplitIntoSubRoutes(g.Problem.Instance, population[0].Route), Cost: population[0].Cost, Route: population[0].Route}
}

// initializePopulation generates the initial population with one greedy solution and the rest random.
func (g Genetic) initializePopulation() []model.Solution {
	calculator := model.RouteCalculator{}
	population := make([]model.Solution, g.Problem.PopulationSize)

	// Add one greedy solution
	greedySolver := Greedy{instance: g.Problem.Instance}
	greedyRoute, _ := greedySolver.Execute()
	population[0] = model.Solution{
		Route: model.Route{NodeIDs: extractNodeIDs(greedyRoute)},
		Cost:  calculator.CalculateWithGreedy(g.Problem.Instance, model.Route{NodeIDs: extractNodeIDs(greedyRoute)}),
	}

	// Add random solutions
	for i := 1; i < g.Problem.PopulationSize; i++ {
		randomRoute := g.generateRandomRoute()
		population[i] = model.Solution{
			Route: randomRoute,
			Cost:  calculator.CalculateWithGreedy(g.Problem.Instance, randomRoute),
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

	// Step 2: Selection and mutation
	for len(nextGeneration) < g.Problem.PopulationSize {
		parent := g.selectParent(population)

		child := parent.Route
		if rand.Float32() < g.Problem.MutationRate {
			child = g.mutate(child)
			fmt.Println("Mutation occurred on route:", child.NodeIDs)
		}

		// Calculate cost for the child
		calculator := model.RouteCalculator{}
		childCost := calculator.CalculateWithGreedy(g.Problem.Instance, child)

		nextGeneration = append(nextGeneration, model.Solution{Route: child, Cost: childCost})
	}

	// Sort the new generation by cost
	sort.Slice(nextGeneration, func(i, j int) bool {
		return nextGeneration[i].Cost < nextGeneration[j].Cost
	})

	return nextGeneration
}

// generateRandomRoute creates a random route visiting all nodes.
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
	tournamentSize := 3 // TODO: Make this a parameter
	best := population[rand.Intn(len(population))]
	for i := 1; i < tournamentSize; i++ {
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
		return route // nothing to swap
	}
	i := rand.Intn(size)
	j := rand.Intn(size)
	route.NodeIDs[i], route.NodeIDs[j] = route.NodeIDs[j], route.NodeIDs[i]
	return route
}

// extractNodeIDs converts a slice of Nodes to a slice of their IDs.
func extractNodeIDs(nodes []model.Node) []int {
	nodeIDs := make([]int, len(nodes))
	for i, node := range nodes {
		nodeIDs[i] = node.ID
	}
	return nodeIDs
}
