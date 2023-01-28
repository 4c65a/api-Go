package route

import (
	"encoding/json"
	"net/http"

	"github.com/LeanBorquez/go-apirest/db"
	"github.com/LeanBorquez/go-apirest/models"
	"github.com/gorilla/mux"
)

func GetTasksHandle(w http.ResponseWriter, r *http.Request) {
	var Task []models.Task
	db.DB.Find(&Task)
	json.NewEncoder(w).Encode(&Task)
}

func GetTaskHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])

	if task.UserId == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func PostTaskHandle(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createUser := db.DB.Create(&task)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])

	if task.UserId == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)
}
