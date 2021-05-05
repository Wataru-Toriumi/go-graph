package main

import (
	"fmt"
	"github.com/Wataru-Toriumi/go-graph/src/graph"
)

func main() {
	fmt.Println("Hello, there!!")
	var nodes []string = []string{"1","2","3"}
	var edges [][]string = [][]string{{"1","2"},{"1","3"}}
	g := graph.New(nodes, edges)
	g.Show()
}