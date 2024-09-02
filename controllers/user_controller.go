package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Siddheshk02/go-blog-platform/models"
	"github.com/Siddheshk02/go-blog-platform/services"
	"github.com/gorilla/mux"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.FetchAllUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	user1, err := c.userService.FetchUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user1)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	var user models.UpdateProfileInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.userService.UpdateProfile(id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
