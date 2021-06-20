package graph_test

import (
	"testing"
	"reflect"
	"github.com/Wataru-Toriumi/go-graph/src/graph"
)

func TestGraphNew(t *testing.T) {

	t.Logf("Set up")

	var nodes []graph.Node = []graph.Node{
		{Id: 1, Weight: map[string]interface{}{"weight": 1}},
		{Id: 2, Weight: map[string]interface{}{"weight": 1}},
		{Id: 3, Weight: map[string]interface{}{"weight": 1}},
	}
	var ns [][2]interface{} = [][2]interface{}{
		{1, 1},
		{2, 1},
		{3, 1},
	}
	var edges []graph.Edge = []graph.Edge{
		{From: 1, To: 2, Weight: map[string]interface{}{"weight": 1}},
		{From: 1, To: 3, Weight: map[string]interface{}{"weight": 1}},
	}

	var es [][3]interface{} = [][3]interface{}{
		{1, 2, 1},
		{1, 3, 1},
	}
	var g = graph.New(ns, es)

	t.Run("get nodes", func(t *testing.T) {
		gotten_nodes := reflect.DeepEqual(graph.Get_nodes(g), nodes)
		t.Logf("%v", graph.Get_nodes(g))
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

	var new_nodes []graph.Node = []graph.Node{
		{Id: 4, Weight: map[string]interface{}{"weight": 1}},
	}
	var new_edges []graph.Edge = []graph.Edge{
		{From: 1, To: 4, Weight: map[string]interface{}{"weight": 1}},
		{From: 2, To: 4, Weight: map[string]interface{}{"weight": 1}},
	}

	t.Run("add nodes", func(t *testing.T) {
		expected_nodes := []graph.Node{
			{Id: 1, Weight: map[string]interface{}{"weight": 1}},
			{Id: 2, Weight: map[string]interface{}{"weight": 1}},
			{Id: 3, Weight: map[string]interface{}{"weight": 1}},
			{Id: 4, Weight: map[string]interface{}{"weight": 1}},
		}
		graph.Add_nodes(g, new_nodes)
		added_nodes := reflect.DeepEqual(graph.Get_nodes(g), expected_nodes)

		if (added_nodes) {
			t.Logf("succeeded to add the nodes to graph")
		} else {
			t.Errorf("Not adequate nodes")
		}
	})

	t.Run("add edges", func(t *testing.T) {
		expected_edges := []graph.Edge{
			{From: 1, To: 2, Weight: map[string]interface{}{"weight": 1}},
			{From: 1, To: 3, Weight: map[string]interface{}{"weight": 1}},
			{From: 1, To: 4, Weight: map[string]interface{}{"weight": 1}},
			{From: 2, To: 4, Weight: map[string]interface{}{"weight": 1}},
		}
		graph.Add_edges(g, new_edges)
		added_nodes := reflect.DeepEqual(graph.Get_edges(g), expected_edges)

		if (added_nodes) {
			t.Logf("succeeded to add the edges to graph")
		} else {
			t.Errorf("Not adequate nodes")
		}
	})
}