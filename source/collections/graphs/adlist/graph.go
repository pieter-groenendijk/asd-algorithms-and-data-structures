package adlist

import (
	"fmt"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"
)

func (g *AdjacencyList[TVertex, TEdge]) GetVertex(id graphs.Id) (TVertex, error) {
	vertex, hasVertex := g.vertices[id]
	if !hasVertex {
		return vertex, fmt.Errorf("failed to get vertex: %w", collections.ErrNotFound)
	}

	return vertex, nil
}

func (g *AdjacencyList[TVertex, TEdge]) AddVertex(vertex TVertex) graphs.Id {
	id := g.newId()

	g.vertices[id] = vertex
	
	return id
}

func (g *AdjacencyList[TVertex, TEdge]) RemoveVertex(id graphs.Id) {
	delete(g.vertices, id)
}

func (g *AdjacencyList[TVertex, TEdge]) GetEdge(fromVertex graphs.Id, edgeId graphs.Id) (TEdge, error) {
	vertex, err := g.GetVertex(fromVertex)
	if err != nil {
		var edge TEdge
		return edge, fmt.Errorf("failed to get edge: %w", err)
	}

	edge, err := vertex.GetEdge(edgeId)
	if err != nil {
		return edge, err
	}

	return edge, nil
}

func (g *AdjacencyList[TVertex, TEdge]) AddEdge(fromVertexId graphs.Id, edge TEdge) (graphs.Id, error) {
	vertex, err := g.GetVertex(fromVertexId)
	if err != nil {
		return -1, fmt.Errorf("failed to add edge: %w", err)
	}

	id := g.newId()
	vertex.AddEdge(id, edge)

	return id, nil
}

func (g *AdjacencyList[TVertex, TEdge]) RemoveEdge(fromVertexId, id graphs.Id) error {
	vertex, err := g.GetVertex(fromVertexId)
	if err != nil {
		return fmt.Errorf("failed to remove edge: %w", err)
	}

	vertex.RemoveEdge(fromVertexId)

	return nil
}

