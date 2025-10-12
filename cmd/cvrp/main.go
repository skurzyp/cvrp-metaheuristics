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
	popSize := flag.Int("pop", 5, "Population size")
	maxGen := flag.Int("gen", 10, "Maximum generations")
	elitism := flag.Int("elitism", 2, "Number of elite solutions retained each generation")
	mutation := flag.Float64("mutation", 0.05, "Mutation rate (0-1)")
	tournament := flag.Int("tournament", 3, "Tournament size for parent selection")

	flag.Parse()

	// Step 2: Parse the input file
	instance, err := io.ParseInstance(*filePath)
	if err != nil {
		fmt.Println("Error parsing instance:", err)
		os.Exit(1)
	}

	// Step 3: Define GA problem
	problem := model.Problem{
		Instance:       instance,
		ElitismCount:   *elitism,
		MaxGenerations: *maxGen,
		MutationRate:   float32(*mutation),
		PopulationSize: *popSize,
		TournamentSize: *tournament,
	}

	// Step 4: Run GA
	ga := solver.Genetic{Problem: problem}
	solution := ga.Run()

	// Step 5: Print solution
	fmt.Println(solution.String())
}
