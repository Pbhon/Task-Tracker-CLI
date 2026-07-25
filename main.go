package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SaveData struct {
	NextID int           `json:"next_id"`
	Tasks  map[int]*Task `json:"tasks"`
}

func main() {
	fmt.Println(
		"This is a Task Tracker CLI. \n" +
			"This is the list of functions available:\n" +
			"add - Creates a new task\n" +
			"update - Updates a task description using its id number\n" +
			"mark-in-progress - Sets a task status as in progress\n" +
			"mark-completed - Sets a task status as done\n" +
			"delete - Deletes a task using its id number\n" +
			"list - Lists out all available tasks along with their status\n" +
			"list-todo - Lists all tasks with status 'To-Do' \n" +
			"list-in-progress - Lists all tasks with status 'In Progress' \n" +
			"list-completed - Lists all tasks with status 'Completed'\n" +
			"finish - End task changing",
	)

	saveData := SaveData{}
	readJSON(&saveData)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var action string

		fmt.Println("Enter desired action: ")

		scanner.Scan()
		action = scanner.Text()

		if !IsValidAction(action) {
			fmt.Println("Invalid action, try again: ")
		} else {
			switch action {
			case "add":
				var description string

				fmt.Print("Enter a description for the task: ")
				scanner.Scan()
				description = scanner.Text()
				if err := scanner.Err(); err != nil {
					fmt.Println("Error reading input:", err)
					break
				}

				AddTask(description, &saveData)

			case "update":
				var idString string
				var id int
				var description string

				var err error

				fmt.Println("Enter task ID: ")
				if scanner.Scan() {
					idString = scanner.Text()

					id, err = strconv.Atoi(idString)

					if err != nil {
						fmt.Println("Error task ID: ", err)
						break
					}
				}

				fmt.Println("Enter task description: ")
				if scanner.Scan() {
					description = scanner.Text()
				}

				UpdateTask(id, description, &saveData)

			case "delete":
				var idString string
				var id int

				var err error
				fmt.Println("Enter task ID: ")

				if scanner.Scan() {
					idString = scanner.Text()

					id, err = strconv.Atoi(idString)
				}

				if err != nil {
					fmt.Println("Error task ID: ", err)
					break
				}

				DeleteTask(id, &saveData)

			case "mark-in-progress":
				var idString string
				var id int

				var err error

				fmt.Println("Enter task ID: ")

				if scanner.Scan() {
					idString = scanner.Text()

					id, err = strconv.Atoi(idString)
					if err != nil {
						fmt.Println("Error task ID: ", err)
						break
					}
				}

				MarkInProgress(id, &saveData)

			case "mark-completed":
				var idString string
				var id int

				var err error

				fmt.Println("Enter task ID: ")

				if scanner.Scan() {
					idString = scanner.Text()

					id, err = strconv.Atoi(idString)
					if err != nil {
						fmt.Println("Error task ID:", err)
						break
					}
				}

				MarkCompleted(id, &saveData)

			case "list":
				List(saveData)

			case "list-todo":
				ListTodo(saveData)

			case "list-in-progress":
				ListInProgress(saveData)

			case "list-completed":
				ListCompleted(saveData)

			case "finish":
				fmt.Println("Closing Program")
				return
			}

		}
	}

}
