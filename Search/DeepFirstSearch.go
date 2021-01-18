package Search

import (
	"fmt"
)

type Vertex struct {
	visited    bool      // Mark if some Vertex was visited or not
	value      string    // The value that the Vertex have
	neighbours []*Vertex // All linked vertexes
}

type Graph struct{}

// Add a new Vertex to the Graph, every vertex has 1 value
func NewVertex(value string) *Vertex {
	return &Vertex{
		value: value,
	}
}

// Make a Link between two or more Vertex
func (v *Vertex) connect(vertex ...*Vertex) {
	v.neighbours = append(v.neighbours, vertex...)
}

// Deep First Search algorithm
func (g *Graph) dfs(vertex *Vertex) {
	// If the vertex has already been visited, just return
	if vertex.visited {
		return
	}

	// If its not, then it has been visited now
	vertex.visited = true
	fmt.Println(vertex.value)

	if len(vertex.neighbours) == 0 {
		return
	}
	// For every connection that some Vertex "holds" we have to go deeper into that connection
	for _, v := range vertex.neighbours {
		g.dfs(v)
	}
}

func (g *Graph) disconnected(vertices ...*Vertex) {
	for _, v := range vertices {
		g.dfs(v)
	}
}

func main() {
	v1 := NewVertex("1")
	v2 := NewVertex("2")
	// v3 := NewVertex("3")
	// v4 := NewVertex("4")
	v5 := NewVertex("5")
	v6 := NewVertex("6")
	v7 := NewVertex("7")
	v8 := NewVertex("8")
	v9 := NewVertex("9")
	v10 := NewVertex("10")

	g := Graph{}

	v1.connect(v9, v5, v2)
	v6.connect(v7)
	v9.connect(v10)
	v5.connect(v8, v6)
	g.dfs(v1)

	/*
	               v1
	           / | \__ v2
	           v9   v5
	           |    | \
	       v10  v8  v6
	                   \
	                   v7


	   * 1 -> 'g.dfs(v1) loop' -> print(1) -> g.dfs(v9)..(v5)...(v2)
	       * 'g.dfs(v9) loop' -> print(9) -> g.dfs(v10) -> return
	           *'g.dfs(v10) loop' -> print(10) -> return
	       * 'g.dfs(v5) loop' -> print(5) -> g.dfs(v8)..(v6) -> return
	           *'g.dfs(v8) loop' -> print(8) -> return
	           *'g.dfs(v6) loop' -> print(6) -> return
	       * 'g.dfs(v2) loop' -> print(2) -> end all vertex have been visited
	*/
}