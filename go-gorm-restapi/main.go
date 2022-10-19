package main

import (
	"fmt"
	"net/http"

	"github.com/Cutshadows/go-gorm-restapi/db"
	"github.com/Cutshadows/go-gorm-restapi/models"
	"github.com/Cutshadows/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	fmt.Println("hello world")
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeletetTaskHandler).Methods("DELETE")
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeletetUserHandler).Methods("DELETE")

	http.ListenAndServe(":2030", r)
}
