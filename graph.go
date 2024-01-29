package main

import (
	"fmt"
	"strconv"
)

type Vertices struct {
	Vertices []string
}

type Graph struct {
	Matrix [][]int
}

func printGraph(vertices *Vertices, graph *Graph) {
	fmt.Printf(" ")
	for x := range vertices.Vertices {
		fmt.Printf(" %s", vertices.Vertices[x])
	}
	fmt.Println()
	for i := range graph.Matrix {
		fmt.Printf("%s ", vertices.Vertices[i])
		for j := range graph.Matrix[i] {
			fmt.Printf(strconv.Itoa(graph.Matrix[i][j]) + " ")
		}
		fmt.Println()
	}
}

func addVertices(vertices *Vertices, data string) *Vertices {
	vertices.Vertices = append(vertices.Vertices, data)
	return &Vertices{
		Vertices: vertices.Vertices,
	}
}

func addEdge(graph *Graph, src int, dst int) {
	status := checkEdge(graph, src, dst)
	if status {
		fmt.Println("cannot add edge")
	} else {
		graph.Matrix[src][dst] = 1
		fmt.Println("done")
	}
}

func checkEdge(graph *Graph, src int, dst int) bool {
	if graph.Matrix[src][dst] == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	row := 4
	col := 4

	sliceOfVertices := &Vertices{}

	addVertices(sliceOfVertices, "A") //0
	addVertices(sliceOfVertices, "B") //1
	addVertices(sliceOfVertices, "C") //2
	addVertices(sliceOfVertices, "D") //3
	fmt.Println(sliceOfVertices)

	twoD := &Graph{
		Matrix: make([][]int, row),
	} //row

	for i := range twoD.Matrix {
		twoD.Matrix[i] = make([]int, col) //col
	}

	fmt.Println(sliceOfVertices.Vertices[1])
	addEdge(twoD, 0, 1)
	addEdge(twoD, 0, 1)
	addEdge(twoD, 1, 0)
	addEdge(twoD, 0, 2)
	addEdge(twoD, 0, 3)
	addEdge(twoD, 2, 3)
	printGraph(sliceOfVertices, twoD)

}
