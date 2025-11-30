package adlist

import (
	"fmt"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"
)

type Vertex[TEdge graphs.Edge[any]] struct {
	edges map[graphs.Id]TEdge
}

func NewVertex[TEdge graphs.Edge[any]]() *Vertex[TEdge] {
	return &Vertex[TEdge]{
		edges: make(map[graphs.Id]TEdge),
	}
}

func (v Vertex[TEdge]) GetEdge(id graphs.Id) (TEdge, error) {
	edge, hasEdge := v.edges[id]
	if !hasEdge {
		return edge, fmt.Errorf("failed to get edge: %w", collections.ErrNotFound)
	}

	return edge, nil
}

func (v Vertex[TEdge]) AddEdge(id graphs.Id, edge TEdge) {
	v.edges[id] = edge
}

func (v Vertex[TEdge]) RemoveEdge(id graphs.Id) {
	delete(v.edges, id)
}
