package main

import (
	"fmt"
	"sort"
)

type Node struct {
	nodeID string
}

type Edge struct {
	Parent *Node
	Child  *Node
	Cost   int
}

// Graph is a collection of Edges and Nodes, represents the problem itself
type Graph struct {
	Edges []*Edge
	Nodes []*Node
}

// ___________ cost 2
// Link between two nodes ex: A --------- B
func (graph *Graph) AddEdge(parent, child *Node, cost int) {
	// Create a edge with  (Parent) A --------- B (Child) and cost between these two edges
	edge := &Edge{
		Parent: parent,
		Child:  child,
		Cost:   cost,
	}

	graph.Edges = append(graph.Edges, edge)
	graph.AddNode(parent)
	graph.AddNode(child)
}

func (graph *Graph) AddNode(nodeToAdd *Node) {
	var alreadyIn bool

	for _, _node := range graph.Nodes {
		if _node == nodeToAdd {
			alreadyIn = true
		}
	}

	if !alreadyIn {
		graph.Nodes = append(graph.Nodes, nodeToAdd)
	}
}

func (graph *Graph) Djikistra(startNode *Node) (shortestPath string) {

	// Initialize startNode = 0 and the other nodes with Infinity
	costTable := graph.newCostTable(startNode)

	// All visited nodes
	var visited []*Node

	// A loop to visit all Nodes
	for len(visited) != len(graph.Nodes) {

		//
		node := graph.closestNonVisited(costTable, visited)

		visited = append(visited, node)

		nodeEdges := graph.getNodeEdges(node)

		for _, edge := range nodeEdges {
			// fmt.Printf("Node: %s - CostTable[node]: %d  - EdgeCost: %d - CostTable[edgechild]: %d - EdgeChild: %s\n", node.nodeID, costTable[node], edge.Cost, costTable[edge.Child], edge.Child.nodeID)

			distanceToNeighbor := costTable[node] + edge.Cost
			// fmt.Printf("Distance to neighbour: %d  - costTable[edge.Child]: %d\n", distanceToNeighbor, costTable[edge.Child])

			// If the distance to the neighbor is less than the current value of the child node, then this is the shortest path so far
			if distanceToNeighbor < costTable[edge.Child] {

				// Update the costTable for that neighbor
				costTable[edge.Child] = distanceToNeighbor
				//	fmt.Printf("now node: %s, has value %d\n", edge.Child.nodeID, costTable[edge.Child])
				//	fmt.Printf("\n")
			}
		}
	}
	for node, cost := range costTable {
		shortestPath += fmt.Sprintf("Distance from %s to %s = %d\n", startNode.nodeID, node.nodeID, cost)
	}

	return shortestPath

}

func (graph *Graph) newCostTable(startNode *Node) map[*Node]int {
	costTable := make(map[*Node]int)

	costTable[startNode] = 0

	for _, node := range graph.Nodes {
		if node != startNode {
			costTable[node] = 99999
		}
	}

	return costTable

}

func (graph *Graph) closestNonVisited(costTable map[*Node]int, visited []*Node) *Node {

	type CostTableToSort struct {
		Node *Node
		Cost int
	}
	var sorted []CostTableToSort

	// Verify if the Node has been visited already
	for node, cost := range costTable {
		var isVisited bool
		for _, visitedNode := range visited {
			if node == visitedNode {
				//	fmt.Printf("node: %s - cost: %d\n - visited: %t\n", node.nodeID, cost, !isVisited)
				isVisited = true
			}
		}
		// If not, add them to the sorted slice
		if !isVisited {
			// fmt.Printf("node: %s - cost: %d\n - visited: %t\n", node.nodeID, cost, isVisited)
			sorted = append(sorted, CostTableToSort{node, cost})
		}
	}

	// We need the Node with the lower cost from the costTable
	// So it's important to sort it
	// Here I'm using an anonymous struct to make it easier to sort a map
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Cost < sorted[j].Cost
	})

	return sorted[0].Node
}

func (graph *Graph) getNodeEdges(node *Node) []*Edge {
	var edges []*Edge

	// If the edge is connected with the requested node, then they're linked
	for _, edge := range graph.Edges {
		if edge.Parent == node {
			edges = append(edges, edge)
		}
	}
	return edges
}
func main() {

	a := &Node{nodeID: "a"}
	b := &Node{nodeID: "b"}
	c := &Node{nodeID: "c"}
	d := &Node{nodeID: "d"}
	e := &Node{nodeID: "e"}
	f := &Node{nodeID: "f"}
	g := &Node{nodeID: "g"}

	graph := Graph{}
	graph.AddEdge(a, c, 2)
	graph.AddEdge(a, b, 5)
	graph.AddEdge(c, b, 1)
	graph.AddEdge(c, d, 9)
	graph.AddEdge(b, d, 4)
	graph.AddEdge(d, e, 2)
	graph.AddEdge(d, g, 30)
	graph.AddEdge(d, f, 10)
	graph.AddEdge(f, g, 1)

	fmt.Println(graph.Djikistra(a))
}
