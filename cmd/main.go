package main

import (
	"fmt"
	"log"

	"github.com/ArthurRod/go-todo/internal/controller"
	"github.com/ArthurRod/go-todo/internal/model"
	"github.com/ArthurRod/go-todo/internal/view"
)

func main() {
	
	err := model.LoadTasks()
	if err != nil {
		log.Fatal("Erro ao carregar tarefas:", err)
	}

	for {
		view.ShowMenu() 
		var option int
		fmt.Print("\nEscolha uma opção: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			var description string
			fmt.Print("\nDescrição da tarefa: ")
			fmt.Scanf("%s\n", &description)
			controller.AddTask(description)
		case 2:
			controller.ListTasks()
		case 3:
			var id int
			fmt.Print("\nDigite o ID da tarefa para remover: ")
			fmt.Scan(&id)
			controller.RemoveTask(id)
		case 4:
			var id int
			fmt.Print("\nDigite o ID da tarefa para marcar como concluída: ")
			fmt.Scan(&id)
			controller.MarkTaskAsDone(id) 
		case 5:
			fmt.Println("\nSaindo...")
			return
		default:
			fmt.Println("Opção inválida.")
		}
	}
}
