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
			err := cli.ReadCli()
			if err != nil {
				fmt.Println(err)
			}
			continue

		case 2:
			err := file.ReadFile()
			if err != nil {
				fmt.Println(err)
			}
			continue

		case 3:
			task.ShowAll()
			continue

		case 4: // Поиск заметки по тегу
			err := cli.Search()
			if err != nil {
				fmt.Println(err)
			}
			continue

		case 0:
			return

		default:
			fmt.Println("Ошибка, введите цифру из предложенного списка")
			continue
		}
	}
}
