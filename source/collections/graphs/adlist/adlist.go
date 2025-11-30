package adlist

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"

type AdjacencyList[TVertex graphs.Vertex[TEdge], TEdge graphs.Edge[TVertex]] struct {
	vertices map[graphs.Id]TVertex

	lastUsedId graphs.Id
}

func New[TVertex graphs.Vertex[TEdge], TEdge graphs.Edge[TVertex]]() *AdjacencyList[TVertex, TEdge] {
	return &AdjacencyList[TVertex, TEdge]{
		vertices: make(map[graphs.Id]TVertex),
		lastUsedId: -1,
	}
}

