package graph_test

import (
	"testing"
	"reflect"
	"github.com/Wataru-Toriumi/go-graph/src/graph"
)

func TestGraphNew(t *testing.T) {

	t.Logf("Set up")

	var nodes []string = []string{"1","2","3"}
	var edges [][]string = [][]string{{"1","2"},{"1","3"}}
	var g = graph.New(nodes, edges)

	t.Run("get nodes", func(t *testing.T) {
		if (reflect.DeepEqual(graph.Get_nodes(g), nodes)) {
			t.Logf("get nodes test succeeded.")
		} else {
			t.Errorf("Not adequate nodes.")
		}
	})

	t.Run("get edges", func(t *testing.T) {
		if (reflect.DeepEqual(graph.Get_edges(g), edges)) {
			t.Logf("get edges test succeeded.")
		} else {
			t.Errorf("Not adequate edges.")
		}
	})
}