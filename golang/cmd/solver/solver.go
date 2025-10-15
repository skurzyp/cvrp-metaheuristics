package main

import (
	"flag"
	"fmt"
	"os"

	"pwr.lab0/algorithms"
	"pwr.lab0/io"
	"pwr.lab0/model"
)

func main() {
	// Step 1: Define CLI flags
	filePath := flag.String("file", "./data/toy.vrp", "Path to the VRP instance file")
	popSize := flag.Int("pop", 200, "Population size (GA only)")
	maxGen := flag.Int("gen", 10000, "Maximum generations (GA only)")
	elitism := flag.Int("elitism", 10, "Number of elite solutions retained each generation (GA only)")
	mutation := flag.Float64("mutation", 0.05, "Mutation rate (0-1, GA only)")
	crossover := flag.Float64("crossover", 0.9, "Crossover rate (0-1, GA only)")
	tournament := flag.Int("tournament", 10, "Tournament size for parent selection (GA only)")
	algorithm := flag.String("algorithm", "ga", "Algorithm to use: 'ga' for Genetic Algorithm, 'sa' for Simulated Annealing, 'greedy' for Greedy Algorithm")
	starting := flag.Int("start", 0, "Starting node ID for Greedy Algorithm (default 0 - depot)")

	// Simulated Annealing parameters
	initialTemp := flag.Float64("temp", 2000.0, "Initial temperature (SA only)")
	minTemp := flag.Float64("minTemp", 0.001, "Minimum temperature (SA only)")
	coolingRate := flag.Float64("cooling", 0.999, "Cooling rate (0-1, SA only)")
	innerLoop := flag.Int("innerLoop", 2000, "Iterations per temperature step (SA only)")

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
		ga := algorithms.Genetic{Problem: problem}
		solution := ga.Run()

		// Step 5: Print solution
		fmt.Println(solution.String())

	case "sa":
		// Step 3: Define SA problem
		problem := model.Problem{
			Instance:           instance,
			InitialTemperature: *initialTemp,
			MinimumTemperature: *minTemp,
			CoolingRate:        *coolingRate,
			InnerLoop:          *innerLoop,
		}

		// Step 4: Run Simulated Annealing
		sa := algorithms.SimulatedAnnealing{Problem: problem}
		solution := sa.Run()

		// Step 5: Print solution
		fmt.Println(solution.String())

	case "greedy":
		// Step 3: Run Greedy algorithm
		greedy := algorithms.Greedy{Instance: instance}
		if *starting < 0 || *starting >= len(instance.NodesMatrix) {
			fmt.Printf("Starting node ID must be between 0 and %d\n", len(instance.NodesMatrix)-1)
			os.Exit(1)
		}
		solution := greedy.Run(*starting)
		fmt.Println(solution.String())

	default:
		fmt.Println("Unknown algorithm. Use -algorithm=ga, -algorithm=sa, or -algorithm=greedy")
		os.Exit(1)
	}
}
