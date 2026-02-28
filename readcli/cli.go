package cli

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Viperesss/taskmanager/task"
	"github.com/Viperesss/taskmanager/utils"
)

func Add(tList *[]task.Task) {
	counter := 1
	fmt.Println("Введите: Название заметки; теги(через запятую); приоритет от 1 до 5")
	input, err := utils.StringReader()
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(input, ";")
	noteName := parts[0]
	noteTags := strings.Split(parts[1], ",")
	notePriority, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Printf("Ошибка при вводе уровня приоритета, %v", err)
		return
	}

	newTask := task.Task{
		ID:   counter,
		Name: noteName,
		Meta: task.Meta{
			Tags:     noteTags,
			Priority: notePriority,
		},
	}

	*tList = append(*tList, newTask)
	fmt.Println("Заметка успешно добавлена")
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
