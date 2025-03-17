package controller

import (
	"github.com/ArthurRod/go-todo/internal/model"
	"github.com/ArthurRod/go-todo/internal/view"
)

func AddTask(description string) {
	task := model.NewTask(description)
	view.ShowTaskAdded(task) 
}

func ListTasks() {
	tasks := model.ListTasks()
	view.ShowTaskList(tasks) 
}

func RemoveTask(id int) {
	if success := model.RemoveTask(id); success {
		view.ShowTaskRemoved(id) 
	} else {
		view.ShowTaskNotFound(id) 
	}
}

func MarkTaskAsDone(id int) {
	if success := model.MarkTaskAsDone(id); success {
		view.ShowTaskMarkedAsDone(id) 
	} else {
		view.ShowTaskNotFound(id) 
	}
}
