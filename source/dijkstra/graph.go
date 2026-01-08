package dijkstra

import (
	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/collections/graphs/adlist"
)

type WeightedEdge struct {
	Weight int
}

type Graph[TVertex any, TEdge WeightedEdge] struct {
	adlist.AdjacencyList[TVertex, TEdge]
}

func New[TVertex any, TEdge WeightedEdge](initVertCap int) *Graph[TVertex, TEdge] {
	return &Graph[TVertex, TEdge]{
		AdjacencyList: *adlist.New[TVertex, TEdge](initVertCap),
	}
}
