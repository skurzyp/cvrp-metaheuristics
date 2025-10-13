package algorithms

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"pwr.lab0/io"
	"pwr.lab0/model"
)

type SimulatedAnnealing struct {
	Problem model.Problem
}

// Run executes the Simulated Annealing algorithm and returns the best CVRP solution.
func (sa SimulatedAnnealing) Run() model.FinalSolution {
	calculator := model.RouteCalculator{}
	rand.Seed(time.Now().UnixNano())

	// Step 1: Initialize solution (greedy or random)
	greedySolver := Greedy{Instance: sa.Problem.Instance}
	startNode := rand.Intn(len(sa.Problem.Instance.NodesMatrix)-1) + 1
	initialSolution := greedySolver.Run(startNode)

	currentRoute := initialSolution.Route
	currentCost := initialSolution.Cost
	bestRoute := currentRoute
	bestCost := currentCost

	// Step 2: Define annealing parameters
	T := sa.Problem.InitialTemperature        // e.g. 1000.0
	alpha := sa.Problem.CoolingRate           // e.g. 0.995
	minT := sa.Problem.MinimumTemperature     // e.g. 0.001
	iterationsPerTemp := sa.Problem.InnerLoop // e.g. 100

	// Generate log file path
	logPath := "results/sa_progress_" +
		sa.Problem.Instance.Name +
		"_T" + fmt.Sprintf("%.2f", T) +
		"_alpha" + fmt.Sprintf("%.3f", alpha) +
		"_minT" + fmt.Sprintf("%.3f", minT) +
		"_inner%d.csv"
	logPath = fmt.Sprintf(logPath, iterationsPerTemp)

	// Log initial state
	log := io.GenerationLog{
		Instance:   sa.Problem.Instance.Name,
		Algorithm:  "sa",
		Run:        1, // This should be passed as parameter if needed
		Best:       bestCost,
		Worst:      currentCost,
		Avg:        currentCost,
		Std:        0,
		PopSize:    0,
		Generation: 0,
		Px:         0,
		Pm:         0,
		Tournament: 0,
		Elitism:    0,
		Cooling:    float32(alpha),
		InnerLoop:  iterationsPerTemp,
	}
	io.LogGeneration(log, logPath)

	generation := 1
	// Step 3: Annealing loop
	for T > minT {
		worstCost := currentCost
		avgCost := currentCost
		std := float32(0)

		for range iterationsPerTemp {
			// Generate neighbor solution (e.g. 2-swap)
			newRoute := sa.perturb(currentRoute)
			newCost := calculator.CalculateCost(sa.Problem.Instance, newRoute)

			delta := newCost - currentCost
			if delta < 0 || math.Exp(-float64(delta)/float64(T)) > rand.Float64() {
				// Accept new solution
				currentRoute = newRoute
				currentCost = newCost
			}

			// Track global best and statistics
			if currentCost < bestCost {
				bestRoute = currentRoute
				bestCost = currentCost
			}
			if currentCost > worstCost {
				worstCost = currentCost
			}
			avgCost = (avgCost + currentCost) / 2
		}

		// Log current temperature iteration
		log.Generation = generation
		log.Best = bestCost
		log.Worst = worstCost
		log.Avg = avgCost
		log.Std = std
		io.LogGeneration(log, logPath)

		generation++
		T *= alpha // Cool down
	}

	// Step 4: Return final solution
	subRoutes := calculator.SplitIntoSubRoutes(sa.Problem.Instance, bestRoute)
	return model.FinalSolution{
		SubRoutes: subRoutes,
		Cost:      bestCost,
		Route:     bestRoute,
	}
}

// perturb generates a neighbor solution by swapping two nodes.
func (sa SimulatedAnnealing) perturb(route model.Route) model.Route {
	newRoute := make([]int, len(route.NodeIDs))
	copy(newRoute, route.NodeIDs)

	i := rand.Intn(len(newRoute))
	j := rand.Intn(len(newRoute))
	newRoute[i], newRoute[j] = newRoute[j], newRoute[i]

	return model.Route{NodeIDs: newRoute}
}
