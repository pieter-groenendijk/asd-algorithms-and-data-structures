package adlist

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"

func (g *AdjacencyList) newId() graphs.Id {
	g.lastUsedId++
	return g.lastUsedId
}
