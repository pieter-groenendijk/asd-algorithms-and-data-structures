package adlist

import (
)

type AdjacencyList struct {
	holesAt []int // contains where holes exist in `vertexToEdges`
	sourceToTargets [][]int // [fromVertexId] -> []toVertexId, answers: What do I point to?
	targetToSources [][]int // [toVertexId] -> []fromVertexId, answers: What edges point to me?
}

// initVertCap: the initial capacity of vertices
// 
// initEdgeCap: the initial capacity of edges per vertices
func New(initVertCap int) *AdjacencyList {
	return &AdjacencyList{
		holesAt: make([]int, 8),
		sourceToTargets: make([][]int, initVertCap),
		targetToSources: make([][]int, initVertCap),
	}
}
