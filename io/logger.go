package io

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"path/filepath"
)

type GenerationLog struct {
	Instance   string
	Algorithm  string
	Run        int
	Best       float32
	Worst      float32
	Avg        float32
	Std        float32
	PopSize    int
	Generation int
	Px         float32
	Pm         float32
	Tournament int
	Elitism    int
	Cooling    float32
	InnerLoop  int
}

func LogGeneration(log GenerationLog, outputPath string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Check if file exists to write headers
	var file *os.File
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		file, err = os.Create(outputPath)
		if err != nil {
			return err
		}
		writer := csv.NewWriter(file)
		headers := []string{
			"Instance", "Algorithm", "Run", "Best", "Worst", "Avg", "Std",
			"pop_size", "gen", "Px", "Pm", "tournament", "Elitism", "cooling", "innerLoop",
		}
		if err := writer.Write(headers); err != nil {
			return err
		}
		writer.Flush()
	} else {
		file, err = os.OpenFile(outputPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		log.Instance,
		log.Algorithm,
		fmt.Sprintf("%d", log.Run),
		fmt.Sprintf("%.4f", log.Best),
		fmt.Sprintf("%.4f", log.Worst),
		fmt.Sprintf("%.4f", log.Avg),
		fmt.Sprintf("%.4f", log.Std),
		fmt.Sprintf("%d", log.PopSize),
		fmt.Sprintf("%d", log.Generation),
		fmt.Sprintf("%.4f", log.Px),
		fmt.Sprintf("%.4f", log.Pm),
		fmt.Sprintf("%d", log.Tournament),
		fmt.Sprintf("%d", log.Elitism),
		fmt.Sprintf("%.4f", log.Cooling),
		fmt.Sprintf("%d", log.InnerLoop),
	}

	return writer.Write(record)
}

// CalculateStats calculates statistics (best, worst, avg, std) from a slice of costs
func CalculateStats(costs []float32) (best, worst, avg, std float32) {
	if len(costs) == 0 {
		return 0, 0, 0, 0
	}

	best = costs[0]
	worst = costs[0]
	sum := float64(0)

	for _, cost := range costs {
		if cost < best {
			best = cost
		}
		if cost > worst {
			worst = cost
		}
		sum += float64(cost)
	}

	avg = float32(sum / float64(len(costs)))

	// Calculate standard deviation
	sumSquares := float64(0)
	for _, cost := range costs {
		diff := float64(cost - avg)
		sumSquares += diff * diff
	}
	std = float32(math.Sqrt(sumSquares / float64(len(costs))))

	return best, worst, avg, std
}
