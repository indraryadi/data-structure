package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func show(tasks []string) {
	fmt.Println("SHOW TASK")
	for i, v := range tasks {
		fmt.Printf("%d.%s\n", i+1, v)
	}
}

func add(tasks []string) []string {
	fmt.Println("ADD TASK")
	// get user input
	var newTask string
	fmt.Printf("task:")
	userInput := bufio.NewScanner(os.Stdin)
	if !userInput.Scan() {
		fmt.Println("error reading input:", userInput.Err())
	}
	newTask = userInput.Text()

	// add user input into slice
	tasks = append(tasks, newTask)
	fmt.Println(newTask + " added to tasks")
	return tasks
}
func edit(tasks []string) []string {
	fmt.Println("EDIT TASK")
	show(tasks)
	// select index that want to edit
	var index string
	fmt.Println("select no of task:")
	fmt.Scanln(&index)
	indexConv, err := strconv.Atoi(index)
	indexConv -= 1
	if err != nil {
		panic(err)
	}
	if indexConv > len(tasks)-1 {
		fmt.Println("wrong number")
	}

	// update value of selected index
	var updateTask string
	fmt.Println("input update task:")
	userInput := bufio.NewScanner(os.Stdin)
	if !userInput.Scan() {
		fmt.Println("error reading input:", userInput.Err())
	}
	updateTask = userInput.Text()
	tasks[indexConv] = updateTask
	return tasks
}

func delete(tasks []string) []string {
	fmt.Println("DELETE TASK")
	show(tasks)
	var index string
	// select index of task that want to delete
	fmt.Println("select no of task:")
	fmt.Scanln(&index)
	deletedIndex, err := strconv.Atoi(index)
	deletedIndex -= 1
	if err != nil {
		panic(err)
	}
	if deletedIndex > len(tasks)-1 {
		fmt.Println("wrong number")
	}

	// fast way, result order not same like before edited
	// tasks[deletedIndex] = tasks[len(tasks)-1] //copy last element to index of deleted item
	// tasks[len(tasks)-1] = ""                  //empty the last element
	// tasks = tasks[:len(tasks)-1]              //truncate the slice, it is like slicing array on python

	// slower way, result order same like before deleted
	copy(tasks[deletedIndex:], tasks[deletedIndex+1:]) // shift all data from deletedIndex+1 until the end to the left one index ex: (1,2,3,4) index 2 is deletedIndex, result: (1,3,4,4)
	tasks[len(tasks)-1] = ""
	tasks = tasks[:len(tasks)-1]
	return tasks
}

func clearScreen(name string) {
	cmd := exec.Command(name)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
	var tasks []string

	var menu int

	for loop := true; loop; {
		fmt.Printf("menu:\n1.Show tasks\n2.Add task\n3.Edit task\n4.Delete task\n5.Exit\n")
		fmt.Println("select menu (1-4):")
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			clearScreen("clear")
			show(tasks)
		case 2:
			clearScreen("clear")
			tasks = add(tasks)
		case 3:
			clearScreen("clear")
			tasks = edit(tasks)
		case 4:
			clearScreen("clear")
			tasks = delete(tasks)
		default:
			loop = false
			break
		}
	}

}
