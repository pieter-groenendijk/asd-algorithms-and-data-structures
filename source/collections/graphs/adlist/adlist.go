package adlist

type Id int

type AdjacencyList struct {
	vertices map[Id]*Vertex

	lastUsedId Id
}


