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

	// Print the full route
	sb.WriteString("Full Route: ")
	for i, nodeID := range fs.Route.NodeIDs {
		if i > 0 {
			sb.WriteString(" -> ")
		}
		sb.WriteString(fmt.Sprint(nodeID))
	}
	sb.WriteString("\n\n")

	// Print each sub-route
	sb.WriteString("Sub-routes:\n")
	for i, sub := range fs.SubRoutes {
		sb.WriteString(fmt.Sprintf("Route #%d: ", i+1))
		for j, nodeID := range sub.NodeIDs {
			if j > 0 {
				sb.WriteString(" -> ")
			}
			sb.WriteString(fmt.Sprint(nodeID))
		}
		sb.WriteString("\n")
	}

	// Print total cost
	sb.WriteString(fmt.Sprintf("\nTotal Cost: %.2f\n", fs.Cost))

	return sb.String()
}
