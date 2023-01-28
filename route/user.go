package route

import (
	"encoding/json"
	"net/http"

	"github.com/LeanBorquez/go-apirest/db"
	"github.com/LeanBorquez/go-apirest/models"
	"github.com/gorilla/mux"
)




func GetUsersHandle(w http.ResponseWriter, r *http.Request){
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandle(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user,params["id"])

	if user.ID == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandle(w http.ResponseWriter, r *http.Request){
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createUser := db.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandle(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}