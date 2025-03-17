package view

import (
	"fmt"

	"github.com/ArthurRod/go-todo/internal/model"
)

func ShowTaskList(tasks []model.Task) {
	if len(tasks) == 0 {
		fmt.Println("Não há tarefas no momento.")
		return
	}
	fmt.Println("\nLista de tarefas:")
	for _, task := range tasks {
		status := "não concluída"
		if task.Done {
			status = "concluída"
		}
		fmt.Printf("ID: %d | Descrição: %s | Status: %s\n", task.ID, task.Description, status)
	}
}

func ShowTaskAdded(task model.Task) {
	fmt.Printf("Tarefa '%s' adicionada com sucesso! (ID: %d)\n", task.Description, task.ID)
}

func ShowTaskRemoved(id int) {
	fmt.Printf("Tarefa com ID %d removida com sucesso.\n", id)
}

func ShowTaskNotFound(id int) {
	fmt.Printf("Tarefa com ID %d não encontrada.\n", id)
}

func ShowTaskMarkedAsDone(id int) {
	fmt.Printf("Tarefa com ID %d marcada como concluída!\n", id)
}