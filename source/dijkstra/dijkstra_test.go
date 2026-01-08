package dijkstra

import (
	"math"
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

type Vertex struct{}
type Edge = WeightedEdge

func new() *Graph[Vertex, Edge] {
	return New[Vertex, Edge](16)
}

func TestShortestPath(t *testing.T) {
	t.Run("emptyGraph", func(t *testing.T) {
		g := new()

		defer func() {
			recover()
		}()

		g.ShortestPathsTo(0, 1)
	})

	t.Run("startIsEnd", func(t *testing.T) {
		g := new()

		id := g.AddVertex(Vertex{}, 3)

		gottenPaths := g.ShortestPathsTo(id, id)

		testutils.AssertEquals(
			t,
			Path{
				src:  -1,
				dist: 0,
			},
			gottenPaths[id],
		)
	})

	t.Run("noPath", func(t *testing.T) {
		g := new()

		idOne := g.AddVertex(Vertex{}, 3)
		idTwo := g.AddVertex(Vertex{}, 3)

		gottenPaths := g.ShortestPathsTo(idOne, idTwo)

		testutils.AssertEquals(
			t,
			gottenPaths[idTwo],
			Path{
				src:  -1,
				dist: math.MaxInt,
			},
		)
	})

	t.Run("directShortestPath", func(t *testing.T) {
		g := new()

		idOne := g.AddVertex(Vertex{}, 3)
		idTwo := g.AddVertex(Vertex{}, 3)
		idThree := g.AddVertex(Vertex{}, 3)

		g.AddEdge(idOne, idTwo, Edge{Weight: 5})
		g.AddEdge(idTwo, idThree, Edge{Weight: 5})
		g.AddEdge(idThree, idOne, Edge{Weight: 2}) // Just to possibly throw the algo off
		g.AddEdge(idTwo, idOne, Edge{Weight: 5})   // Just to possibly throw the algo off

		gottenPaths := g.ShortestPathsTo(idOne, idThree)

		testutils.AssertEquals(t, Path{
			src:  idTwo,
			dist: 10,
		}, gottenPaths[idThree])
		testutils.AssertEquals(t, Path{
			src:  idOne,
			dist: 5,
		}, gottenPaths[idTwo])
		testutils.AssertEquals(t, Path{
			src:  -1,
			dist: 0,
		}, gottenPaths[idOne])
	})

	t.Run("indirectShortestPath", func(t *testing.T) {
		g := new()

		idOne := g.AddVertex(Vertex{}, 3)
		idTwo := g.AddVertex(Vertex{}, 3)
		idThree := g.AddVertex(Vertex{}, 3)

		// Direct path
		g.AddEdge(idOne, idThree, Edge{Weight: 15})

		// Indirect path
		g.AddEdge(idOne, idTwo, Edge{Weight: 5})
		g.AddEdge(idTwo, idThree, Edge{Weight: 5})

		// Just to possibly throw off the algo

		gottenPaths := g.ShortestPathsTo(idOne, idThree)

		testutils.AssertEquals(t, Path{
			src:  idTwo,
			dist: 10,
		}, gottenPaths[idThree])
		testutils.AssertEquals(t, Path{
			src:  idOne,
			dist: 5,
		}, gottenPaths[idTwo])
		testutils.AssertEquals(t, Path{
			src:  -1,
			dist: 0,
		}, gottenPaths[idOne])
	})
}
