package main

import (
	"flag"
	"fmt"
	"os"

	"pwr.lab0/io"
	"pwr.lab0/model"
	"pwr.lab0/solver"
)

func main() {
	// Step 1: Define CLI flags
	filePath := flag.String("file", "./data/mock_data.txt", "Path to the VRP instance file")
	popSize := flag.Int("pop", 100, "Population size")
	maxGen := flag.Int("gen", 100, "Maximum generations")
	elitism := flag.Int("elitism", 5, "Number of elite solutions retained each generation")
	mutation := flag.Float64("mutation", 0.1, "Mutation rate (0-1)")
	crossover := flag.Float64("crossover", 0.7, "Crossover rate (0-1)")
	tournament := flag.Int("tournament", 5, "Tournament size for parent selection")
	algorithm := flag.String("algorithm", "ga", "Algotihm to use: 'ga' for Genetic Algorithm (default), 'sa' for Simulated Annealing (not implemented), greedy for Greedy Algorithm")
	starting := flag.Int("start", 1, "Starting node ID for Greedy Algorithm (default 1)")

	flag.Parse()

	// Step 2: Parse the input file
	instance, err := io.ParseInstance(*filePath)
	if err != nil {
		fmt.Println("Error parsing instance:", err)
		os.Exit(1)
	}

	switch *algorithm {
	case "ga":
		// Step 3: Define GA problem
		problem := model.Problem{
			Instance:       instance,
			ElitismCount:   *elitism,
			MaxGenerations: *maxGen,
			MutationRate:   float32(*mutation),
			CrossoverRate:  float32(*crossover),
			PopulationSize: *popSize,
			TournamentSize: *tournament,
		}

		// Step 4: Run GA
		ga := solver.Genetic{Problem: problem}
		solution := ga.Run()

		// Step 5: Print solution
		fmt.Println(solution.String())
	case "greedy":
		greedy := solver.Greedy{Instance: instance}
		if *starting < 1 || *starting >= len(instance.NodesMatrix) {
			fmt.Printf("Starting node ID must be between 1 and %d\n", len(instance.NodesMatrix)-1)
			os.Exit(1)
		}
		solution := greedy.Run(*starting)
		fmt.Println(solution.String())
	}

}
