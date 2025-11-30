package adlist

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"

func (g *AdjacencyList[TVertex, TEdge]) newId() graphs.Id {
	g.lastUsedId++
	return g.lastUsedId
}
