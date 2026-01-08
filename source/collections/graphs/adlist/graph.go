package adlist

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/swapback"
)

func (g *AdjacencyList[TVertex, TEdge]) NumOfVertices() int {
	return len(g.sourceToTargets)
}

func (g *AdjacencyList[TVertex, TEdge]) AddVertex(vertex TVertex, edgeCap int) int {
	holesAt, holeAt, holeFound := swapback.Pop(g.holesAt)

	g.holesAt = holesAt

	if holeFound {
		g.sourceToTargets[holeAt] = make([]int, 0, edgeCap)
		g.edges[holeAt] = make([]TEdge, 0, edgeCap)
		g.vertices[holeAt] = vertex

		return holeAt
	} else {
		g.sourceToTargets = append(g.sourceToTargets, make([]int, 0, edgeCap))
		g.edges = append(g.edges, make([]TEdge, 0, edgeCap))
		g.vertices = append(g.vertices, vertex)

		return len(g.sourceToTargets) - 1
	}
}

func (g *AdjacencyList[TVertex, TEdge]) RemoveVertex(vertex int) {
	// Disconnect outgoing
	targets := g.sourceToTargets[vertex]
	for _, target := range targets {
		sources := g.targetToSources[target]
		for at, source := range sources {
			if source == vertex {
				g.targetToSources[target] = swapback.Remove(sources, at)
				break
			}
		}
	}

	// Disconnect incoming
	sources := g.targetToSources[vertex]
	for _, source := range sources {
		targets := g.sourceToTargets[source]
		for at, target := range targets {
			if target == vertex {
				g.sourceToTargets[source] = swapback.Remove(targets, at)
				g.edges[source] = swapback.Remove(g.edges[source], at)
				break
			}
		}
	}

	// Wipe data
	g.sourceToTargets[vertex] = nil
	g.edges[vertex] = nil
	g.targetToSources[vertex] = nil
	var zeroVertex TVertex // this only helps if the user gave a pointer type, or a struct with inner pointer types.
	g.vertices[vertex] = zeroVertex

	// Make our delete tracable
	g.holesAt = append(g.holesAt, vertex)
}

// May return nil slice if the vertex does not exist
func (g *AdjacencyList[TVertex, TEdge]) GetTargetsOf(sourceVertex int) []int {
	return g.sourceToTargets[sourceVertex]
}

func (g *AdjacencyList[TVertex, TEdge]) GetSourcesOf(targetVertex int) []int {
	return g.targetToSources[targetVertex]
}

func (g *AdjacencyList[TVertex, TEdge]) AddEdge(fromVertex int, toVertex int, edge TEdge) {
	g.sourceToTargets[fromVertex] = append(g.sourceToTargets[fromVertex], toVertex)
	g.edges[fromVertex] = append(g.edges[fromVertex], edge)

	g.targetToSources[toVertex] = append(g.targetToSources[toVertex], fromVertex)
}

// O(SE + TE), where SE is the amount of edges the source has, and TE is the amount of edges the target has
func (g *AdjacencyList[TVertex, TEdge]) RemoveEdge(sourceVertex int, targetVertex int) {
	// Disconnect incoming
	sourcesOfTarget := g.targetToSources[targetVertex]
	for at, sourceOfTarget := range sourcesOfTarget {
		if sourceOfTarget == sourceVertex {
			g.targetToSources[targetVertex] = swapback.Remove(sourcesOfTarget, at)
		}
	}

	// Disconnect outgoing
	targetsOfSource := g.sourceToTargets[sourceVertex]
	for at, targetOfSource := range targetsOfSource {
		if targetOfSource == targetVertex {
			g.sourceToTargets[sourceVertex] = swapback.Remove(targetsOfSource, at)
			g.edges[sourceVertex] = swapback.Remove(g.edges[sourceVertex], at)
		}
	}
}
