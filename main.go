package main

import (
	"net/http"

	"github.com/LeanBorquez/go-apirest/db"
	"github.com/LeanBorquez/go-apirest/models"
	"github.com/LeanBorquez/go-apirest/route"
	"github.com/gorilla/mux"
)

func main() {

	db.Connection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", route.HomeHandler)

	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/tasks", route.GetTasksHandle).Methods("GET")
	s.HandleFunc("/tasks/{id}", route.GetTaskHandle).Methods("GET")
	s.HandleFunc("/tasks", route.PostTaskHandle).Methods("POST")

	// users routes
	s.HandleFunc("/users", route.GetUsersHandle).Methods("GET")
	s.HandleFunc("/users/{id}", route.GetUserHandle).Methods("GET")
	s.HandleFunc("/users", route.PostUserHandle).Methods("POST")
	s.HandleFunc("/users/{id}", route.DeleteUserHandle).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
