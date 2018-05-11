package main

import (
	"fmt"
)

// Graph : Data type for graphs
type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

// Node : Data type for nodes
type Node struct {
	Key      string
	LinkedTo []*Node
}

// Edge : aksjsdlakj
type Edge struct {
	Start, End *Node
	Cost       int
}

// addNode : Adds a node to the graph
func (graph *Graph) addNode(newNode *Node) {
	newArr := make([]*Node, len(graph.Nodes)+1)
	if len(graph.Nodes) == 0 {
		newArr[0] = newNode
		graph.Nodes = newArr
		return
	}
	for i := 0; i < len(graph.Nodes); i++ {
		newArr[i] = graph.Nodes[i]
	}
	newArr[len(graph.Nodes)] = newNode
	graph.Nodes = newArr
}

// createNode : Creates and returns a new node
func createNode(key string) *Node {
	node := new(Node)
	node = &Node{
		Key: key,
	}
	return node
}

func main() {
	graph := new(Graph)
	arr := [3]string{"a", "b", "c"}
	for i := 0; i < len(arr); i++ {
		graph.addNode(createNode(arr[i]))
	}
	fmt.Println(graph.Nodes[0])
	fmt.Println(graph.Nodes[1])
	fmt.Println(graph.Nodes[2])
}
