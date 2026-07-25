package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

func readJSON(saveData *SaveData) {
	jsonFile, err := os.OpenFile("tasks.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return
	}

	defer func() {
		if err := jsonFile.Close(); err != nil {
			log.Println("failed to close file:", err)
		}
	}()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON: ", err)
	}

	if len(data) == 0 {
		saveData.Tasks = make(map[int]*Task)
	}

	err = json.Unmarshal(data, saveData)

	if saveData.Tasks == nil {
		saveData.Tasks = make(map[int]*Task)
	}
}

func writeJSON(saveData *SaveData) {
	jsonBytes, err := json.MarshalIndent(saveData, "", "    ")
	if err != nil {
		fmt.Println("Error converting to json", err)
		return
	}

	err = os.WriteFile("tasks.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println("Error writing to JSON", err)
		return
	}
}

func IsValidAction(action string) bool {
	actions := []string{
		"add",
		"update",
		"delete",
		"mark-in-progress",
		"mark-completed",
		"list",
		"list-todo",
		"list-in-progress",
		"list-completed",
		"finish",
	}

	if slices.Contains(actions, action) {
		return true
	}

	return false
}

func CompletedAction() {
	fmt.Println("Action Completed!")
	fmt.Println("")
}
