package graph

type Node struct {
	Id interface{}
	Weight map[string] interface{}
}

func NewNode(id interface{}, weight map[string] interface{}) *Node{
	return &Node{Id: id, Weight: weight}
}
