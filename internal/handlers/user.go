package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "go-rest-api/internal/models"
    "go-rest-api/internal/utils"
    "time"
)

// Untuk menyimpan data sementara (ganti dengan database di production)
var users = make(map[string]models.User)

type CreateUserRequest struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}

// GetUsers handles GET /api/v1/users
func GetUsers(w http.ResponseWriter, r *http.Request) {
    userList := make([]models.User, 0)
    for _, user := range users {
        // Remove password before sending
        user.Password = ""
        userList = append(userList, user)
    }
    utils.RespondWithJSON(w, http.StatusOK, userList)
}

// GetUser handles GET /api/v1/users/{id}
func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    user, exists := users[id]
    if !exists {
        utils.RespondWithError(w, http.StatusNotFound, "User not found")
        return
    }

    // Remove password before sending
    user.Password = ""
    utils.RespondWithJSON(w, http.StatusOK, user)
}

// CreateUser handles POST /api/v1/users
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Check if email already exists
    for _, user := range users {
        if user.Email == req.Email {
            utils.RespondWithError(w, http.StatusConflict, "Email already exists")
            return
        }
    }

    // Create new user
    user := models.User{
        ID:        utils.GenerateUUID(), // Implement this in utils
        Name:      req.Name,
        Email:     req.Email,
        Password:  req.Password,
        CreatedAt: time.Now(),
    }

    // Hash password
    if err := user.HashPassword(); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
        return
    }

    // Save user
    users[user.ID] = user

    // Remove password before sending response
    user.Password = ""
    utils.RespondWithJSON(w, http.StatusCreated, user)
}

// UpdateUser handles PUT /api/v1/users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    user, exists := users[id]
    if !exists {
        utils.RespondWithError(w, http.StatusNotFound, "User not found")
        return
    }

    var req UpdateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Update fields
    user.Name = req.Name
    user.Email = req.Email

    // Save updated user
    users[id] = user

    // Remove password before sending response
    user.Password = ""
    utils.RespondWithJSON(w, http.StatusOK, user)
}

// DeleteUser handles DELETE /api/v1/users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    if _, exists := users[id]; !exists {
        utils.RespondWithError(w, http.StatusNotFound, "User not found")
        return
    }

    delete(users, id)
    utils.RespondWithJSON(w, http.StatusNoContent, nil)
}
