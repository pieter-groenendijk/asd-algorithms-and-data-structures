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
	vertsToVisit := priorqueue.New[int, int](16) // (infinity - knownShortestDistToSrc) -> vertex

	for i := 0; i < numOfVertices; i++ {
		pathsToSrc[i] = pathToSrc{
			src: -1,
			dist: math.MaxInt,
		}
	}

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
				// update priority
			}
		}
	}
}
