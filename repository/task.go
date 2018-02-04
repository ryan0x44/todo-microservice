package repository

type Task struct {
	ID          int64
	Project     *Project
	Description string
}

type TaskRepository interface {
	FindAll() ([]Task, error)
	FindByProject(Project) ([]Task, error)
	Store(Task) (Task, error)
	Delete(Task) error
}
