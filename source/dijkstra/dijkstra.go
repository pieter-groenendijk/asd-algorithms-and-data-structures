package dijkstra

import (
	"math"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs/adlist"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/priorqueue"
)

// Any vertex which has been checked has a pathToSrc
type pathToSrc struct {
	src int 
	dist int
}

type EdgeKey struct {
	sourceVertex int
	targetVertex int
}

func ShortestPathsTo(g *adlist.AdjacencyList, sourceVertex int, edgeWeights map[EdgeKey]int) {
	numOfVertices := g.NumOfVertices()

	pathsToSrc := make([]pathToSrc, numOfVertices)
	vertsToVisit := priorqueue.New[int, int](16) // (-knownShortestDistToSrc) -> vertex

	for vertex := 0; vertex < numOfVertices; vertex++ {
		pathsToSrc[vertex] = pathToSrc{
			src: -1,
			dist: 0,
		}
		vertsToVisit.Push(0, vertex)
	}

	pathsToSrc[sourceVertex] = 
	vertsToVisit.Push(math.MaxInt, sourceVertex)	
	for {
		curr, shouldVisit := vertsToVisit.Pop()
		if !shouldVisit {
			break
		}
		distToSrc := pathsToSrc[curr].dist

		targets := g.GetTargetsOf(curr)
		for _, target := range targets {
			altDist := distToSrc + edgeWeights[EdgeKey{curr, target}]
			if altDist < pathsToSrc[target].dist {
				pathsToSrc[target].src = curr
				pathsToSrc[target].dist = altDist
				// TODO: update priority, should still work without it though
			}
		}
	}
}
