//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	_ "github.com/lib/pq"
	http "todo-list/internal/api"
	"todo-list/internal/api/handler"
	"todo-list/internal/config"
	"todo-list/internal/db"
	"todo-list/internal/repository"
	"todo-list/internal/service"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		handler.NewTaskHandler,
		repository.NewTaskRepository,
		service.NewTaskService,
		http.NewServer,
	)
	return &http.Server{}, nil
}
