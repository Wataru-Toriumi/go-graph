package graph

import (
	"fmt"
)

type Graph struct {
	nodes []Node
	edges []Edge
}

type GraphInterface interface {
	getNodes() []Node
	getEdges() []Edge
	addNodes([]Node)
	addEdges([]Edge)
	show()
}

func (G Graph) getNodes() []Node {
	return G.nodes
}

func (G Graph) getEdges() []Edge {
	return G.edges
}

func (G *Graph) addNodes(nodes []Node) {
	G.nodes = append(G.nodes, nodes...)
}

func (G *Graph) addEdges(edges []Edge) {
	G.edges = append(G.edges, edges...)
}

func (G Graph) show() {
	fmt.Printf("nodes: %v\n", G.nodes)
	fmt.Printf("edges: %v\n", G.edges)
}

func New (nodes [][2]interface{}, edges [][3]interface{}) *Graph {
	var ns []Node = []Node{}
	var es []Edge = []Edge{}

	for _, node := range nodes {
		n := NewNode(node[0], map[string]interface{}{"weight": node[1]})
		ns = append(ns, *n)
	}

	for _, edge := range edges {
		e := NewEdges(edge[0], edge[1], map[string]interface{}{"weight": edge[2]})
		es = append(es, *e)
	}
	return &Graph{ns, es}
}


func Get_nodes(g GraphInterface) []Node {
	return g.getNodes()
}

func Get_edges(g GraphInterface) []Edge {
	return g.getEdges()
}

func Add_nodes(g GraphInterface, nodes []Node) {
	g.addNodes(nodes)
}

func Add_edges(g GraphInterface, edges []Edge) {
	g.addEdges(edges)
}

func Show(g GraphInterface) {
	g.show()
}