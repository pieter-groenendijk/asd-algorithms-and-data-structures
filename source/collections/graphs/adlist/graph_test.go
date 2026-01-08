package adlist

import (
	"testing"

	"github.com/pieter-groenendijk/asd-algorithms-and-data-structures/testutils"
)

type Vertex struct {
	name string
}

type Edge struct {
	name string
}

// A blackbox usage test
func Test(t *testing.T) {
	g := New[Vertex, Edge](8)

	// Testing empty state
	testutils.AssertEquals(t, 0, g.NumOfVertices())

	// Testing holding of vertex state
	idOne := g.AddVertex(Vertex{name: "one"}, 3)
	testutils.AssertEquals(t, 1, g.NumOfVertices())
	testutils.AssertEquals(t, []int{}, g.GetSourcesOf(idOne))
	testutils.AssertEquals(t, []int{}, g.GetTargetsOf(idOne))
	testutils.AssertEquals(t, Vertex{name: "one"}, g.GetVertexValue(idOne))

	g.RemoveVertex(idOne)
	testutils.AssertEquals(t, 0, g.NumOfVertices())

	idTwo := g.AddVertex(Vertex{name: "two"}, 6)
	testutils.AssertEquals(t, 1, g.NumOfVertices())
	testutils.AssertEquals(t, []int{}, g.GetSourcesOf(idTwo))
	testutils.AssertEquals(t, []int{}, g.GetTargetsOf(idTwo))
	testutils.AssertEquals(t, Vertex{name: "two"}, g.GetVertexValue(idTwo))

	idThree := g.AddVertex(Vertex{name: "three"}, 6)
	testutils.AssertEquals(t, 2, g.NumOfVertices())
	testutils.AssertEquals(t, []int{}, g.GetSourcesOf(idTwo))
	testutils.AssertEquals(t, []int{}, g.GetTargetsOf(idTwo))
	testutils.AssertEquals(t, Vertex{name: "two"}, g.GetVertexValue(idTwo))
	testutils.AssertEquals(t, []int{}, g.GetSourcesOf(idThree))
	testutils.AssertEquals(t, []int{}, g.GetTargetsOf(idThree))
	testutils.AssertEquals(t, Vertex{name: "three"}, g.GetVertexValue(idThree))

	// Testing holding of edge state
	g.AddEdge(idTwo, idThree, Edge{name: "edgeOne"})
	testutils.AssertEquals(t, 2, g.NumOfVertices())
	testutils.AssertEquals(t, []int{}, g.GetSourcesOf(idTwo))
	testutils.AssertEquals(t, []int{idThree}, g.GetTargetsOf(idTwo))
	testutils.AssertEquals(t, Vertex{name: "two"}, g.GetVertexValue(idTwo))
	testutils.AssertEquals(t, []int{idTwo}, g.GetSourcesOf(idThree))
	testutils.AssertEquals(t, []int{}, g.GetTargetsOf(idThree))
	testutils.AssertEquals(t, Vertex{name: "three"}, g.GetVertexValue(idThree))
	testutils.AssertEquals(t, Edge{name: "edgeOne"}, g.GetEdgeValue(idTwo, idThree))

	g.RemoveEdge(idTwo, idThree)
	testutils.AssertEquals(t, 2, g.NumOfVertices())
	testutils.AssertEquals(t, []int{}, g.GetSourcesOf(idTwo))
	testutils.AssertEquals(t, []int{}, g.GetTargetsOf(idTwo))
	testutils.AssertEquals(t, Vertex{name: "two"}, g.GetVertexValue(idTwo))
	testutils.AssertEquals(t, []int{}, g.GetSourcesOf(idThree))
	testutils.AssertEquals(t, []int{}, g.GetTargetsOf(idThree))
	testutils.AssertEquals(t, Vertex{name: "three"}, g.GetVertexValue(idThree))

	g.AddEdge(idTwo, idThree, Edge{name: "edgeTwo"})
	testutils.AssertEquals(t, 2, g.NumOfVertices())
	testutils.AssertEquals(t, []int{}, g.GetSourcesOf(idTwo))
	testutils.AssertEquals(t, []int{idThree}, g.GetTargetsOf(idTwo))
	testutils.AssertEquals(t, Vertex{name: "two"}, g.GetVertexValue(idTwo))
	testutils.AssertEquals(t, []int{idTwo}, g.GetSourcesOf(idThree))
	testutils.AssertEquals(t, []int{}, g.GetTargetsOf(idThree))
	testutils.AssertEquals(t, Vertex{name: "three"}, g.GetVertexValue(idThree))
	testutils.AssertEquals(t, Edge{name: "edgeTwo"}, g.GetEdgeValue(idTwo, idThree))

	// Accumulation
	idFour := g.AddVertex(Vertex{name: "four"}, 3)
	idFive := g.AddVertex(Vertex{name: "five"}, 5)
	g.AddEdge(idThree, idFive, Edge{name: "edgeThree"})
	g.AddEdge(idFour, idFive, Edge{name: "edgeFour"})
	g.AddEdge(idFive, idThree, Edge{name: "edgeFive"})
	g.AddEdge(idFour, idTwo, Edge{name: "edgeSix"})

	testutils.AssertEquals(t, 4, g.NumOfVertices())

	testutils.AssertEquals(t, []int{idThree}, g.GetTargetsOf(idTwo))
	testutils.AssertEquals(t, []int{idFour}, g.GetSourcesOf(idTwo))

	testutils.AssertEquals(t, []int{idFive}, g.GetTargetsOf(idThree))
	testutils.AssertEquals(t, []int{idTwo, idFive}, g.GetSourcesOf(idThree))

	testutils.AssertEquals(t, []int{idFive, idTwo}, g.GetTargetsOf(idFour))
	testutils.AssertEquals(t, []int{}, g.GetSourcesOf(idFour))

	testutils.AssertEquals(t, []int{idThree}, g.GetTargetsOf(idFive))
	testutils.AssertEquals(t, []int{idThree, idFour}, g.GetSourcesOf(idFive))

	testutils.AssertEquals(t, Vertex{name: "two"}, g.GetVertexValue(idTwo))
	testutils.AssertEquals(t, Vertex{name: "three"}, g.GetVertexValue(idThree))
	testutils.AssertEquals(t, Vertex{name: "four"}, g.GetVertexValue(idFour))
	testutils.AssertEquals(t, Vertex{name: "five"}, g.GetVertexValue(idFive))

	testutils.AssertEquals(t, Edge{name: "edgeTwo"}, g.GetEdgeValue(idTwo, idThree))
	testutils.AssertEquals(t, Edge{name: "edgeThree"}, g.GetEdgeValue(idThree, idFive))
	testutils.AssertEquals(t, Edge{name: "edgeFour"}, g.GetEdgeValue(idFour, idFive))
	testutils.AssertEquals(t, Edge{name: "edgeFive"}, g.GetEdgeValue(idFive, idThree))
	testutils.AssertEquals(t, Edge{name: "edgeSix"}, g.GetEdgeValue(idFour, idTwo))
}
