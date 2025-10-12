package solver

import (
	"math"

	"pwr.lab0/model"
)

type Greedy struct {
	instance model.Instance
}

func (g *Greedy) Execute(startingNodeID int) ([]model.Node, float32) {
	instance := g.instance
	visited := make(map[int]bool)
	var route []model.Node
	var totalCost float32

	// Mark depot as visited (we don't want to include it in the route)
	depotID := instance.GetDepotID()
	visited[depotID] = true

	// Start from the specified node
	currentNode := instance.NodesMatrix[startingNodeID]
	visited[currentNode.ID] = true
	route = append(route, currentNode)

	// Continue visiting remaining nodes
	for len(visited) < len(instance.NodesMatrix) {
		nextNode := model.Node{}
		minDistance := float32(math.MaxFloat32)

		// Find the nearest unvisited node
		for _, node := range instance.NodesMatrix {
			if !visited[node.ID] {
				distance := instance.GetNodesDistance(currentNode, node)
				if distance < minDistance {
					minDistance = distance
					nextNode = node
				}
			}
		}

		// Visit the next node
		visited[nextNode.ID] = true
		route = append(route, nextNode)
		totalCost += minDistance
		currentNode = nextNode
	}

	return route, totalCost
}
