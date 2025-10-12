package main

import (
	"fmt"

	"pwr.lab0/io"
	"pwr.lab0/model"
	"pwr.lab0/solver"
)

func main() {
	// Step 1: Parse the input file
	instance, err := io.ParseInstance("./data/mock_data.txt")
	if err != nil {
		fmt.Println("Error parsing instance:", err)
		return
	}

	// Step 2: Define GA parameters (user input)
	problem := model.Problem{
		Instance:       instance,
		ElitismCount:   2,
		MaxGenerations: 3,
		MutationRate:   0.05,
		PopulationSize: 5,
	}

	// Step 3: Create an instance of the Genetic Algorithm solver
	ga := solver.Genetic{Problem: problem}

	// Step 4: Run the Genetic Algorithm
	solution := ga.Run()

	// Step 5: Print the solution and its cost
	fmt.Println("Generated Route:", solution.Route.NodeIDs)
	fmt.Println("Route Cost:", solution.Cost)
}
