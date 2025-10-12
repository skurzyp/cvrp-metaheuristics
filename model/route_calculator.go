package model

type RouteCalculator struct{}

func (r RouteCalculator) SplitIntoSubRoutes(instance Instance, route Route) []Route {
	var subRoutes []Route
	var currentSubRoute []int
	currentLoad := 0
	depot := instance.NodesMatrix[instance.GetDepotID()]

	for _, nodeID := range route.NodeIDs {
		node := instance.NodesMatrix[nodeID]

		if currentLoad+node.Load > instance.TruckMaxLoad {
			subRoutes = append(subRoutes, Route{NodeIDs: append([]int{depot.ID}, currentSubRoute...)})
			currentLoad = 0
			currentSubRoute = []int{}
		}

		currentSubRoute = append(currentSubRoute, node.ID)
		currentLoad += node.Load
	}

	if len(currentSubRoute) > 0 {
		subRoutes = append(subRoutes, Route{NodeIDs: append([]int{depot.ID}, currentSubRoute...)})
	}

	return subRoutes
}

// CalculateWithGreedy evaluates the cost of a route using the Greedy algorithm.
// Assumptions:
// - Node 0 (depot) is added at the beginning and end of the route.
// - The truck has limited capacity and must return to the depot when overloaded.
func (r RouteCalculator) CalculateWithGreedy(instance Instance, route Route) float32 {
	totalCost := float32(0)
	subRoutes := r.SplitIntoSubRoutes(instance, route)

	for _, sub := range subRoutes {
		currentNode := instance.NodesMatrix[instance.GetDepotID()]
		for _, nodeID := range sub.NodeIDs[1:] { // skip depot at start
			nextNode := instance.NodesMatrix[nodeID]
			totalCost += instance.GetNodesDistance(currentNode, nextNode)
			currentNode = nextNode
		}
		// Ensure return to depot
		totalCost += instance.GetNodesDistance(currentNode, instance.NodesMatrix[instance.GetDepotID()])
	}

	return totalCost
}
