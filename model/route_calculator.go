package model

import "fmt"

type RouteCalculator struct{}

// SplitIntoSubRoutes splits a route into sub-routes based on truck capacity.
// This follows the greedy approach: visit nodes in order, return to depot when capacity exceeded.
func (r RouteCalculator) SplitIntoSubRoutes(instance Instance, route Route) []Route {
	var subRoutes []Route
	var currentSubRoute []int
	currentLoad := 0

	for _, nodeID := range route.NodeIDs {
		node := instance.NodesMatrix[nodeID]

		// If adding this node exceeds capacity, start a new sub-route
		if currentLoad+node.Load > instance.TruckMaxLoad {
			if len(currentSubRoute) > 0 {
				subRoutes = append(subRoutes, Route{NodeIDs: currentSubRoute})
			}
			currentLoad = 0
			currentSubRoute = []int{}
		}

		currentSubRoute = append(currentSubRoute, node.ID)
		currentLoad += node.Load
	}

	// Add the last sub-route if it has nodes
	if len(currentSubRoute) > 0 {
		subRoutes = append(subRoutes, Route{NodeIDs: currentSubRoute})
	}

	return subRoutes
}

func (r RouteCalculator) CalculateCost(instance Instance, route Route) float32 {
	fmt.Println("[ROUTE CALCULATOR] Calculating cost for route:", route)
	fmt.Println("rout length:", len(route.NodeIDs))
	if len(route.NodeIDs) == 0 {
		return 0
	}

	totalCost := float32(0)
	depot := instance.NodesMatrix[instance.GetDepotID()]
	currentLoad := 0

	// Start from first node
	currentNode := instance.NodesMatrix[route.NodeIDs[0]]
	currentLoad = currentNode.Load

	// Visit remaining nodes in order
	for i := 1; i < len(route.NodeIDs)-1; i++ {
		nextNode := instance.NodesMatrix[route.NodeIDs[i]]

		// Check if we need to return to depot due to capacity
		if currentLoad+nextNode.Load > instance.TruckMaxLoad {
			totalCost += instance.GetNodesDistance(currentNode, depot)
			totalCost += instance.GetNodesDistance(depot, nextNode)
			currentLoad = nextNode.Load
		} else {
			totalCost += instance.GetNodesDistance(currentNode, nextNode)
			currentLoad += nextNode.Load
		}

		currentNode = nextNode
	}

	// Route ends at route[n-1], no return to depot
	return totalCost
}
