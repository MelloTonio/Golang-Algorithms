package BFS

import (
	"container/list"
	"fmt"
)

const (
	SIZE = 5 // x
)

type Node struct {
	x      int
	y      int
	dist   int
	Parent *Node
}

func isValid(mat [][]string, visited [SIZE][SIZE]bool, row int, col int) bool {
	return (row >= 0) && (row < SIZE) && (col >= 0) && (col < SIZE) && mat[row][col] == " " && !visited[row][col]
}

func BFS(mat [][]string, i int, j int, x int, y int) {

	row := []int{-1, 0, 0, 1}
	col := []int{0, -1, 1, 0}

	// Visited is a SIZExSIZE Slice of bool, indicates wich (x,y) have been visited
	var visited [SIZE][SIZE]bool

	queue := list.New()

	// Start with initial values visited
	visited[i][j] = true

	// Push the first Node into the queue
	queue.PushFront(&Node{x: i, y: j, dist: 0, Parent: nil})

	min_dist := 9999

	var node *list.Element

	//var helper *list.Element

	for queue.Len() > 0 {

		// Push a node from the front
		node = queue.Back()
		//	helper = node

		// Catch the node values
		i := node.Value.(*Node).x
		j := node.Value.(*Node).y
		dist := node.Value.(*Node).dist

		// If node values == requested values, then we found the path
		if i == x && j == y {
			min_dist = dist
			break
		}

		// Pop node from the queue
		queue.Remove(node)

		// 4 possible moves
		for k := 0; k < 4; k++ {
			// check if it is possible to go to position
			// (i + row[k], j + col[k]) from current position
			if isValid(mat, visited, i+row[k], j+col[k]) {
				// mark next cell as visited and enqueue it
				visited[i+row[k]][j+col[k]] = true
				queue.PushFront(&Node{x: i + row[k], y: j + col[k], dist: dist + 1, Parent: node.Value.(*Node)})
				// Just a helper dont mind it
				//helper = append(helper, node.Value.(*Node))
			}
		}

	}

	if min_dist != 9999 {
		fmt.Println("The shortest path from source to destination ", "has length ", min_dist)
		printPath(node, mat)

	} else {
		fmt.Println("Destination can't be reached from source")
	}
}

func printPath(node *list.Element, mat [][]string) {
	newNode := node.Value.(*Node)

	for newNode != nil {

		mat[newNode.x][newNode.y] = "o"

		newNode = newNode.Parent
	}

	for _, v := range mat {
		fmt.Println(v)
	}
}

func Run(x, y int) {
	// If we change the maze, we must change "CONST SIZE"
	SliceOfSlices := [][]string{
		//x
		{" ", " ", " ", "#", " ", " ", " "}, // y
		{" ", "#", " ", " ", " ", "#", " "},
		{" ", "#", " ", " ", " ", " ", " "},
		{" ", " ", "#", "#", " ", " ", " "},
		{"#", " ", "#", " ", " ", "#", " "},
	}

	BFS(SliceOfSlices, 0, 0, x, y)
}

/*
{" ", " ", " ", " ", " ", "#", "#", " ", " ", " "},
{"#", " ", " ", " ", " ", " ", "#", " ", "#", " "},
{"#", "#", " ", "#", " ", " ", " ", "#", "#", " "},
{" ", "#", " ", " ", " ", "#", " ", " ", "#", " "},
{"#", "#", "#", " ", "#", "#", "#", " ", "#", " "},
{" ", "#", " ", " ", " ", "#", "#", " ", " ", "#"},
{"#", "#", "#", "#", " ", "#", "#", " ", "#", " "},
{"#", " ", " ", " ", " ", " ", " ", " ", "#", "#"},
{" ", " ", " ", " ", " ", "#", "#", " ", " ", " "},
{"#", "#", " ", "#", "#", " ", " ", "#", "#", " "},
*/
