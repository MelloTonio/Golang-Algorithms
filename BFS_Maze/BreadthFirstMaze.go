package BFS

import (
	"container/list"
	"fmt"
)

const (
	M = 10
	N = 10
)

type Node struct {
	x      int
	y      int
	dist   int
	Parent *Node
}

func isValid(mat [][]string, visited [M][N]bool, row int, col int) bool {
	return (row >= 0) && (row < M) && (col >= 0) && (col < N) && mat[row][col] == " " && !visited[row][col]
}

func BFS(mat [][]string, i int, j int, x int, y int) {

	row := []int{-1, 0, 0, 1}
	col := []int{0, -1, 1, 0}

	// Visited is a SIZExSIZE Slice of bool, indicates wich (x,y) have been visited
	var visited [M][N]bool

	queue := list.New()

	// Start with initial values visited
	visited[i][j] = true

	// Push the first Node into the queue
	queue.PushFront(&Node{x: i, y: j, dist: 0, Parent: nil})

	// All the distances before (0,0) are unknown, so they are 9999 (infinity)
	min_dist := 9999

	var node *list.Element

	//var helper *list.Element

	for queue.Len() > 0 {

		// Push a node from the back
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
			// (i + row[k], j + col[k]) -> it only increases or decrease 1 in both x,y
			if isValid(mat, visited, i+row[k], j+col[k]) {
				// mark next cell as visited and enqueue it
				visited[i+row[k]][j+col[k]] = true
				// enqueue all the valid parent nodes
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
	// Get the node that reached the requested (x,y)
	newNode := node.Value.(*Node)

	lastNodeX := newNode.x
	lastNodeY := newNode.y
	// directions := []string{}

	// Backtracking - Go from top to down, all the way until we reach (0,0)
	for newNode != nil {

		switch true {

		// y + 1
		case newNode.y > lastNodeY:
			mat[newNode.x][newNode.y] = "←"
			// directions = append(directions, "d")
			lastNodeY = newNode.y

		// x - 1
		case newNode.x < lastNodeX:
			mat[newNode.x][newNode.y] = "↓"
			// 	directions = append(directions, "c")
			lastNodeX = newNode.x

		// y - 1
		case newNode.y < lastNodeY:
			mat[newNode.x][newNode.y] = "➜"
			// directions = append(directions, "e")
			lastNodeY = newNode.y

		// x + 1
		case newNode.x > lastNodeX:
			mat[newNode.x][newNode.y] = "↑"
			lastNodeX = newNode.x

		// Result
		default:
			mat[newNode.x][newNode.y] = "⟰"
		}

		// mat[newNode.x][newNode.y] = "o"

		// fmt.Printf("%d %d\n", newNode.x, newNode.y)
		newNode = newNode.Parent
	}

	// Print the final path
	for _, v := range mat {
		fmt.Println(v)
	}

}

func Run(x, y int) {
	// If we change the maze, we must change "CONST SIZE"
	SliceOfSlices := [][]string{
		//x
		{" ", " ", " ", " ", " ", "#", "#", " ", " ", " "},
		{"#", " ", " ", " ", " ", " ", "#", " ", "#", " "},
		{"#", "#", " ", "#", " ", " ", " ", " ", "#", " "},
		{" ", "#", " ", " ", " ", "#", " ", " ", "#", " "},
		{"#", "#", "#", " ", "#", "#", "#", " ", "#", " "},
		{" ", "#", " ", " ", " ", "#", "#", " ", " ", "#"},
		{"#", "#", "#", "#", " ", "#", "#", " ", "#", " "},
		{"#", " ", " ", " ", " ", " ", " ", " ", "#", "#"},
		{" ", " ", " ", " ", " ", "#", "#", " ", " ", " "},
		{"#", "#", " ", "#", "#", " ", " ", "#", "#", " "},
	}

	BFS(SliceOfSlices, 0, 0, x, y)
}

/*
{" ", " ", " ", " ", " ", "#", "#", " ", " ", " "},
{"#", " ", " ", " ", " ", " ", "#", " ", "#", " "},
{"#", "#", " ", "#", " ", " ", " ", " ", "#", " "},
{" ", "#", " ", " ", " ", "#", " ", " ", "#", " "},
{"#", "#", "#", " ", "#", "#", "#", " ", "#", " "},
{" ", "#", " ", " ", " ", "#", "#", " ", " ", "#"},
{"#", "#", "#", "#", " ", "#", "#", " ", "#", " "},
{"#", " ", " ", " ", " ", " ", " ", " ", "#", "#"},
{" ", " ", " ", " ", " ", "#", "#", " ", " ", " "},
{"#", "#", " ", "#", "#", " ", " ", "#", "#", " "},
*/
