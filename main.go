package main

import (
	"fmt"
)

// Graph : asda
type Graph struct {
	Nodes []*Node
}

// Node : asdasda
type Node struct {
	Key string
}

func main() {
	a := new(Node)
	a = &Node{
		Key: "a",
	}
	b := new(Node)
	b = &Node{
		Key: "b",
	}
	c := new(Node)
	c = &Node{
		Key: "c",
	}
	nodes := make([]*Node, 3)
	nodes[0] = a
	nodes[1] = b
	nodes[2] = c
	graph := new(Graph)
	graph = &Graph{
		Nodes: nodes,
	}
	fmt.Println(graph.Nodes)
}
