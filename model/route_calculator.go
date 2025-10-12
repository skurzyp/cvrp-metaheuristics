package model

type RouteCalculator struct{}

// SplitIntoSubRoutes splits a route into sub-routes based on truck capacity.
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

// CalculateWithGreedy evaluates the cost of a route.
func (r RouteCalculator) CalculateWithGreedy(instance Instance, route Route) float32 {
	totalCost := float32(0)
	subRoutes := r.SplitIntoSubRoutes(instance, route)
	depot := instance.NodesMatrix[instance.GetDepotID()]

	for _, sub := range subRoutes {
		// Start from depot to first node
		firstNode := instance.NodesMatrix[sub.NodeIDs[0]]
		totalCost += instance.GetNodesDistance(depot, firstNode)

		// Travel between nodes in the sub-route
		for i := 0; i < len(sub.NodeIDs)-1; i++ {
			currentNode := instance.NodesMatrix[sub.NodeIDs[i]]
			nextNode := instance.NodesMatrix[sub.NodeIDs[i+1]]
			totalCost += instance.GetNodesDistance(currentNode, nextNode)
		}

		// Return from last node to depot
		lastNode := instance.NodesMatrix[sub.NodeIDs[len(sub.NodeIDs)-1]]
		totalCost += instance.GetNodesDistance(lastNode, depot)
	}

	return totalCost
}
