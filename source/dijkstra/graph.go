package dijkstra

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"
)

// A graph which dijkstra's shortest path algorithm can run upon
type DijkstraGraph[TVertex DistVertex[TEdge], TEdge WeightedEdge[TVertex]] interface {
	graphs.Graph[TVertex, TEdge]
}

// generic needs to be any to prevent recursive type definition
// types enforced in graph interface
type DistVertex[TEdge any] interface {
	graphs.Vertex[TEdge]

	Previous() DistVertex[TEdge]
	SetPrevious(vertex DistVertex[TEdge]) 
	
	Distance() int // Distance from this node towards the source (where it came from)
	SetDistance(int) 
}

// generic needs to be any to prevent recursive type definition
// types enforced in graph interface
type WeightedEdge[TVertex any] interface {
	graphs.Edge[TVertex]

	Weight() int
	SetWeight(int)
}
