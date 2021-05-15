package graph

type Edge struct {
	From string
	To string
	Weight map[string] int
}

func NewEdges(from string, to string, weight map[string] int) *Edge {
	return &Edge{From: from, To: to, Weight: weight}
}
