package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"pwr.lab0/analytics"
	"pwr.lab0/io"
	"pwr.lab0/model"
)

func main() {
	// Basic benchmark parameters
	filePath := flag.String("file", "./data/toy.vrp", "Path to the VRP instance file")
	output := flag.String("out", "./results/summary.csv", "Path to CSV output file")
	runs := flag.Int("runs", 10, "Number of repetitions for metaheuristic algorithms - GA and SA")
	randomRuns := flag.Int("randomRuns", 10000, "Number of random solution generations")

	// Genetic Algorithm parameters (defaults aligned with solver)
	popSize := flag.Int("pop", 200, "Population size (GA only)")
	maxGen := flag.Int("gen", 10000, "Maximum generations (GA only)")
	elitism := flag.Int("elitism", 10, "Number of elite solutions retained (GA only)")
	mutation := flag.Float64("mutation", 0.05, "Mutation rate (0-1, GA only)")
	crossover := flag.Float64("crossover", 0.90, "Crossover rate (0-1, GA only)")
	tournament := flag.Int("tournament", 10, "Tournament size for parent selection (GA only)")

	// Simulated Annealing parameters (defaults aligned with solver)
	initialTemp := flag.Float64("temp", 2000.0, "Initial temperature (SA only)")
	minTemp := flag.Float64("minTemp", 0.001, "Minimum temperature (SA only)")
	coolingRate := flag.Float64("cooling", 0.999, "Cooling rate (0-1, SA only)")
	innerLoop := flag.Int("innerLoop", 2000, "Iterations per temperature step (SA only)")

	flag.Parse()

	instance, err := io.ParseInstance(*filePath)
	if err != nil {
		fmt.Println("Error parsing instance:", err)
		os.Exit(1)
	}

	problem := model.Problem{
		Instance:       instance,
		PopulationSize: *popSize,
		MaxGenerations: *maxGen,
		MutationRate:   float32(*mutation),
		CrossoverRate:  float32(*crossover),
		TournamentSize: *tournament,
		ElitismCount:   *elitism,

		InitialTemperature: *initialTemp,
		MinimumTemperature: *minTemp,
		CoolingRate:        *coolingRate,
		InnerLoop:          *innerLoop,
	}

	fmt.Println("🔍 Running benchmark suite (concurrent)...")

	resultsChan := make(chan analytics.ResultSummary, 4)
	var wg sync.WaitGroup

	// Helper to run one experiment concurrently
	run := func(name string, runs int) {
		defer wg.Done()
		start := time.Now()
		result := analytics.RunExperiment(problem, name, runs)
		elapsed := time.Since(start)
		fmt.Printf("✅ %s finished in %v\n", name, elapsed)
		resultsChan <- result
	}

	// Add all 4 algorithms
	wg.Add(4)
	go run("ga", *runs)
	go run("sa", *runs)
	go run("greedy", len(instance.NodesMatrix)-1) // one run per possible starting node
	go run("random", *randomRuns)                 // configurable number of random runs

	wg.Wait()
	close(resultsChan)

	// Collect results
	results := []analytics.ResultSummary{}
	for r := range resultsChan {
		results = append(results, r)
	}

	// Save results to CSV
	if err := analytics.SaveResults(*output, results); err != nil {
		fmt.Println("Error saving results:", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Benchmark complete! Results saved to %s\n", *output)
}
