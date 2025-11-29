package graphs

type Graph interface {
	AddVertex() int
	RemoveVertex(id int)
	AddEdge() int
	RemoveEdge(id int)
}

type Vertex interface {
	AddEdge() 
	RemoveEdge()
}

type Edge interface {
	Target() Vertex
}
