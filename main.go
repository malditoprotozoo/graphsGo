package main

import (
	"fmt"
)

// Graph : Data type for graphs
type Graph struct {
	Nodes []*Node
}

// Node : Data type for nodes
type Node struct {
	Key   string
	Links []*Node
}

// ShowAllKeys : Show all keys
func (graph *Graph) ShowAllKeys() {
	for i := 0; i < len(graph.Nodes); i++ {
		fmt.Println(graph.Nodes[i])
	}
}

// addEdge : Add new links between nodes
func (node *Node) addEdge(newNode *Node) {
	// if edge already exists
	for i := 0; i < len(node.Links); i++ {
		if node.Links[i] != newNode {
			node.Links = append(node.Links, newNode)
		}
	}
	return
}

// addNode : Add new node to graph
func (graph *Graph) addNode(newNode *Node) {
	for i := 0; i < len(graph.Nodes); i++ {
		if graph.Nodes[i] != newNode {
			graph.Nodes = append(graph.Nodes, newNode)
		}
	}
	return
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
	graph.ShowAllKeys()
}
