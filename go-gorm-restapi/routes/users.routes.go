package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Cutshadows/go-gorm-restapi/db"
	"github.com/Cutshadows/go-gorm-restapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world users"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world users"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createUser := db.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

func DeletetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world users"))
}
