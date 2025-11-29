package adlist 

type Vertex struct {
	edges map[Id]*Edge
}

func (v *Vertex) AddEdge(id Id, edge *Edge) {
	v.edges[id] = edge
}

func (v *Vertex) RemoveEdge(id Id) {
	delete(v.edges, id)
}
