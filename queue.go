package main

// if the max len of queue is 5 we only can add untill 4 data, it used to check the queue is full or not
// empty queueu head and tail has same position
// full queue tail is one position behind front pointer

import "fmt"

func checkFull(headIndex *int, tailIndex *int, queue *[]interface{}) bool {
	if *headIndex-*tailIndex == (len(*queue)-1)*-1 || *headIndex-*tailIndex == 1 {
		return true
	}
	return false
}
func enqueue(headIndex *int, tailIndex *int, queue *[]interface{}, data interface{}) {
	fmt.Println("ENQUEUE: ", data)
	isFull := checkFull(headIndex, tailIndex, queue)
	if !isFull {
		(*queue)[*tailIndex] = data
		*tailIndex += 1
		if *tailIndex > len(*queue)-1 {
			*tailIndex = 0
		}
	} else {
		fmt.Println("QUEUE FULL")
	}
}
func dequeue(headIndex *int, tailIndex *int, queue *[]interface{}) {
	fmt.Println("DEQUEUE", *headIndex)
	(*queue)[*headIndex] = nil
	*headIndex += 1
}
func main() {
	// implementation of queue using array (fix length queue)
	var headIndex int // this used for dequeue, when want to dequeue use this index to indicate which index need to dequeue
	var tailIndex int // this used for enqueue, when want to enqueue use this index to indicate which index need to enqueue
	queue := make([]interface{}, 5)
	fmt.Println(queue)

	enqueue(&headIndex, &tailIndex, &queue, "print a")
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	enqueue(&headIndex, &tailIndex, &queue, "print b")
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	enqueue(&headIndex, &tailIndex, &queue, "print c")
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	enqueue(&headIndex, &tailIndex, &queue, "print d")
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	enqueue(&headIndex, &tailIndex, &queue, "print e")
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	dequeue(&headIndex, &tailIndex, &queue)
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	enqueue(&headIndex, &tailIndex, &queue, "print e")
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	enqueue(&headIndex, &tailIndex, &queue, "print f")
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	dequeue(&headIndex, &tailIndex, &queue)
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)

	enqueue(&headIndex, &tailIndex, &queue, "print f")
	fmt.Printf("head:%d | tail:%d\n", headIndex, tailIndex)
	fmt.Println(queue)
}
