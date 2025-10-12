package model

import (
	"fmt"
	"strings"
)

type FinalSolution struct {
	Route     Route
	Cost      float32
	SubRoutes []Route // sub-routes based on truck load capacity
}

func (fs FinalSolution) String() string {
	var sb strings.Builder

	// Print each sub-route
	for i, sub := range fs.SubRoutes {
		sb.WriteString(fmt.Sprintf("Route #%d: ", i+1))
		for j, nodeID := range sub.NodeIDs {
			// Skip the depot if included at start or end
			if nodeID == 0 {
				continue
			}
			if j > 0 {
				sb.WriteString(" ")
			}
			sb.WriteString(fmt.Sprint(nodeID))
		}
		sb.WriteString("\n")
	}

	// Print total cost
	sb.WriteString(fmt.Sprintf("Cost: %.2f\n", fs.Cost))

	return sb.String()
}
