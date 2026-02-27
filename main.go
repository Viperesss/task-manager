package main

import (
	"fmt"
	"strconv"

	cli "github.com/Viperesss/taskmanager/readcli"
	file "github.com/Viperesss/taskmanager/readfile"
	"github.com/Viperesss/taskmanager/task"
	"github.com/Viperesss/taskmanager/utils"
)

func main() {
	var TaskList []task.Task
	for {
		fmt.Printf("\nВыберите опцию:\n")
		fmt.Printf("1. Добавить заметку с помощью консоли\n")
		fmt.Printf("2. Прочитать заметку через файл\n")
		fmt.Print("--------------------------------\n")
		fmt.Printf("3. Посмотреть заметки\n")
		fmt.Printf("4. Поиск заметки по тегу\n")
		fmt.Print("--------------------------------\n")
		fmt.Printf("0. Выход\n")

		input, err := utils.StringReader()
		if err != nil {
			fmt.Println(err)
			continue
		}
		chose, err := strconv.Atoi(input)

		switch chose {
		case 1:
			cli.Add()
		case 2:
			file.ReadFile(&TaskList)
			fmt.Println("Заметки успешно добавлены")
		case 3:
			if len(TaskList) != 0 {
				for _, task := range TaskList {
					fmt.Printf("\n%v. %s, теги: %v, уровень приоритета - %v\n", task.ID, task.Name, task.Meta.Tags, task.Meta.Priority)
				}
			} else {
				fmt.Println("Заметки отсутствуют")
			}

		case 4: // Поиск заметки по тегу
			tasks := cli.Search(TaskList)
			for _, task := range tasks {
				fmt.Printf("%v. %s, теги: %v, уровень приоритета - %v\n", task.ID, task.Name, task.Meta.Tags, task.Meta.Priority)
			}
		case 0:
			return
		default:
			fmt.Println("Ошибка, введите цифру из предложенного списка")
			continue
		}
	}

}
