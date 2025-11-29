package adlist

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"

func (g *AdjacencyList) AddVertex(vertex *Vertex) Id {
	id := g.newId()

	g.vertices[id] = vertex
	
	return id
}

func (g *AdjacencyList) RemoveVertex(id Id) {
	delete(g.vertices, id)
}

func (g *AdjacencyList) AddEdge(fromVertexId Id, edge *Edge) (Id, error) {
	vertex, hasVertex := g.vertices[fromVertexId]
	if !hasVertex {
		return -1, collections.ErrNotFound
	}

	id := g.newId()
	vertex.AddEdge(id, edge)

	return id, nil
}

func (g *AdjacencyList) RemoveEdge(fromVertexId Id, id Id) {
	vertex, hasVertex := g.vertices[fromVertexId]
	if !hasVertex {
		return
	}

	vertex.RemoveEdge(fromVertexId)
}

