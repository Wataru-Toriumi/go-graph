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
	g := graph.New(nodes, edges)

	t.Run("get nodes", func(t *testing.T) {
		if (reflect.DeepEqual(g.Get_nodes, nodes)) {
			t.Errorf("Not adequate nodes.")
		} else {
			t.Logf("get nodes test succeeded.")
		}
	})

	t.Run("get edges", func(t *testing.T) {
		if (reflect.DeepEqual(g.Get_edges, edges)) {
			t.Errorf("Not adequate edges.")
		} else {
			t.Logf("get edges test succeeded.")
		}
	})
}