package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Viperesss/taskmanager/task"
	"github.com/Viperesss/taskmanager/utils"
)

// читаем файл
func ReadFile() error {
	lineCounter := 1
	var newTask task.Task

	fmt.Println("\nЗаметка должна содержать следующий формат записи: Название заметки; теги; приоритет от 1 до 5")
	fmt.Print("Введите название файла: ")

	input, err := utils.StringReader()
	if err != nil {
		return err
	}

	file, err := os.Open(input) // Открываем файл для работы с ним
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file) // Инициализируем сканнер для чтения файла

	for scanner.Scan() { // Цикл работает до тех пор пока не прочитает все строки
		parts := strings.Split(scanner.Text(), ";")
		tags := strings.Split(parts[1], ",")
		priority, err := strconv.Atoi(parts[2])
		if err != nil {
			return fmt.Errorf("Ошибка при получении уровня приоритета в %v строке\n Ошибка: %v", lineCounter, err)
		}

		err = newTask.AddTask(parts[0], tags, priority)
		if err != nil {
			return fmt.Errorf("Ошибка при создании новой заметки, строка: %v\nОшибка: %v", lineCounter, err)
		}
	}
	fmt.Println("Заметки из файла успешно добавлены")

	err = file.Close() // Закрываем файл для экономии ресурсов, err один
	if err != nil {
		return err
	}
	if scanner.Err() != nil { // Если при чтении файла возникла ошибка - выдаем ошибку
		return scanner.Err()
	}
	return nil
}
