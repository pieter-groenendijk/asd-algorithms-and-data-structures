package adlist

type AdjacencyList[TVertex any, TEdge any] struct {
	holesAt         []int   // contains where holes exist in `vertexToEdges`
	sourceToTargets [][]int // [fromVertexId] -> []toVertexId, answers: What do I point to?
	targetToSources [][]int // [toVertexId] -> []fromVertexId, answers: What edges point to me?

	// Additional data
	vertices []TVertex // may be nil, [vertexId] -> TVertex
	edges    [][]TEdge // may be nil, mirrors sourceToTargets, [fromVertexId] -> []toVertexId -> TEdge

	numOfVertices int
}

// initVertCap: the initial capacity of vertices
//
// initEdgeCap: the initial capacity of edges per vertices
func New[TVertex any, TEdge any](initVertCap int) *AdjacencyList[TVertex, TEdge] {
	return &AdjacencyList[TVertex, TEdge]{
		holesAt:         make([]int, 0, 8),
		sourceToTargets: make([][]int, 0, initVertCap),
		targetToSources: make([][]int, 0, initVertCap),
		vertices:        make([]TVertex, 0, initVertCap),
		edges:           make([][]TEdge, 0, initVertCap),
		numOfVertices:   0,
	}
}
