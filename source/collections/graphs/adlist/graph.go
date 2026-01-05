package adlist

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/bag"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/swapback"
)

func (g *AdjacencyList) HaveVertex(vertexId int) bool {
	return vertexId < len(g.toVertices)
}

func (g *AdjacencyList) HaveEdge(fromVertexId int, toVertexId int) bool {
	edgeExists := toVertexId < len(g.fromVertices)
	if !edgeExists {
		return false
	}

	edges := g.fromVertices[toVertexId]
	edgesLen := len(edges)
	for i := 0; i < edgesLen; i++ {
		if edges[i] == fromVertexId {
			return true
		}
	}

	return false
}

func (g *AdjacencyList) AddVertex(edgeCapacity int) int {
	holesAt, holeAt, holeFound := swapback.Pop(g.holesAt)
	g.holesAt = holesAt
	if holeFound {
		g.toVertices[holeAt] = make([]int, edgeCapacity)

		return holeAt
	} else {
		g.toVertices = append(g.toVertices, make([]int, edgeCapacity))

		return len(g.toVertices) - 1
	}
}

func (g *AdjacencyList) RemoveVertex(vertexId int) {
	fromVertices := g.fromVertices[vertexId] // 
	


}

// May return nil slice if the vertex does not exist
func (g *AdjacencyList) GetEdges(vertexId int) []int {
	return g.toVertices[vertexId]
}

func (g *AdjacencyList) AddEdge(fromVertexId int, toVertexId int) {
	g.toVertices[fromVertexId] = append(g.toVertices[fromVertexId], toVertexId)
}

func (g *AdjacencyList) RemoveEdge(fromVertexId int, toVertexId int) {
	edges := g.toVertices[fromVertexId]
	if edges == nil {
		return
	}

	length := len(edges)
	for i := 0; i < length; i++ {
		if edges[i] == toVertexId {
			g.toVertices[fromVertexId] = swapback.Remove(g.toVertices[fromVertexId], i)
			break
		}
	}
}
