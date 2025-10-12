package solver

import (
	"math"

	"pwr.lab0/model"
)

type Greedy struct {
	instance model.Instance
}

func (g *Greedy) Execute() ([]model.Node, float32) {
	instance := g.instance
	visited := make(map[int]bool)
	var route []model.Node
	var totalCost float32

	// Start from the depot
	currentNode := instance.NodesMatrix[instance.GetDepotID()]
	visited[currentNode.ID] = true
	route = append(route, currentNode)

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

	// Return to the depot
	totalCost += instance.GetNodesDistance(currentNode, instance.NodesMatrix[instance.GetDepotID()])
	route = append(route, instance.NodesMatrix[instance.GetDepotID()])

	return route, totalCost
}
