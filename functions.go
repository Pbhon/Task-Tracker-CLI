package main

import (
	"fmt"
	"strconv"
	"time"
)

func AddTask(description string, saveData *SaveData) {
	task := Task{
		Id:          saveData.NextID,
		Description: description,
		Status:      "To-do",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	saveData.Tasks[saveData.NextID] = &task
	saveData.NextID++

	writeJSON(saveData)

	CompletedAction()
}

func UpdateTask(id int, description string, saveData *SaveData) {
	saveData.Tasks[id].Description = description
	saveData.Tasks[id].UpdatedAt = time.Now()

	writeJSON(saveData)

	CompletedAction()
}

func DeleteTask(id int, saveData *SaveData) {
	delete(saveData.Tasks, id)

	writeJSON(saveData)

	CompletedAction()
}

func MarkInProgress(id int, saveData *SaveData) {
	saveData.Tasks[id].Status = "In Progress"
	saveData.Tasks[id].UpdatedAt = time.Now()

	writeJSON(saveData)

	CompletedAction()
}

func MarkCompleted(id int, saveData *SaveData) {
	saveData.Tasks[id].Status = "Completed"
	saveData.Tasks[id].UpdatedAt = time.Now()

	writeJSON(saveData)

	CompletedAction()
}

func List(saveData SaveData) {
	for _, task := range saveData.Tasks {
		var idString string
		idString = strconv.Itoa(task.Id)
		fmt.Println("ID: " + idString + " | " + task.Description + " | " + task.Status)
	}

	CompletedAction()
}

func ListTodo(saveData SaveData) {
	for _, task := range saveData.Tasks {
		if task.Status == "To-do" {
			var idString string
			idString = strconv.Itoa(task.Id)
			fmt.Println("ID" + idString + " | " + task.Description + " | " + task.Status)
		}
	}

	CompletedAction()
}

func ListInProgress(saveData SaveData) {
	for _, task := range saveData.Tasks {
		if task.Status == "In Progress" {
			var idString string
			idString = strconv.Itoa(task.Id)
			fmt.Println("ID: " + idString + " | " + task.Description + " | " + task.Status)
		}
	}

	CompletedAction()
}

func ListCompleted(saveData SaveData) {
	for _, task := range saveData.Tasks {
		if task.Status == "Completed" {
			var idString string
			idString = strconv.Itoa(task.Id)
			fmt.Println("ID: " + idString + " | " + task.Description + " | " + task.Status)
		}
	}

	CompletedAction()
}
