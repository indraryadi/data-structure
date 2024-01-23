package main

import (
	"fmt"
	"strings"
)

func pop(undostack *[]string, redostack *[]string) {
	fmt.Println("POP")
	// fmt.Println((*undostack)[len(*undostack)-1])
	fmt.Println((*undostack))

	push(redostack, (*undostack)[len(*undostack)-1])
	*undostack = (*undostack)[:len(*undostack)-1]
}

func push(stack *[]string, data string) {
	*stack = append((*stack), data)
}

func redo(undostack *[]string, redostack *[]string) {
	fmt.Println("REDO")
	fmt.Println((*undostack))
	// fmt.Println((*redostack)[len(*redostack)-1])

	push(undostack, (*redostack)[len(*redostack)-1])
	*redostack = (*redostack)[:len(*redostack)-1]
}

func splitting(data string) (string, string) {
	dat := strings.Split(data, " ")
	var out string
	for _, v := range dat[1:] {
		out += v + " "
	}
	out = strings.Trim(out, " ")
	return dat[0], out
}

func main() {
	undoStack := []string{}
	redoStack := []string{}
	undoStack = append(undoStack, "hello")
	undoStack = append(undoStack, "world")
	undoStack = append(undoStack, "test")

	input := []string{"write A", "write B C D", "write testing input", "undo", "redo", "write hello world"}

	for _, v := range input {
		command, out := splitting(v)
		if command == "write" {
			undoStack = append(undoStack, out)
		} else if command == "undo" {
			pop(&undoStack, &redoStack)
		} else if command == "redo" {
			redo(&undoStack, &redoStack)
		}
	}

	fmt.Println("undo stack: ", undoStack)
	fmt.Println("redo stack", redoStack)
}
