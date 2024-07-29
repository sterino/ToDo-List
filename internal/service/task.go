package service

import (
	"context"
	"todo-list/internal/domain/task"
	interfaces "todo-list/internal/repository/interface"
	services "todo-list/internal/service/interface"
)

type TaskService struct {
	taskRepository interfaces.TaskRepository
}

func NewTaskService(repository interfaces.TaskRepository) services.TaskService {
	return &TaskService{
		taskRepository: repository,
	}
}

func (ts *TaskService) CreateTask(ctx context.Context, req task.Request) (id string, err error) {
	data := task.Entity{
		Title:    req.Title,
		ActiveAt: task.ParseDate(req.ActiveAt),
	}
	id, err = ts.taskRepository.Create(ctx, data)
	return
}

func (ts *TaskService) ListTasks(ctx context.Context, status string) (res []task.Response, err error) {
	if !task.IsValidStatus(status) {
		err = task.ErrorInvalidStatus
		return
	}
	if status == "" {
		status = "active"
	}
	data, err := ts.taskRepository.List(ctx, status)
	if err != nil {
		return nil, err
	}

	for i := range data {
		_, err = data[i].ParseToDayoffs()
		if err != nil {
			return
		}
	}
	res = task.ParseFromEntities(data)
	return
}

func (ts *TaskService) GetTask(ctx context.Context, id string) (res task.Response, err error) {
	data, err := ts.taskRepository.Get(ctx, id)
	if err != nil {
		return
	}
	_, err = data.ParseToDayoffs()
	if err != nil {
		return
	}
	res = task.ParseFromEntity(data)
	return
}

func (ts *TaskService) DeleteTask(ctx context.Context, id string) (err error) {
	err = ts.taskRepository.Delete(ctx, id)
	return
}

func (ts *TaskService) UpdateTask(ctx context.Context, id string, req task.Request) (err error) {
	data := task.Entity{
		Title:    req.Title,
		ActiveAt: task.ParseDate(req.ActiveAt),
	}
	err = ts.taskRepository.Update(ctx, id, data)
	return
}

func (ts *TaskService) DoneTask(ctx context.Context, id string) (err error) {
	err = ts.taskRepository.Done(ctx, id)
	return
}
