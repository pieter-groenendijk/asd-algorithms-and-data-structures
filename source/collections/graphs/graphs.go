package graphs

type Graph interface {
	HaveVertex(vertex int) bool
	HaveEdge(fromVertex int, toVertex int) bool
	AddVertex(edgeCap int) int 
	RemoveVertex(vertex int)
	GetEdgesFrom(fromVertex int)
	GetEdgesTo(toVertex int)
	AddEdge(fromVertex int, toVertex int) 
	RemoveEdge(fromVertex int, toVertex int)
}
