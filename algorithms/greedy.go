package algorithms

import (
	"math"

	"pwr.lab0/model"
)

type Greedy struct {
	Instance model.Instance
}

// Run executes the Greedy Algorithm and returns the best CVRP solution.
// It starts from a specified node and greedily visits nearest unvisited nodes,
// returning to depot when capacity is exceeded.
func (g *Greedy) Run(startingNodeID int) model.FinalSolution {
	instance := g.Instance
	visited := make(map[int]bool)
	var route []int // Only customer node IDs
	var subRoutes []model.Route

	depot := instance.NodesMatrix[instance.GetDepotID()]
	visited[depot.ID] = true

	currentNode := instance.NodesMatrix[startingNodeID]
	currentLoad := currentNode.Load
	currentSubRoute := []int{currentNode.ID}
	visited[currentNode.ID] = true
	route = append(route, currentNode.ID)

	for len(visited) < len(instance.NodesMatrix) {
		nextNode := model.Node{}
		minDistance := float32(math.MaxFloat32)
		found := false

		// Find the nearest unvisited node
		for _, node := range instance.NodesMatrix {
			if !visited[node.ID] {
				distance := instance.GetNodesDistance(currentNode, node)
				if distance < minDistance {
					minDistance = distance
					nextNode = node
					found = true
				}
			}
		}

		if !found {
			break
		}

		// Check if we can add this node to current sub-route
		if currentLoad+nextNode.Load <= instance.TruckMaxLoad {
			// Add to current sub-route
			currentSubRoute = append(currentSubRoute, nextNode.ID)
			currentLoad += nextNode.Load
			visited[nextNode.ID] = true
			route = append(route, nextNode.ID)
			currentNode = nextNode
		} else {
			// Capacity exceeded - finish current sub-route and start new one
			subRoutes = append(subRoutes, model.Route{NodeIDs: currentSubRoute})

			// Start new sub-route from depot to nextNode
			currentSubRoute = []int{nextNode.ID}
			currentLoad = nextNode.Load
			visited[nextNode.ID] = true
			route = append(route, nextNode.ID)
			currentNode = nextNode
		}
	}

	// Add the last sub-route
	if len(currentSubRoute) > 0 {
		subRoutes = append(subRoutes, model.Route{NodeIDs: currentSubRoute})
	}

	// Calculate total cost
	calculator := model.RouteCalculator{}
	fullRoute := model.Route{NodeIDs: route}
	cost := calculator.CalculateCost(instance, fullRoute)

	return model.FinalSolution{
		SubRoutes: subRoutes,
		Cost:      cost,
		Route:     fullRoute,
	}
}
