package repository

type Project struct {
	ID   int64
	Name string
}

type ProjectRepository interface {
	FindAll() ([]Project, error)
	Store(Project) (Project, error)
	Delete(Project) error
}
