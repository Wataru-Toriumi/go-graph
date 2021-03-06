package graph

import (
	"fmt"
)

type Graph struct {
	nodes []string
	edges [][]string
}

func (G Graph) Show_graph() {
	fmt.Printf("nodes: %v\n", G.nodes)
	fmt.Printf("edges: %v\n", G.edges)
}

func New (nodes []string, edges [][]string) *Graph {
	return &Graph{nodes, edges}
}