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
	Key        string
	LinkedTo   []*Node
	PathLength int
	Prev       *Node
	Temp       bool
}

// Edge : Data type for edges
type Edge struct {
	StartKey, EndKey string
	Start, End       *Node
	Cost             int
	Visited          bool
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

// addEdge : Adds a edges to the graph
func (graph *Graph) addEdge(newEdge *Edge) {
	newArr := make([]*Edge, len(graph.Edges)+1)
	if len(graph.Edges) == 0 {
		newArr[0] = newEdge
		graph.Edges = newArr
		return
	}
	for i := 0; i < len(graph.Edges); i++ {
		newArr[i] = graph.Edges[i]
	}
	newArr[len(graph.Edges)] = newEdge
	graph.Edges = newArr
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

// createEdges : Creates edges
func createEdge(start, end *Node) *Edge {
	edge := new(Edge)
	edge = &Edge{
		StartKey: start.Key,
		EndKey:   end.Key,
		Start:    start,
		End:      end,
	}
	return edge
}

func (graph *Graph) obtainIndex(node *Node) int {
	for i := 0; i < len(graph.Edges); i++ {
		if graph.Edges[i].StartKey == node.Key {
			return i
		}
	}
	return -1
}

func (graph *Graph) lowestPath() (node *Node) {
	minPath := graph.Nodes[0]
	for i := 0; i < len(graph.Nodes); i++ {
		if graph.Nodes[i].PathLength < minPath.PathLength {
			minPath = graph.Nodes[i]
		}
	}
	minPath.Temp = false
	return minPath
}

func (graph *Graph) dijkstra(startNode, endNode *Node) {
	for i := 0; i < len(graph.Nodes); i++ {
		graph.Nodes[i].PathLength = 10000000
		graph.Nodes[i].Prev = nil
		graph.Nodes[i].Temp = true
	}
	graph.Nodes[graph.obtainIndex(startNode)].PathLength = 0
	current := graph.lowestPath()
	fmt.Println(current)
}

func main() {
	graph := new(Graph)
	keys := [7]string{"a", "b", "c", "d", "e", "f", "g"}
	costs := [9]int{2, 1, 1, 2, 5, 7, 6, 9, 8}
	for i := 0; i < len(keys); i++ {
		graph.addNode(createNode(keys[i]))
	}
	connections := [7][]int{{1, 6}, {2, 4, 3}, {4}, {4}, {5}, {-1}, {3}}
	for i := 0; i < len(graph.Nodes); i++ {
		for j := 0; j < len(connections[i]); j++ {
			if connections[i][j] > -1 {
				graph.Nodes[i].createLink(graph.Nodes[connections[i][j]])
				graph.addEdge(createEdge(graph.Nodes[i], graph.Nodes[connections[i][j]]))
			}
		}
	}
	for i := 0; i < len(graph.Edges); i++ {
		graph.Edges[i].Cost = costs[i]
	}
	for i := 0; i < len(graph.Nodes); i++ {
		fmt.Println(graph.Nodes[i])
	}
	graph.dijkstra(graph.Nodes[0], graph.Nodes[3])
}
