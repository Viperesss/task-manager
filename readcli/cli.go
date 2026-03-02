package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Viperesss/taskmanager/task"
	"github.com/Viperesss/taskmanager/utils"
)

func Add(tList *[]task.Task) error {
	counter := len(*tList) + 1
	fmt.Println("Введите: Название заметки; теги(через запятую); приоритет от 1 до 5")
	input, err := utils.StringReader()
	if err != nil {
		return err
	}

	parts := strings.Split(input, ";")
	noteName := parts[0]
	noteTags := strings.Split(parts[1], ",")
	notePriority, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Ошибка при вводе уровня приоритета")
		return err
	}

	priority := task.Priority(notePriority)
	if priority.IsCorrect() {
		newTask := task.Task{
			ID:   counter,
			Name: noteName,
			Meta: task.Meta{
				Tags:     noteTags,
				Priority: priority,
			},
		}

		*tList = append(*tList, newTask)
		fmt.Println("Заметка успешно добавлена")
	} else {
		return fmt.Errorf("Ошибка при получении уровня приоритета")
	}
	return nil
}

func Search(tList []task.Task) ([]task.Task, error) {
	var Tasks []task.Task

	fmt.Print("Введите тег: ")
	input, err := utils.StringReader()
	if err != nil {
		return nil, err
	}

	for _, task := range tList {
		for _, tags := range task.Meta.Tags {
			if tags == input {
				Tasks = append(Tasks, task)
			}
		}
	}
	return Tasks, nil
}
