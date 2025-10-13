// Package io provides functionality to parse input files and generate output files.
package io

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"pwr.lab0/model"
)

type section int

const (
	sectionNone section = iota
	sectionNode
	sectionDemand
	sectionDepot
)

func ParseInstance(filePath string) (model.Instance, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return model.Instance{}, err
	}
	defer file.Close()

	var nodes []model.Node
	var distanceMatrix [][]float32
	var truckMaxLoad int
	var depotFileID int // Depot ID from file (1-based)

	scanner := bufio.NewScanner(file)
	currentSection := sectionNone

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		// Switch sections
		switch parts[0] {
		case "NODE_COORD_SECTION":
			currentSection = sectionNode
			continue
		case "DEMAND_SECTION":
			currentSection = sectionDemand
			continue
		case "DEPOT_SECTION":
			currentSection = sectionDepot
			continue
		case "EOF":
			currentSection = sectionNone
			continue
		case "CAPACITY":
			truckMaxLoad, _ = strconv.Atoi(parts[len(parts)-1])
			continue
		}

		// Parse line based on section
		switch currentSection {
		case sectionNode:
			id, _ := strconv.Atoi(parts[0])
			x, _ := strconv.Atoi(parts[1])
			y, _ := strconv.Atoi(parts[2])
			nodes = append(nodes, model.Node{
				ID:       id - 1, // Temporary ID, will be reordered
				Position: model.Position{X: x, Y: y},
				Load:     0, // will be set later in DEMAND_SECTION
			})
		case sectionDemand:
			id, _ := strconv.Atoi(parts[0])
			load, _ := strconv.Atoi(parts[1])
			nodes[id-1].Load = load // NOTE: IDs in file are 1-based, slice is 0-based
		case sectionDepot:
			id, _ := strconv.Atoi(parts[0])
			if id == -1 {
				// End of depot section
				continue
			}
			depotFileID = id // Store depot ID (1-based)
		}
	}

	// Reorder nodes so depot is at index 0
	if depotFileID > 0 {
		depotIndex := depotFileID - 1 // Convert to 0-based
		if depotIndex != 0 {
			// Swap depot to position 0
			nodes[0], nodes[depotIndex] = nodes[depotIndex], nodes[0]
		}
	}

	// Reassign IDs to match array indices
	for i := range nodes {
		nodes[i].ID = i
	}

	// Build distance matrix
	n := len(nodes)
	distanceMatrix = make([][]float32, n)
	for i := range nodes {
		distanceMatrix[i] = make([]float32, n)
		for j := range nodes {
			distanceMatrix[i][j] = nodes[i].Position.CalculateDistance(nodes[i], nodes[j])
		}
	}

	return model.Instance{
		Name:           strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1],
		NodesMatrix:    nodes,
		DistanceMatrix: distanceMatrix,
		TruckMaxLoad:   truckMaxLoad,
	}, nil
}
