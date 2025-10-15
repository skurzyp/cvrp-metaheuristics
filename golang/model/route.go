package model

import (
	"fmt"
	"strings"
)

type Route struct {
	NodeIDs []int
}

func (r Route) String() string {
	if len(r.NodeIDs) == 0 {
		return "empty route"
	}
	strIDs := make([]string, len(r.NodeIDs))
	for i, id := range r.NodeIDs {
		strIDs[i] = fmt.Sprint(id)
	}
	return strings.Join(strIDs, " -> ")
}
