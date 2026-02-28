package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Viperesss/taskmanager/task"
	"github.com/Viperesss/taskmanager/utils"
)

// читаем файл
func ReadFile(tList *[]task.Task) {
	fmt.Println("\nЗаметка должна содержать следующий формат записи: Название заметки; теги; приоритет от 1 до 5")
	for {
		counter := len(*tList) + 1
		fmt.Print("Введите название файла: ")

		input, err := utils.StringReader()
		if err != nil {
			fmt.Println(err)
			return
		}

		file, err := os.Open(input) // Открываем файл для работы с ним
		if err != nil {
			log.Fatal(err)
			continue
		}
		scanner := bufio.NewScanner(file) // Инициализируем сканнер для чтения файла

		for scanner.Scan() { // Цикл работает до тех пор пока не прочитает все строки
			parts := strings.Split(scanner.Text(), ";")
			tags := strings.Split(parts[1], ",")
			priority, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println(err)
				fmt.Println("Ошибка при получении уровня приоритета")
				continue
			}
			task := task.Task{
				ID:   counter,
				Name: parts[0],
				Meta: task.Meta{
					Tags:     tags,
					Priority: priority,
				},
			}
			*tList = append(*tList, task)
			counter++

		}
		fmt.Println("Заметки из файла успешно добавлены")

		err = file.Close() // Закрываем файл для экономии ресурсов, err один
		if err != nil {
			log.Fatal(err)
		}
		if scanner.Err() != nil { // Если при чтении файла возникла ошибка - выдаем ошибку
			log.Fatal(scanner.Err())
		}

		return
	}

}
