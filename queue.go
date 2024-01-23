package main

import "fmt"

func enqueue(headIndex *int, tailIndex *int, queue *[]interface{}, data interface{}) {
	(*queue)[*tailIndex] = data
	*tailIndex += 1
}
func dequeue(headIndex *int, tailIndex *int, queue *[]interface{}) {
	(*queue)[*headIndex] = nil
	*headIndex += 1
}
func main() {
	// implementation queue using array
	var headIndex int
	var tailIndex int
	queue := make([]interface{}, 5)
	fmt.Println(queue)

	enqueue(&headIndex, &tailIndex, &queue, "print a")
	enqueue(&headIndex, &tailIndex, &queue, "print b")
	enqueue(&headIndex, &tailIndex, &queue, "print c")

	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	dequeue(&headIndex, &tailIndex, &queue)
	dequeue(&headIndex, &tailIndex, &queue)
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	enqueue(&headIndex, &tailIndex, &queue, "print d")
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)
}
