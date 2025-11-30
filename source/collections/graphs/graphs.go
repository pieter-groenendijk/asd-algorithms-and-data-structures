package graphs

import "iter"

type Id int

type Graph[TVertex Vertex[TEdge], TEdge Edge[TVertex]] interface {
	GetVertex(id Id) (TVertex, error)
	AddVertex(vertex TVertex) Id
	RemoveVertex(id Id)
	GetEdge(fromVertex Id, edgeId Id) (TEdge, error)
	AddEdge(fromVertex Id, edge TEdge) (Id, error)
	RemoveEdge(fromVertex Id, id Id) (Id, error)
	AllVertices() iter.Seq2[Id, TVertex]
}

// generic needs to be any to prevent recursive type definition
// types enforced in graph interface
type Vertex[TEdge any] interface {
	GetEdge(id Id) (TEdge, error)
	AddEdge(edgeId Id, edge TEdge) 
	RemoveEdge(edgeId Id)
}

// generic needs to be any to prevent recursive type definition
// types enforced in graph interface
type Edge[TVertex any] interface {
	Target() TVertex
}
