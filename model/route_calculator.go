package model

type RouteCalculator struct{}

// CalculateWithGreedy evaluates the cost of a route using the Greedy algorithm.
// Assumptions:
// - Node 0 (depot) is added at the beginning and end of the route.
// - The truck has limited capacity and must return to the depot when overloaded.
func (r RouteCalculator) CalculateWithGreedy(instance Instance, route Route) float32 {
	totalCost := float32(0)
	currentLoad := 0
	depot := instance.NodesMatrix[instance.GetDepotID()]
	currentNode := depot

	for _, nodeID := range route.NodeIDs {
		nextNode := instance.NodesMatrix[nodeID]

		// Check if adding the next node's load would exceed truck capacity
		if currentLoad+nextNode.Load > instance.TruckMaxLoad {
			// Return to depot first
			totalCost += instance.GetNodesDistance(currentNode, depot)
			currentNode = depot
			currentLoad = 0
		}

		// Go to next node
		totalCost += instance.GetNodesDistance(currentNode, nextNode)
		currentNode = nextNode
		currentLoad += nextNode.Load
	}

	// Return to depot at the end
	totalCost += instance.GetNodesDistance(currentNode, depot)
	return totalCost
}
