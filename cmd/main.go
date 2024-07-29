package main

import (
	"log"
	"os"
	"todo-list/internal/config"
	"todo-list/internal/di"

	_ "todo-list/docs"
)

// @title ToDoList API
// @version 1.0
// @description API Server for ToDoList Application
// @BasePath /api
func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal(diErr)
	} else {
		server.Run(infoLog, errorLog)
	}
}
