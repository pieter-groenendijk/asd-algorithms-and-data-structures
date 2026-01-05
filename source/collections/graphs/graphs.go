package graphs

type Id int

type Graph interface {
	HaveEdge(fromVertexId Id, toVertexId Id) bool
	GetVertices() map[Id][]Id
	AddVertex() Id
	RemoveVertex(id Id)
	GetEdgesOf(vertexId Id) []Id
	AddEdge(fromVertexId Id, toVertexId Id) 
	RemoveEdge(fromVertexId Id, toVertexId Id)
}
