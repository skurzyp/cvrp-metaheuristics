package model

type RouteCalculator struct{}

// CalculateWithGreedy evaluates the cost of a route using the Greedy algorithm.
// Assumptions:
// - Node 0 (depot) is added at the beginning and end of the route.
// - The truck has limited capacity and must return to the depot when overloaded.
func (r RouteCalculator) CalculateWithGreedy(instance Instance, route Route) float32 {
	totalCost := float32(0)
	currentLoad := 0
	currentNode := instance.NodesMatrix[instance.GetDepotID()]

	for _, nodeID := range route.NodeIDs {
		nextNode := instance.NodesMatrix[nodeID]
		distance := instance.GetNodesDistance(currentNode, nextNode)

		// Check if the truck is overloaded
		if currentLoad+1 > instance.TruckMaxLoad {
			// Return to the depot
			totalCost += instance.GetNodesDistance(currentNode, instance.NodesMatrix[instance.GetDepotID()])
			currentLoad = 0
		}

		// Visit the next node
		totalCost += distance
		currentNode = nextNode
		currentLoad++
	}

	// Return to the depot at the end
	totalCost += instance.GetNodesDistance(currentNode, instance.NodesMatrix[instance.GetDepotID()])
	return totalCost
}
