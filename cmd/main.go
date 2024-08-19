// cmd/main.go
package main

import (
	"fmt"
	"net/http"

	//"github.com/<ваше имя на GitHub>/todo-list/database"
	//"github.com/<ваше имя на GitHub>/todo-list/handlers"
	//"proekt_doker/database"
	//"proekt_doker/handlers"
	//"Users/LapTOP/go/Projekts/proekt_doker/database"
	//"Users/LapTOP/go/Projekts/proekt_doker/handlers"

	"github.com/andrey-boyar/proekt_doker-rep/database"
	"github.com/andrey-boyar/proekt_doker-rep/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	database.ConnectDB()

	r := chi.NewRouter()

	r.Get("/tasks", handlers.GetTasks)
	r.Post("/tasks", handlers.PostTask)
	r.Get("/tasks/{id}", handlers.GetTaskByID)
	r.Delete("/tasks/{id}", handlers.DeleteTask)

	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Printf("Start server error: %s", err.Error())
		return
	}
}
