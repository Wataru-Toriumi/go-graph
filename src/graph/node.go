package graph

type Node struct {
	Id string
	Weight map[string] int
}

func NewNode(id string, weight map[string] int) *Node{
	return &Node{Id: id, Weight: weight}
}
