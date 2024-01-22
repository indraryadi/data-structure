package main

import "fmt"

type Node struct {
	MusicName string
	Next      *Node
}

type List struct {
	Head *Node
}

func (list *List) Append(musicName string) {
	// step
	// create new list
	// check if current list head is nil, use newList as Head
	// else, create var to store Head called previous
	// loop through list to get the last node,
	// by check if next node is not nil (nil indicate current node is node last node) change value of previousNode into Next value of previous node
	// after got last node, fill that node last value using newNode
	fmt.Println("APPEND")
	newNode := &Node{
		MusicName: musicName,
		Next:      nil,
	}

	if list.Head == nil {
		list.Head = newNode
	} else {
		previousNode := list.Head
		for previousNode.Next != nil {
			previousNode = previousNode.Next
		}
		previousNode.Next = newNode
	}
}

func Show(list *List) {
	fmt.Println("SHOW")
	previous := list.Head
	if previous.Next == nil {
		fmt.Printf("%v ->\n", previous.MusicName)
	} else {
		for previous.Next != nil {
			fmt.Printf("%v ->", *&previous.MusicName)
			previous = previous.Next
		}
		fmt.Printf("%v ->\n", previous.MusicName)
	}
}

func (list *List) Prepend(musicName string) {
	fmt.Println("PREPEND")
	newNode := &Node{
		MusicName: musicName,
		Next:      nil,
	}

	if list.Head == nil {
		list.Head = newNode
	} else {
		previousNode := list.Head
		list.Head = newNode
		list.Head.Next = previousNode
	}
}
func main() {
	fmt.Println("LINKED LIST")
	list := &List{}

	list.Append("goose")
	Show(list)

	list.Append("up into the ether")
	list.Append("gravity")
	list.Append("slow dancing in the burning room")
	list.Append("who you are")
	Show(list)

	list.Prepend("pyscosocial")
	Show(list)
}
