package adlist

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"
)

/*
type Graph interface {
	haveEdge(fromVertexId Id, toVertexId Id) bool
	addVertex() Id
	removeVertex(id Id)
	addEdge(fromVertexId Id, toVertexId Id) 
	removeEdge(fromVertexId Id, toVertexId Id)
}
*/

func (g *AdjacencyList) HaveEdge(fromVertexId graphs.Id, toVertexId graphs.Id) bool {
	edges := g.edges[fromVertexId]
	if edges == nil {
		return false
	}

	length := len(edges)
	for i := 0; i < length; i++ {
		if edges[i] == toVertexId {
			return true
		}
	}

	return false
}

func (g *AdjacencyList) GetVertices(id Id) map[Id][]Id {
	return g.edges
}

func (g *AdjacencyList) AddVertex(edgeCapacity int) graphs.Id {
	id := g.newId()

	g.edges[id] = make([]graphs.Id, edgeCapacity)

	return id
}

func (g *AdjacencyList) RemoveVertex(vertexId graphs.Id) {
	delete(g.edges, vertexId)
}

// May return nil if the vertex does not exist
func (g *AdjacencyList) GetEdges(vertexId graphs.Id) []graphs.Id {
	return g.edges[vertexId]
}

// addEdge duplicate edges are not prevented
func (g *AdjacencyList) AddEdge(fromVertexId graphs.Id, toVertexId graphs.Id) {
	g.edges[fromVertexId] = append(g.edges[fromVertexId], toVertexId)
}

func (g *AdjacencyList) RemoveEdge(fromVertexId graphs.Id, toVertexId graphs.Id) {
	edges := g.edges[fromVertexId]
	if edges == nil {
		return
	}

	length := len(edges)
	for i := 0; i < length; i++ {
		if edges[i] == toVertexId {
			edges[i] = edges[length - 1] // order does not matter, swap with last vertex id, instead of splicing. `O(1) instead of O(n)`
			g.edges[fromVertexId] = edges[:length - 1] // then we change our view, doing as it never existed
			// TODO: Update slice for bigger discrepencies to avoid bigger memory leaks
		}
	}
}
