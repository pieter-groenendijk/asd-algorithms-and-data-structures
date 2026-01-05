package adlist

import "github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs"

type AdjacencyList struct {
	edges map[graphs.Id][]graphs.Id // map[fromNodeId][]toNodeId
	lastUsedId graphs.Id
}

func New(capacity int) *AdjacencyList {
	return &AdjacencyList{
		edges: make(map[graphs.Id][]graphs.Id, capacity),
		lastUsedId: -1,
	}
}

