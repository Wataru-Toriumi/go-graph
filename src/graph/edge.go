package graph

type Edge struct {
	From interface{}
	To interface{}
	Weight map[string] interface{}
}

func NewEdges(from interface{}, to interface{}, weight map[string] interface{}) *Edge {
	return &Edge{From: from, To: to, Weight: weight}
}
