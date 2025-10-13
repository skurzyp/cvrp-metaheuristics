// Package analytics provides statistical evaluation utilities
// to benchmark and compare CVRP algorithms.
package analytics

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"pwr.lab0/algorithms"
	"pwr.lab0/io"
	"pwr.lab0/model"
)

// ResultSummary stores aggregated results from multiple runs.
type ResultSummary struct {
	Algorithm    string
	InstanceName string
	Runs         int

	Best  float64
	Worst float64
	Avg   float64
	Std   float64

	Config map[string]any
}

// RunExperiment executes a given algorithm multiple times (e.g. GA, SA)
// concurrently and computes aggregated statistics.
func RunExperiment(problem model.Problem, algorithm string, runs int) ResultSummary {
	results := make([]float64, runs)
	var wg sync.WaitGroup
	var mu sync.Mutex // protects writes to results[]

	wg.Add(runs)
	for i := 0; i < runs; i++ {
		go func(i int) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())

			var cost float64
			switch algorithm {
			case "ga":
				problemCopy := problem // defensive copy for concurrent safety
				ga := algorithms.Genetic{Problem: problemCopy}
				sol := ga.Run()
				cost = float64(sol.Cost)

			case "sa":
				problemCopy := problem
				sa := algorithms.SimulatedAnnealing{Problem: problemCopy}
				sol := sa.Run()
				cost = float64(sol.Cost)

			case "greedy":
				startNode := (i % (len(problem.Instance.NodesMatrix) - 1)) + 1
				greedy := algorithms.Greedy{Instance: problem.Instance}
				sol := greedy.Run(startNode)
				cost = float64(sol.Cost)

			case "random":
				calculator := model.RouteCalculator{}
				route := algorithms.RandomRoute(len(problem.Instance.NodesMatrix))
				cost = float64(calculator.CalculateCost(problem.Instance, route))
			}

			// Thread-safe write
			mu.Lock()
			results[i] = cost
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	best, worst, avg, std := summarize(results)
	return ResultSummary{
		Algorithm:    algorithm,
		InstanceName: problem.Instance.Name,
		Runs:         runs,
		Best:         best,
		Worst:        worst,
		Avg:          avg,
		Std:          std,
		Config: map[string]any{
			"pop_size":   problem.PopulationSize,
			"gen":        problem.MaxGenerations,
			"Px":         problem.CrossoverRate,
			"Pm":         problem.MutationRate,
			"Tournament": problem.TournamentSize,
			"Elitism":    problem.ElitismCount,
			"cooling":    problem.CoolingRate,
			"innerLoop":  problem.InnerLoop,
		},
	}
}

// summarize computes min, max, mean, and standard deviation.
func summarize(values []float64) (best, worst, avg, std float64) {
	if len(values) == 0 {
		return 0, 0, 0, 0
	}
	best, worst = values[0], values[0]
	sum := 0.0
	for _, v := range values {
		sum += v
		if v < best {
			best = v
		}
		if v > worst {
			worst = v
		}
	}
	avg = sum / float64(len(values))

	var variance float64
	for _, v := range values {
		variance += (v - avg) * (v - avg)
	}
	std = math.Sqrt(variance / float64(len(values)))
	return
}

// SaveResults writes multiple summaries to a CSV file using the custom io.Writer.
func SaveResults(filePath string, summaries []ResultSummary) error {
	writer := io.NewCSVWriter(filePath)
	defer writer.Close()

	// Header
	writer.Write([]string{
		"Instance", "Algorithm", "Runs", "Best", "Worst", "Avg", "Std",
		"pop_size", "gen", "Px", "Pm", "tournament", "Elitism", "cooling", "innerLoop",
	})

	// Each experiment summary
	for _, s := range summaries {
		writer.Write([]string{
			s.InstanceName,
			s.Algorithm,
			fmt.Sprintf("%d", s.Runs),
			fmt.Sprintf("%.4f", s.Best),
			fmt.Sprintf("%.4f", s.Worst),
			fmt.Sprintf("%.4f", s.Avg),
			fmt.Sprintf("%.4f", s.Std),
			fmt.Sprintf("%v", s.Config["pop_size"]),
			fmt.Sprintf("%v", s.Config["gen"]),
			fmt.Sprintf("%v", s.Config["Px"]),
			fmt.Sprintf("%v", s.Config["Pm"]),
			fmt.Sprintf("%v", s.Config["Tournament"]),
			fmt.Sprintf("%v", s.Config["Elitism"]),
			fmt.Sprintf("%v", s.Config["cooling"]),
			fmt.Sprintf("%v", s.Config["innerLoop"]),
		})
	}

	return writer.Flush()
}
