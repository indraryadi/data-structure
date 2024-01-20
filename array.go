package main

import (
	"fmt"
	"strconv"
)

func show(tasks []string) {
	for i, v := range tasks {
		fmt.Printf("%d.%s\n", i+1, v)
	}
}

func add(tasks []string) []string {
	var newTask string
	fmt.Printf("add task:")
	fmt.Scanln(&newTask)
	tasks = append(tasks, newTask) //it will not work bc using append will will create new slice not modify the original value of slice
	fmt.Println(newTask + " added to tasks")
	return tasks
}
func edit(tasks []string) []string {
	show(tasks)
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

	var updateTask string
	fmt.Println("input update task:")
	fmt.Scanln(&updateTask)

	tasks[indexConv] = updateTask
	return tasks
}

func delete(tasks []string) []string {
	show(tasks)
	var index string
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

func main() {
	// task list
	// functionality:
	// 1.add task
	// 2.edit task
	// 3.delete task
	// 4.show all task

	var tasks []string
	tasks = append(tasks, "good")
	tasks = append(tasks, "notgood")
	tasks = append(tasks, "ggwp")
	tasks = append(tasks, "well")
	tasks = append(tasks, "enough")

	// fmt.Printf("tasks:\n1.Show tasks\n2.Add task\n3.Edit task\n4.Delete task\n")

	show(tasks)

	tasks = add(tasks)
	show(tasks)

	tasks = edit(tasks)
	show(tasks)

	tasks = delete(tasks)
	show(tasks)

}
