package adlist 

func (g *AdjacencyList) newId() Id {
	g.lastUsedId++
	return g.lastUsedId
}
