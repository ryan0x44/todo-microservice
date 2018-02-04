package internal

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetTasksEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetTasksEndpoint: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return s.GetTasks(ctx)
		},
	}
}
