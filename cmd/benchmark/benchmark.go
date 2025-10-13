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
	filePath := flag.String("file", "./data/mock_data.txt", "Path to the VRP instance file")
	output := flag.String("out", "./results/summary.csv", "Path to CSV output file")
	runs := flag.Int("runs", 10, "Number of repetitions for metaheuristic algorithms - GA and SA")
	flag.Parse()

	instance, err := io.ParseInstance(*filePath)
	if err != nil {
		fmt.Println("Error parsing instance:", err)
		os.Exit(1)
	}

	problem := model.Problem{
		Instance:       instance,
		PopulationSize: 200,
		MaxGenerations: 5000,
		MutationRate:   0.05,
		CrossoverRate:  0.90,
		TournamentSize: 10,
		ElitismCount:   10,

		InitialTemperature: 1000,
		MinimumTemperature: 0.001,
		CoolingRate:        0.9995,
		InnerLoop:          400,
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
	go run("random", 10000)                       // 10,000 random runs

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
