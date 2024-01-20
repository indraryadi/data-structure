package main

import "fmt"

type Node struct {
	MusicName string
	Next      *Node
}

type List struct {
	Head *Node
}

func (list *List) Insert(musicName string) {
	// step
	// create new list
	// check if current list head is nil, use newList as Head
	// else, create var to store Head called previous
	// loop through list to get the last node,
	// by check if next node is not nil (nil indicate current node is node last node) change value of previousNode into Next value of previous node
	// after got last node, fill that node last value using newNode
	fmt.Println("INSERT")
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

func main() {
	fmt.Println("LINKED LIST")
	list := &List{}

	list.Insert("goose")
	Show(list)

	list.Insert("up into the ether")
	list.Insert("gravity")
	list.Insert("slow dancing in the burning room")
	list.Insert("who you are")
	Show(list)
}
