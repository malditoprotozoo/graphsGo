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

// Edge : Data type for edges
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

// createLink : Creates a single connection
func (node *Node) createLink(link *Node) {
	newArr := make([]*Node, len(node.LinkedTo)+1)
	if len(node.LinkedTo) == 0 {
		newArr[0] = link
		node.LinkedTo = newArr
		return
	}
	for i := 0; i < len(node.LinkedTo); i++ {
		newArr[i] = node.LinkedTo[i]
	}
	newArr[len(node.LinkedTo)] = link
	node.LinkedTo = newArr
}

func main() {
	graph := new(Graph)
	// 								0,  1,		2,		3,	4,	5,	6
	arr := [7]string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i < len(arr); i++ {
		graph.addNode(createNode(arr[i]))
	}
	connections := [7][]int{{1, 6}, {2, 4, 3}, {4}, {4}, {5}, {-1}, {3}}
	for i := 0; i < len(graph.Nodes); i++ {
		for j := 0; j < len(connections[i]); j++ {
			if connections[i][j] > -1 {
				graph.Nodes[i].createLink(graph.Nodes[connections[i][j]])
			}
		}
	}
	fmt.Println(graph.Nodes[0].LinkedTo[0].Key)
}
