package graph

import (
	"fmt"
)

type Graph struct {
	nodes []string
	edges [][]string
}

type GraphInterface interface {
	Nodes() []string
	Edges() [][]string
	Show()
}

func (G Graph) Show() {
	fmt.Printf("nodes: %v\n", G.nodes)
	fmt.Printf("edges: %v\n", G.edges)
}

func (G Graph) Nodes() []string {
	return G.nodes
}

func (G Graph) Edges() [][]string {
	return G.edges
}

func New (nodes []string, edges [][]string) *Graph {
	return &Graph{nodes, edges}
}

func Get_nodes(g GraphInterface) []string {
	return g.Nodes()
}

func Get_edges(g GraphInterface) [][]string {
	return g.Edges()
}