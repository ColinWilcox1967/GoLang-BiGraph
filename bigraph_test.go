package bigraph

import "testing"

func TestGraphOperations(t *testing.T) {
	g := NewGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddEdge(1, 2)

	if g.IsIsolated(1) {
		t.Errorf("Vertex 1 should not be isolated")
	}

	g.RemoveEdge(1, 2)
	if !g.IsIsolated(1) || !g.IsIsolated(2) {
		t.Errorf("Vertices 1 and 2 should be isolated")
	}

	g.RemoveVertex(1)
	if _, exists := g.adjacencyList[1]; exists {
		t.Errorf("Vertex 1 should have been removed")
	}
}

func TestShortestPath(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	distance, err := g.ShortestPath(1, 4)
	if err != nil || distance != 3 {
		t.Errorf("Expected shortest path 3, got %d, error: %v", distance, err)
	}
}

func TestLoadFromAdjacencyFile(t *testing.T) {
	g := NewGraph()
	err := g.LoadFromAdjacencyFile("test_adj_list.txt")
	if err != nil {
		t.Errorf("Error loading adjacency file: %v", err)
	}
}

func TestRemoveNonExistentVertex(t *testing.T) {
	g := NewGraph()
	g.RemoveVertex(10)
}

func TestRemoveNonExistentEdge(t *testing.T) {
	g := NewGraph()
	g.RemoveEdge(1, 2)
}
