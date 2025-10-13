// Package model provides data structures and functions for handling problem instances.
package model

type Instance struct {
	Name           string
	NodesMatrix    []Node
	DistanceMatrix [][]float32
	TruckMaxLoad   int
}

func (i Instance) GetNodesDistance(n1 Node, n2 Node) float32 {
	return i.DistanceMatrix[n1.ID][n2.ID]
}

// GetDepotID returns the depot ID, which is always 0; other nodes are 1..n.
func (i Instance) GetDepotID() int {
	return 0
}
