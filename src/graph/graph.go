package graph

import (
	"fmt"
)

type Node struct {
	Id string
	Weight map[string] int
}

type Edge struct {
	From string
	To string
	Weight map[string] int
}

type Graph struct {
	nodes []Node
	edges []Edge
}

type GraphInterface interface {
	get_nodes() []Node
	get_edges() []Edge
	add_nodes([]Node)
	add_edges([]Edge)
	show()
}

func (G Graph) get_nodes() []Node {
	return G.nodes
}

func (G Graph) get_edges() []Edge {
	return G.edges
}

func (G *Graph) add_nodes(nodes []Node) {
	G.nodes = append(G.nodes, nodes...)
}

func (G *Graph) add_edges(edges []Edge) {
	G.edges = append(G.edges, edges...)
}

func (G Graph) show() {
	fmt.Printf("nodes: %v\n", G.nodes)
	fmt.Printf("edges: %v\n", G.edges)
}

func New (nodes []Node, edges []Edge) *Graph {
	return &Graph{nodes, edges}
}


func Get_nodes(g GraphInterface) []Node {
	return g.get_nodes()
}

func Get_edges(g GraphInterface) []Edge {
	return g.get_edges()
}

func Add_nodes(g GraphInterface, nodes []Node) {
	g.add_nodes(nodes)
}

func Add_edges(g GraphInterface, edges []Edge) {
	g.add_edges(edges)
}

func Show(g GraphInterface) {
	g.show()
}