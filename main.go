package main

import (
	"fmt"
	"strconv"

	cli "github.com/Viperesss/taskmanager/readcli"
	file "github.com/Viperesss/taskmanager/readfile"
	"github.com/Viperesss/taskmanager/utils"
)

func main() {
	for {
		fmt.Printf("\nВыберите опцию:\n")
		fmt.Printf("1. Добавить заметку с помощью консоли\n")
		fmt.Printf("2. Прочитать заметку через файл\n")
		fmt.Print("--------------------------------\n")
		fmt.Printf("3. Поиск заметки по тегу\n")
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
			tasks := file.ReadFile()
			for _, task := range tasks {
				fmt.Printf("%v. %s, теги: %v, уровень приоритета - %v\n", task.ID, task.Name, task.Meta.Tags, task.Meta.Priority)
			}
		case 3:
			tasks := cli.Search()
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
