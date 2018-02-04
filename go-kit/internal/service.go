package internal

import (
	"context"
)

type Service interface {
	GetTasks(ctx context.Context) ([]string, error)
}

type service struct{}

func NewService() service {
	return service{}
}

func (s service) GetTasks(ctx context.Context) ([]string, error) {
	return []string{
		"test 1",
		"test 2",
	}, nil
}
