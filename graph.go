package main

import (
	"fmt"
	"strconv"
)

type Vertices struct {
	Data string
}

type Graph struct {
	Matrix [][]int
}

func printGraph(graph *Graph) {
	for i := range graph.Matrix {
		for j := range graph.Matrix[i] {
			fmt.Printf(strconv.Itoa(graph.Matrix[i][j]) + " ")
		}
		fmt.Println()
	}
}

func main() {
	twoD := &Graph{
		Matrix: make([][]int, 3),
	} //row

	for i := range twoD.Matrix {
		twoD.Matrix[i] = make([]int, 4, 4) //col
	}

	twoD.Matrix[1][1] = 3

	printGraph(twoD)

}
