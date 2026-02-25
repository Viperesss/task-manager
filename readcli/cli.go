package cli

import (
	"fmt"
	"log"

	file "github.com/Viperesss/taskmanager/readfile"
	"github.com/Viperesss/taskmanager/task"
	"github.com/Viperesss/taskmanager/utils"
)

// Введите название, теги, приоритет
func Add() {
	fmt.Println("Введите: Название заметки; теги; приоритет от 1 до 5")

}

func Search() []task.Task {
	var Tasks []task.Task

	fmt.Print("Введите тег: ")
	input, err := utils.StringReader()
	if err != nil {
		log.Fatal(err)
	}

	tasks := file.ReadFile()
	for _, task := range tasks {
		for _, tags := range task.Meta.Tags {
			if tags == input {
				Tasks = append(Tasks, task)
			}
		}
	}
	return Tasks
}
