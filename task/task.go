package task

type Priority int

type Meta struct {
	Tags []string
	Priority
}

// Сама задача
type Task struct {
	ID   int
	Name string
	Done bool
	Meta Meta
}

// метод проверки корректности
func (p Priority) IsCorrect() bool {
	return p > 0 && p < 6
}
