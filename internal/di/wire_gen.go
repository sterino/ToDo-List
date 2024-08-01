// Code generated by Wire. DO NOT EDIT.
//go:build !wireinject
// +build !wireinject

package di

import (
	_ "github.com/lib/pq"
	"todo-list/internal/api"
	"todo-list/internal/api/handler"
	"todo-list/internal/config"
	"todo-list/internal/db"
	"todo-list/internal/repository"
	"todo-list/internal/service"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	sqlxDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	err = db.Migrate(sqlxDB)
	if err != nil {
		return nil, err
	}
	taskRepository := repository.NewTaskRepository(sqlxDB)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)
	server := http.NewServer(taskHandler)
	return server, nil
}
