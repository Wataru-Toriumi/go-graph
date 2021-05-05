package graph

import (
	"fmt"
)

type Graph struct {
	nodes []string
	edges [][]string
}

type GraphInterface interface {
	get_nodes() []string
	get_edges() [][]string
	add_nodes([]string)
	add_edges([][]string)
	show()
}

func (G Graph) get_nodes() []string {
	return G.nodes
}

func (G Graph) get_edges() [][]string {
	return G.edges
}

func (G *Graph) add_nodes(nodes []string) {
	G.nodes = append(G.nodes, nodes...)
}

func (G *Graph) add_edges(edges [][]string) {
	G.edges = append(G.edges, edges...)
}

func (G Graph) show() {
	fmt.Printf("nodes: %v\n", G.nodes)
	fmt.Printf("edges: %v\n", G.edges)
}

func New (nodes []string, edges [][]string) *Graph {
	return &Graph{nodes, edges}
}


func Get_nodes(g GraphInterface) []string {
	return g.get_nodes()
}

func Get_edges(g GraphInterface) [][]string {
	return g.get_edges()
}

func Add_nodes(g GraphInterface, nodes []string) {
	g.add_nodes(nodes)
}

func Add_edges(g GraphInterface, edges [][]string) {
	g.add_edges(edges)
}

func Show(g GraphInterface) {
	g.show()
}