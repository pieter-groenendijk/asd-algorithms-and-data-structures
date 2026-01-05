package adlist

import (
)

type AdjacencyList struct {
	holesAt []int // contains where holes exist in `vertexToEdges`
	toVertices [][]int // [fromVertexId] -> []toVertexId, answers: What do I point to?
	fromVertices [][]int // [toVertexId] -> []fromVertexId, answers: What edges point to me?
}

// initVertCap: the initial capacity of vertices
// 
// initEdgeCap: the initial capacity of edges per vertices
func New(initVertCap int, initEdgeCap int) *AdjacencyList {
	return &AdjacencyList{
		holesAt: make([]int, 8),
		toVertices: make([][]int, initVertCap),
		fromVertices: make([][]int, initVertCap * initEdgeCap),
	}
}
