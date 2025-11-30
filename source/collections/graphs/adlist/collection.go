package adlist

import (
	"iter"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"
)

func (g *AdjacencyList[TVertex, TEdge]) AllVertices() iter.Seq2[graphs.Id, TVertex] {
	return func(yield func(graphs.Id, TVertex) bool) {
		vertices := g.vertices
		for key, value := range vertices {
			if !yield(key, value) {
				return
			}
		}
	}
}
