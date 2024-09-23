package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var taskMap = make(map[int]string)
var endFlag = false
func main() {

	for !endFlag {
			
		// print menu here
		
		fmt.Println("\n\n1) list\n2) create\n3) update\n4) delete\n5) exit")



		reader := bufio.NewReader(os.Stdin)
		userInstruction, _ := reader.ReadString('\n')
		userInstruction = userInstruction[:len(userInstruction)-1]



		switch userInstruction {
		case "1":
			list()
		
		case "2":
			create()
		
		case "3":
			update()

		case "4":
			deleteTask()
		
		case "5":
			endFlag = true
			return

		default:
			fmt.Println("Invalid choice")
			return
		}
	}
}



// list
func list() {
	if len(taskMap) == 0 {
		fmt.Println("You do not have any tasks!")
		return
	}
	
	fmt.Println("\n\nListing tasks!")
	for taskToList, task := range taskMap {
		fmt.Println(strconv.Itoa(taskToList) + ": " + task)
	}
}



// create
func create() {
	fmt.Println("Enter a task: ")
	reader := bufio.NewReader(os.Stdin)
	newTask, _ := reader.ReadString('\n')
	newTask = newTask[:len(newTask)-1]
	

	var newTaskNumber = len(taskMap)
	taskMap[newTaskNumber] = newTask
	fmt.Println("\n\nTask Added")
	list()
}



// update
func update() {
	if len(taskMap) == 0 {
		fmt.Println("You do not have any tasks!")
		return
	}
	
	fmt.Println("Tell us which task to update: ")
	reader := bufio.NewReader(os.Stdin)
	taskToUpdate, _ := reader.ReadString('\n')
	taskToUpdate = taskToUpdate[:len(taskToUpdate)-1]
	taskToUpdateInt, _ := strconv.Atoi(taskToUpdate)

	fmt.Println("What to change it to: ")
	taskDescription, _ := reader.ReadString('\n')
	taskDescription = taskDescription[:len(taskDescription)-1]
	
	taskMap[taskToUpdateInt] = taskDescription
	fmt.Println("\n\n Task Updated")
	list()

}



// delete
func deleteTask() {

	fmt.Println("Tell us which task to delete: ")
	reader := bufio.NewReader(os.Stdin)
	taskToDelete, _ := reader.ReadString('\n')
	taskToDelete = taskToDelete[:len(taskToDelete)-1]
	taskToDeleteInt, _ := strconv.Atoi(taskToDelete)

	delete(taskMap, taskToDeleteInt)
	fmt.Println("Deleted task")
	list()
}
