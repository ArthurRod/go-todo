package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var Tasks []Task
var NextID int

func LoadTasks() error {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			
			return nil
		}
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Tasks)
	if err != nil {
		return fmt.Errorf("erro ao decodificar arquivo: %w", err)
	}

	
	if len(Tasks) > 0 {
		NextID = Tasks[len(Tasks)-1].ID + 1
	}
	return nil
}

func SaveTasks() error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") 
	err = encoder.Encode(Tasks)
	if err != nil {
		return fmt.Errorf("erro ao codificar arquivo: %w", err)
	}

	return nil
}

func NewTask(description string) Task {
	task := Task{
		ID:          NextID,
		Description: description,
		Done:        false,
	}
	NextID++
	Tasks = append(Tasks, task)
	SaveTasks() 
	return task
}

func ListTasks() []Task {
	return Tasks
}

func RemoveTask(id int) bool {
	for index, task := range Tasks {
		if task.ID == id {
			Tasks = append(Tasks[:index], Tasks[index+1:]...)
			SaveTasks() 
			return true
		}
	}
	return false
}

func MarkTaskAsDone(id int) bool {
	for i, task := range Tasks {
		if task.ID == id {
			Tasks[i].Done = true
			SaveTasks() 
			return true
		}
	}
	return false
}
