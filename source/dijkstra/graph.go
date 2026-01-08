package dijkstra

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs/adlist"
)

type Edge struct {
	Weight int
}

type Graph[TVertex any, TEdge Edge] struct {
	adlist.AdjacencyList[TVertex, TEdge]
}
