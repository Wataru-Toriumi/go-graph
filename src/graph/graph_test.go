package graph_test

import (
	"testing"
	"reflect"
	"github.com/Wataru-Toriumi/go-graph/src/graph"
)

func TestGraphNew(t *testing.T) {

	t.Logf("Set up")

	var nodes []string = []string{"1", "2", "3"}
	var edges [][]string = [][]string{{"1", "2"},{"1", "3"}}
	var g = graph.New(nodes, edges)

	t.Run("get nodes", func(t *testing.T) {
		gotten_nodes := reflect.DeepEqual(graph.Get_nodes(g), nodes)

		if (gotten_nodes) {
			t.Logf("get nodes test succeeded.")
		} else {
			t.Errorf("Not adequate nodes.")
		}
	})

	t.Run("get edges", func(t *testing.T) {
		gotten_edges := reflect.DeepEqual(graph.Get_edges(g), edges)
		if (gotten_edges) {
			t.Logf("get edges test succeeded.")
		} else {
			t.Errorf("Not adequate edges.")
		}
	})

	var new_nodes []string = []string{"4"}
	var new_edges [][]string = [][]string{{"1", "4"}, {"2", "4"}}

	t.Run("add nodes", func(t *testing.T) {
		expected_nodes := []string{"1", "2", "3", "4"}
		graph.Add_nodes(g, new_nodes)
		added_nodes := reflect.DeepEqual(graph.Get_nodes(g), expected_nodes)

		if (added_nodes) {
			t.Logf("succeeded to add the nodes to graph")
		} else {
			t.Errorf("Not adequate nodes")
		}
	})

	t.Run("add edges", func(t *testing.T) {
		expected_edges := [][]string{{"1", "2"}, {"1", "3"}, {"1",  "4"}, {"2", "4"}}
		graph.Add_edges(g, new_edges)
		added_nodes := reflect.DeepEqual(graph.Get_edges(g), expected_edges)

		if (added_nodes) {
			t.Logf("succeeded to add the edges to graph")
		} else {
			t.Errorf("Not adequate nodes")
		}
	})
}