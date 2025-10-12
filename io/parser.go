// Package io provides functionality to parse input files and generate output files.
package io

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"pwr.lab0/model"
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
	scanner := bufio.NewScanner(file)

	// Parse the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		switch parts[0] {
		case "CAPACITY":
			truckMaxLoad, _ = strconv.Atoi(parts[len(parts)-1])
		case "NODE_COORD_SECTION":
			for scanner.Scan() {
				nodeLine := scanner.Text()
				if nodeLine == "DEMAND_SECTION" {
					break
				}
				nodeParts := strings.Fields(nodeLine)
				id, _ := strconv.Atoi(nodeParts[0])
				x, _ := strconv.Atoi(nodeParts[1])
				y, _ := strconv.Atoi(nodeParts[2])
				nodes = append(nodes, model.Node{ID: id - 1, Position: model.Position{X: x, Y: y}})
			}
		case "DEMAND_SECTION":
			// Skip demand section for now
			for scanner.Scan() {
				if scanner.Text() == "DEPOT_SECTION" {
					break
				}
			}
		}
	}

	// Calculate distance matrix
	distanceMatrix = make([][]float32, len(nodes))
	for i := range nodes {
		distanceMatrix[i] = make([]float32, len(nodes))
		for j := range nodes {
			distanceMatrix[i][j] = nodes[i].Position.CalculateDistance(nodes[i], nodes[j])
		}
	}

	return model.Instance{
		NodesMatrix:    nodes,
		DistanceMatrix: distanceMatrix,
		TruckMaxLoad:   truckMaxLoad,
	}, nil
}
