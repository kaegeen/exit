package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var todos []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Simple To-Do List")
	fmt.Println("=================")
	fmt.Println("Commands:")
	fmt.Println(" - ADD <task>: Add a new task")
	fmt.Println(" - LIST: Show all tasks")
	fmt.Println(" - DONE <task number>: Mark a task as done")
	fmt.Println(" - EXIT: Exit the program")
	fmt.Println("=================")

	for {
		fmt.Print("Enter a command: ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)
		command := strings.ToUpper(parts[0])

		switch command {
		case "ADD":
			if len(parts) < 2 {
				fmt.Println("Error: Task description cannot be empty.")
			} else {
				todos = append(todos, parts[1])
				fmt.Println("Task added!")
			}
		case "LIST":
			if len(todos) == 0 {
				fmt.Println("No tasks found.")
			} else {
				fmt.Println("To-Do List:")
				for i, task := range todos {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case "DONE":
			if len(parts) < 2 {
				fmt.Println("Error: Please specify a task number.")
				continue
			}
			var taskNumber int
			_, err := fmt.Sscanf(parts[1], "%d", &taskNumber)
			if err != nil || taskNumber < 1 || taskNumber > len(todos) {
				fmt.Println("Error: Invalid task number.")
			} else {
				todos = append(todos[:taskNumber-1], todos[taskNumber:]...)
				fmt.Println("Task completed!")
			}
		case "EXIT":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}
}
