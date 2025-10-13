package algorithms

import (
	"math/rand"

	"pwr.lab0/model"
)

// RandomRoute generates a random solution (for random baseline).
func RandomRoute(size int) model.Route {
	nodeIDs := make([]int, size-1)
	for i := 1; i < size; i++ {
		nodeIDs[i-1] = i
	}
	rand.Shuffle(len(nodeIDs), func(i, j int) {
		nodeIDs[i], nodeIDs[j] = nodeIDs[j], nodeIDs[i]
	})
	return model.Route{NodeIDs: nodeIDs}
}
