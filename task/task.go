package task

import "fmt"

type priority int

type meta struct {
	tags []string
	priority
}

// Сама задача
type Task struct {
	id   int
	name string
	meta
}

var TaskList []Task

func (t Task) AddTask(name string, tags []string, p int) error {
	id := len(TaskList) + 1
	if p < 1 || p > 5 {
		return fmt.Errorf("Ошибка, уровень приоритета может быть в диапозоне от 1 до 5")
	}

	newTask := Task{
		id:   id,
		name: name,
		meta: meta{
			tags:     tags,
			priority: priority(p),
		},
	}
	TaskList = append(TaskList, newTask)
	return nil
}

func Search(tagName string) {
	var returnedTasks []Task
	for _, task := range TaskList {
		for _, tag := range task.tags {
			if tag == tagName {
				returnedTasks = append(returnedTasks, task)
			}
		}
	}
	for _, task := range returnedTasks {
		fmt.Printf("%v. %s, теги: %v, уровень приоритета - %v\n", task.id, task.name, task.meta.tags, task.meta.priority)
	}
}

func ShowAll() {
	if len(TaskList) != 0 {
		for _, task := range TaskList {
			fmt.Printf("%v. %s, теги: %v, уровень приоритета - %v\n", task.id, task.name, task.meta.tags, task.meta.priority)
		}
	} else {
		fmt.Println("Заметки отсутствуют")
	}
}
