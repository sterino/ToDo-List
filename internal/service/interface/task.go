package interfaces

import (
	"context"
	"todo-list/internal/domain/task"
)

type TaskService interface {
	CreateTask(ctx context.Context, req task.Request) (id string, err error)
	ListTasks(ctx context.Context, status string) (res []task.Response, err error)
	GetTask(ctx context.Context, id string) (res task.Response, err error)
	DeleteTask(ctx context.Context, id string) (err error)
	UpdateTask(ctx context.Context, id string, req task.Request) (err error)
	DoneTask(ctx context.Context, id string) (err error)
}
