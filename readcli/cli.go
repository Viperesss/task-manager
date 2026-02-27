package cli

import (
	"fmt"
	"log"

	"github.com/Viperesss/taskmanager/task"
	"github.com/Viperesss/taskmanager/utils"
)

// Введите название, теги, приоритет
func Add() {
	fmt.Println("Введите: Название заметки; теги; приоритет от 1 до 5")

}

func Search(tList []task.Task) []task.Task {
	var Tasks []task.Task

	fmt.Print("Введите тег: ")
	input, err := utils.StringReader()
	if err != nil {
		log.Fatal(err)
	}

	for _, task := range tList {
		for _, tags := range task.Meta.Tags {
			if tags == input {
				Tasks = append(Tasks, task)
			}
		}
	}
	return Tasks
}
