package pkg

import (
	"reflect"
	"testing"
)

func TestGraph_AddVertex(t *testing.T) {
	g := NewGraph()

	g.AddVertex("A")
	g.AddVertex("B")

	if len(g.vertices) != 2 {
		t.Errorf("expected 2 vertices got %d", len(g.vertices))
	}
	if _, exists := g.vertices["A"]; !exists {
		t.Errorf("expected vertex a to exist")
	}
	if _, exists := g.vertices["B"]; !exists {
		t.Errorf("expected vertex b to exist")
	}
}

func TestGraph_AddEdge(t *testing.T) {
	g := NewGraph()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B")

	if len(g.vertices["A"].Edges) != 1 {
		t.Errorf("expected 1 edge from vertex a got %d", len(g.vertices["A"].Edges))
	}
	if _, exists := g.vertices["A"].Edges["B"]; !exists {
		t.Errorf("expected an edge from a to b")
	}
}

func TestGraph_QueryGrandparents(t *testing.T) {
	g := NewGraph()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")
	g.AddEdge("D", "E")

	grandparentsA := g.QueryGrandparents("A")
	expectedGrandparentsA := []string{"D", "E"}

	if !reflect.DeepEqual(grandparentsA, expectedGrandparentsA) {
		t.Errorf("expected grandparents of a to be %v got %v", expectedGrandparentsA, grandparentsA)
	}
}

func TestGraph_QuerySiblings(t *testing.T) {
	g := NewGraph()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")
	g.AddEdge("D", "E")

	siblingsA := g.QuerySiblings("A")
	expectedSiblingsA := []string{"D", "E"}

	if !reflect.DeepEqual(siblingsA, expectedSiblingsA) {
		t.Errorf("expected siblings of a to be %v got %v", expectedSiblingsA, siblingsA)
	}
}

func TestGraph_QueryCousins(t *testing.T) {
	g := NewGraph()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")
	g.AddEdge("D", "E")

	cousinsA := g.QueryCousins("A")
	expectedCousinsA := []string{"E"}

	if !reflect.DeepEqual(cousinsA, expectedCousinsA) {
		t.Errorf("expected cousins of a to be %v got %v", expectedCousinsA, cousinsA)
	}
}

func TestGraph_FilterVertices(t *testing.T) {
	g := NewGraph()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")

	filteredVertices := g.FilterVertices(func(v *Vertex) bool {
		return v.Name == "A" || v.Name == "C"
	})
	expectedNames := []string{"A", "C"}

	if len(filteredVertices) != len(expectedNames) {
		t.Errorf("expected %d vertices got %d", len(expectedNames), len(filteredVertices))
	}

	for i, vertex := range filteredVertices {
		if vertex.Name != expectedNames[i] {
			t.Errorf("expected vertex %s got %s", expectedNames[i], vertex.Name)
		}
	}
}
