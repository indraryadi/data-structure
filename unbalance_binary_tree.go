package main

import (
	"fmt"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func insertData(trees *TreeNode, newNode *TreeNode) {
	fmt.Printf("compare %d with %d\n", trees.Data, newNode.Data)
	if newNode.Data >= trees.Data {
		if trees.Right == nil {
			// fmt.Println("INSERT TO RIGHT")
			trees.Right = newNode
			// fmt.Printf("%d => %d\n", trees.Data, trees.Right.Data)
			// fmt.Println("DONE")
		} else {
			// fmt.Println("recursive")
			insertData(trees.Right, newNode)
		}
	} else {
		if trees.Left == nil {
			// fmt.Println("INSERT TO LEFT")
			trees.Left = newNode
			// fmt.Printf("%d => %d\n", trees.Data, trees.Left.Data)
			// fmt.Println("DONE")
		} else {
			// fmt.Println("recursive")
			insertData(trees.Left, newNode)
		}
	}
}

func printTrees(trees *TreeNode) {
	if trees.Left != nil {
		printTrees(trees.Left)
	}

	fmt.Printf("%d =>", trees.Data)

	if trees.Right != nil {
		printTrees(trees.Right)
	}
}

func findData(trees *TreeNode, data int) string {
	var res string
	if trees == nil {
		res = "not found"
		return res
	}

	if trees.Data == data {
		res = "found"
		return res
	} else if data < trees.Data {
		return findData(trees.Left, data)
	} else {
		return findData(trees.Right, data)
	}
}

func main() {
	trees := &TreeNode{Data: 10, Left: nil, Right: nil}

	node2 := &TreeNode{Data: 5, Left: nil, Right: nil}
	node3 := &TreeNode{Data: 15, Left: nil, Right: nil}
	node4 := &TreeNode{Data: 8, Left: nil, Right: nil}
	insertData(trees, node2)
	insertData(trees, node3)
	insertData(trees, node4)

	printTrees(trees)
	result := findData(trees, 8)
	fmt.Println(result)
}
