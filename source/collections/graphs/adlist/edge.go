package adlist

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"
)

type Edge[TVertex graphs.Vertex[any]] struct {
	target *TVertex
}

func NewEdge[TVertex graphs.Vertex[any]](target *TVertex) *Edge[TVertex] {
	return &Edge[TVertex]{
		target: target,
	}
}

func (e *Edge[TVertex]) Target() *TVertex {
	return e.target
}
