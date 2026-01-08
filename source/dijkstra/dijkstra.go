package dijkstra

import (
	"math"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/priorqueue"
)

// Any vertex which has been checked has a pathToSrc
type pathToSrc struct {
	src  int
	dist int
}

func (g *Graph[TVertex, TEdge]) ShortestPathsTo(startVertex int) []pathToSrc {
	numOfVertices := g.NumOfVertices()

	pathsToSrc := make([]pathToSrc, numOfVertices)
	vertsToVisit := priorqueue.New[int, int](numOfVertices * 3 / 2) // (-knownShortestDistToSrc) -> vertex
	visitedVerts := make([]bool, numOfVertices)                     // TODO: Optimize as bitset

	for vertex := 0; vertex < numOfVertices; vertex++ {
		pathsToSrc[vertex] = pathToSrc{
			src:  -1,
			dist: math.MaxInt,
		}
		vertsToVisit.Push(0, vertex)
	}

	pathsToSrc[startVertex].dist = 0
	vertsToVisit.Push(math.MaxInt, startVertex) // shorthand of math.MaxInt - 0
	for {
		sourceVertex, exists := vertsToVisit.Pop()
		if !exists {
			break
		}
		if visitedVerts[sourceVertex] {
			continue
		}
		distToSrc := pathsToSrc[sourceVertex].dist

		targetVertices := g.GetTargetsOf(sourceVertex)
		for _, targetVertex := range targetVertices {
			altDist := distToSrc + Edge(g.GetEdgeValue(sourceVertex, targetVertex)).Weight
			if altDist < pathsToSrc[targetVertex].dist {
				pathsToSrc[targetVertex].src = sourceVertex
				pathsToSrc[targetVertex].dist = altDist
				vertsToVisit.Push(math.MaxInt-altDist, targetVertex) // TODO: update priority (make indexed binheap to achieve performant?), instead of pushing again? Works better for dense graphs
			}
		}
	}

	return pathsToSrc
}
