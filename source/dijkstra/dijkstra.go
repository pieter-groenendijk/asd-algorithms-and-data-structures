package dijkstra

import (
	"fmt"
	"math"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/priorqueue"
)

const infinity = math.MaxInt

// returns the shortest path for directed acyclic graphs.
func ShortestPath[TVertex DistVertex[TEdge], TEdge WeightedEdge[TVertex]](graph DijkstraGraph[TVertex, TEdge], fromNodeId graphs.Id) error {
	fromNode, err  := graph.GetVertex(fromNodeId)
	if err != nil {
		return fmt.Errorf("failed to find fromNode: %w", err)
	}

	// the set of vertices, we know the shortest paths to
	processedVertices := make(map[*TVertex]interface{})
	toProcessVertices := make(priorqueue.PriorityQueue, 10)

	for _, vertex := range graph.AllVertices() {
		vertex.SetPrevious(nil)
		vertex.SetDistance(infinity)
		toProcessVertices.Push()
	}
	// we're too greedy, let's correct it
	fromNode.SetPrevious(nil)
	fromNode.SetDistance(0)



	return nil
}
