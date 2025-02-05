package bigraph

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"	
)

type Graph struct {
	adjacencyList map[int]map[int]bool
}

// NewGraph initializes a new Graph.
func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[int]map[int]bool)}
}

// AddVertex adds a vertex to the graph.
func (g *Graph) AddVertex(v int) {
	if _, exists := g.adjacencyList[v]; !exists {
		g.adjacencyList[v] = make(map[int]bool)
	}
}

// RemoveVertex removes a vertex and all its edges.
func (g *Graph) RemoveVertex(v int) {
	delete(g.adjacencyList, v)
	for _, neighbors := range g.adjacencyList {
		delete(neighbors, v)
	}
}

// AddEdge adds an edge between two vertices.
func (g *Graph) AddEdge(v1, v2 int) {
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.adjacencyList[v1][v2] = true
	g.adjacencyList[v2][v1] = true
}

// RemoveEdge removes an edge between two vertices.
func (g *Graph) RemoveEdge(v1, v2 int) {
	delete(g.adjacencyList[v1], v2)
	delete(g.adjacencyList[v2], v1)
}

// IsIsolated checks if a vertex has no edges.
func (g *Graph) IsIsolated(v int) bool {
	if neighbors, exists := g.adjacencyList[v]; exists {
		return len(neighbors) == 0
	}
	return false
}

// LoadFromAdjacencyFile loads the graph from a file with adjacency list format.
func (g *Graph) LoadFromAdjacencyFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) == 0 {
			continue
		}

		v, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		g.AddVertex(v)

		for _, neighbor := range parts[1:] {
			n, err := strconv.Atoi(neighbor)
			if err != nil {
				return err
			}
			g.AddEdge(v, n)
		}
	}

	return scanner.Err()
}

// ShortestPath finds the shortest distance between two vertices using BFS.
func (g *Graph) ShortestPath(start, end int) (int, error) {
	if _, exists := g.adjacencyList[start]; !exists {
		return -1, errors.New("start vertex not found")
	}
	if _, exists := g.adjacencyList[end]; !exists {
		return -1, errors.New("end vertex not found")
	}

	queue := []int{start}
	visited := map[int]bool{start: true}
	distance := map[int]int{start: 0}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		if current == end {
			return distance[current], nil
		}

		for neighbor := range g.adjacencyList[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				distance[neighbor] = distance[current] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	return -1, errors.New("no path found")
}

