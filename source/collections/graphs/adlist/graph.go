package adlist

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/swapback"
)

func (g *AdjacencyList) NumOfVertices() int {
	return len(g.sourceToTargets)
}

func (g *AdjacencyList) AddVertex(edgeCap int) int {
	holesAt, holeAt, holeFound := swapback.Pop(g.holesAt)
	g.holesAt = holesAt
	if holeFound {
		g.sourceToTargets[holeAt] = make([]int, edgeCap)

		return holeAt
	} else {
		g.sourceToTargets = append(g.sourceToTargets, make([]int, edgeCap))

		return len(g.sourceToTargets) - 1
	}
}

func (g *AdjacencyList) RemoveVertex(vertex int) {
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

	sources := g.targetToSources[vertex]
	for _, source := range sources {
		targets := g.sourceToTargets[source]
		for at, target := range targets {
			if target == vertex {
				g.sourceToTargets[source] = swapback.Remove(targets, at)
				break
			}
		}
	}

	g.sourceToTargets[vertex] = nil
	g.targetToSources[vertex] = nil

	g.holesAt = append(g.holesAt, vertex)
}

// May return nil slice if the vertex does not exist
func (g *AdjacencyList) GetTargetsOf(sourceVertex int) []int {
	return g.sourceToTargets[sourceVertex]
}

func (g *AdjacencyList) GetSourcesOf(targetVertex int) []int {
	return g.targetToSources[targetVertex]
}

func (g *AdjacencyList) AddEdge(fromVertex int, toVertex int) {
	g.sourceToTargets[fromVertex] = append(g.sourceToTargets[fromVertex], toVertex)
}

// O(SE + TE), where SE is the amount of edges the source has, and TE is the amount of edges the target has
func (g *AdjacencyList) RemoveEdge(sourceVertex int, targetVertex int) {
	sourcesOfTarget := g.targetToSources[targetVertex]
	for at, sourceOfTarget := range sourcesOfTarget {
		if sourceOfTarget == sourceVertex {
			g.targetToSources[targetVertex] = swapback.Remove(sourcesOfTarget, at)
		}
	}

	targetsOfSource := g.sourceToTargets[sourceVertex]
	for at, targetOfSource := range targetsOfSource {
		if targetOfSource == targetVertex {
			g.sourceToTargets[sourceVertex] = swapback.Remove(targetsOfSource, at)
		}
	}
}
