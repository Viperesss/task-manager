package task

type Meta struct {
	Tags     []string
	Priority int // Важность от 1 до 5
}

// Сама задача
type Task struct {
	ID   int
	Name string
	Done bool
	Meta Meta
}
