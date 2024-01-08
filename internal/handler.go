package handler

import (
	"encoding/json"
	"go-rest/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := findUserByID(userID)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func findUserByID(userID int) *model.User {
	for _, user := range model.Users {
		if user.ID == userID {
			return &user
		}
	}
	return nil
}
