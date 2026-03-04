package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Viperesss/taskmanager/task"
	"github.com/Viperesss/taskmanager/utils"
)

func ReadCli() error {
	lineCounter := len(task.TaskList) + 1
	fmt.Println("Введите: Название заметки; теги(через запятую); приоритет от 1 до 5")
	input, err := utils.StringReader()
	if err != nil {
		return err
	}

	parts := strings.Split(input, ";")
	noteName := parts[0]
	noteTags := strings.Split(parts[1], ",")
	notePriorites := strings.Fields(parts[2])
	notePriority, err := strconv.Atoi(notePriorites[0])
	if err != nil {
		fmt.Println("Ошибка при вводе уровня приоритета")
		return err
	}

	var newTask task.Task
	err = newTask.AddTask(noteName, noteTags, notePriority)
	if err != nil {
		return fmt.Errorf("Ошибка при создании новой заметки, строка: %v\nОшибка: %v", lineCounter, err)
	}

	return nil
}

func Search() error {
	fmt.Print("Введите тег: ")
	targetTag, err := utils.StringReader()
	if err != nil {
		return fmt.Errorf("Ошибка при вводе тега:\n%v\nОшибка: %v", targetTag, err)
	}

	task.Search(targetTag)
	return nil
}
