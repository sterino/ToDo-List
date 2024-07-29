package interfaces

import (
	"context"
	"todo-list/internal/domain/task"
)

type TaskRepository interface {
	Create(ctx context.Context, entity task.Entity) (id string, err error)
	List(ctx context.Context, status string) (res []task.Entity, err error)
	Get(ctx context.Context, id string) (res task.Entity, err error)
	Delete(ctx context.Context, id string) (err error)
	Update(ctx context.Context, id string, entity task.Entity) (err error)
	Done(ctx context.Context, id string) (err error)
}
