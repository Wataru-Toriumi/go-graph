package graph

import (
	"fmt"
	"strconv"
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

func New (nodes [][2]string, edges [][3]string) *Graph {
	var ns []Node = []Node{}
	var es []Edge = []Edge{}

	for _, node := range nodes {
		weight, _ := strconv.Atoi(node[1])
		n := NewNode(node[0], map[string]int{"weight": weight})
		ns = append(ns, *n)
	}

	for _, edge := range edges {
		weight, _ := strconv.Atoi(edge[2])
		e := NewEdges(edge[0], edge[1], map[string]int{"weight": weight})
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