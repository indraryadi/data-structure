package main

import "fmt"

type Node struct {
	MusicName string
	Next      *Node
}

type List struct {
	Head   *Node
	Length int
}

// add data into last list
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
		list.Length += 1
		fmt.Println(list.Length)
	} else {
		previousNode := list.Head
		for previousNode.Next != nil {
			previousNode = previousNode.Next
		}
		previousNode.Next = newNode
		list.Length += 1
		fmt.Println(list.Length)
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
	fmt.Println(list.Length)
}

// add data into first list
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
	list.Length += 1
	fmt.Println(list.Length)
}

// insert data at any index
func (list *List) Insert(index int, musicName string) {
	fmt.Println("INSERT")
	if index > list.Length {
		fmt.Println("out of index")
	} else if index == 0 {
		fmt.Println("add into first")
		list.Prepend(musicName)
		fmt.Println("done insert")
	} else if index == list.Length {
		fmt.Println("add into last")
		list.Append(musicName)
		fmt.Println("done insert")
	} else {
		fmt.Println("add not into middle")
		newNode := &Node{
			MusicName: musicName,
			Next:      nil,
		}
		node := list.Head
		var count int
		for node.Next != nil {
			if count == index-1 {
				newNode.Next = node.Next
				node.Next = newNode
				list.Length += 1
				break
			}
			node = node.Next
			count += 1
		}
	}
}

// delete
func (list *List) Delete(index int) {
	fmt.Println("DELETE")
	if index > list.Length {
		fmt.Println("out of index")
	} else if index == 0 {
		fmt.Println("delete first")
		head := list.Head
		list.Head = head.Next
		head.Next = nil
		list.Length -= 1
		fmt.Println("done delete")
	} else if index == list.Length {
		fmt.Println("delete last")
		prev := list.Head
		for prev.Next.Next != nil {
			fmt.Println(prev.MusicName)
			prev = prev.Next
		}
		fmt.Println(prev.MusicName)
		prev.Next = nil
		list.Length -= 1
		fmt.Println("done delete")
	} else {
		fmt.Println("delete from middle")
	}
}

// update

func main() {
	fmt.Println("LINKED LIST")
	list := &List{}

	list.Append("goose")
	list.Append("up into the ether")
	list.Append("gravity")
	list.Append("slow dancing in the burning room")
	list.Prepend("pyscosocial")
	list.Append("who you are")
	Show(list)
	fmt.Println("====")

	list.Insert(1, "NEW MUSIC")
	Show(list)

	list.Delete(7)
	Show(list)
}
