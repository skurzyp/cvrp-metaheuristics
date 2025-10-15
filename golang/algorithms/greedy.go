package algorithms

import (
	"math"

	"pwr.lab0/model"
)

type Greedy struct {
	Instance model.Instance
}

// Run executes the Greedy Algorithm and returns the best CVRP solution.
// If startingNodeID is 0 (depot), start from depot. Otherwise, start from the given customer.
func (g *Greedy) Run(startingNodeID int) model.FinalSolution {
	instance := g.Instance
	visited := make(map[int]bool)
	var route []int // full route including all nodes in order of visit
	var subRoutes []model.Route

	depotID := instance.GetDepotID()
	depot := instance.NodesMatrix[depotID]

	// Mark depot visited only if we start from depot
	if startingNodeID == depotID {
		visited[depot.ID] = true
		currentNode := depot
		currentLoad := 0
		currentSubRoute := []int{}

		for len(visited) < len(instance.NodesMatrix) {
			nextNode, found := g.findNearestUnvisited(currentNode, visited)
			if !found {
				break
			}

			// Check capacity
			if currentLoad+nextNode.Load <= instance.TruckMaxLoad {
				currentSubRoute = append(currentSubRoute, nextNode.ID)
				currentLoad += nextNode.Load
				visited[nextNode.ID] = true
				route = append(route, nextNode.ID)
				currentNode = nextNode
			} else {
				// Capacity exceeded, finish sub-route by returning to depot
				if len(currentSubRoute) > 0 {
					subRoutes = append(subRoutes, model.Route{NodeIDs: currentSubRoute})
				}
				currentSubRoute = []int{nextNode.ID}
				currentLoad = nextNode.Load
				visited[nextNode.ID] = true
				route = append(route, nextNode.ID)
				currentNode = nextNode
			}
		}

		if len(currentSubRoute) > 0 {
			subRoutes = append(subRoutes, model.Route{NodeIDs: currentSubRoute})
		}

	} else {
		// Start from customer node (not depot)
		currentNode := instance.NodesMatrix[startingNodeID]
		currentLoad := currentNode.Load
		currentSubRoute := []int{currentNode.ID}
		visited[currentNode.ID] = true
		visited[depot.ID] = true
		route = append(route, currentNode.ID)

		for len(visited) < len(instance.NodesMatrix) {
			nextNode, found := g.findNearestUnvisited(currentNode, visited)
			if !found {
				break
			}

			if currentLoad+nextNode.Load <= instance.TruckMaxLoad {
				currentSubRoute = append(currentSubRoute, nextNode.ID)
				currentLoad += nextNode.Load
				visited[nextNode.ID] = true
				route = append(route, nextNode.ID)
				currentNode = nextNode
			} else {
				// Capacity exceeded, finish sub-route by returning to depot
				if len(currentSubRoute) > 0 {
					subRoutes = append(subRoutes, model.Route{NodeIDs: currentSubRoute})
				}
				currentSubRoute = []int{nextNode.ID}
				currentLoad = nextNode.Load
				visited[nextNode.ID] = true
				route = append(route, nextNode.ID)
				currentNode = nextNode
			}
		}

		if len(currentSubRoute) > 0 {
			subRoutes = append(subRoutes, model.Route{NodeIDs: currentSubRoute})
		}
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

// findNearestUnvisited returns the nearest unvisited node and a boolean flag.
func (g *Greedy) findNearestUnvisited(current model.Node, visited map[int]bool) (model.Node, bool) {
	minDistance := float32(math.MaxFloat32)
	var nearest model.Node
	found := false

	for _, node := range g.Instance.NodesMatrix {
		if !visited[node.ID] {
			distance := g.Instance.GetNodesDistance(current, node)
			if distance < minDistance {
				minDistance = distance
				nearest = node
				found = true
			}
		}
	}

	return nearest, found
}
