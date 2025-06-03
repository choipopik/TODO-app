package main

import (
	"log"

	"github.com/choipopik/todo-app"
	"github.com/choipopik/todo-app/pkg/handler"
)

func main() {

	handlers := new(handler.Handler)
	srv := new(todo.Server)

	err := srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
