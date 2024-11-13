package handlers

import (
    "encoding/json"
    "net/http"
    "go-rest-api/internal/models"
    "go-rest-api/internal/utils"
)

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type RegisterRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Login handles POST /api/v1/auth/login
func Login(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Find user by email
    user := findUserByEmail(req.Email)
    if user == nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
        return
    }

    // Check password
    if err := user.CheckPassword(req.Password); err != nil {
        utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
        return
    }

    // Generate token
    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error generating token")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{
        "token": token,
    })
}

// Register handles POST /api/v1/auth/register
func Register(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Check if email already exists
    if findUserByEmail(req.Email) != nil {
        utils.RespondWithError(w, http.StatusConflict, "Email already exists")
        return
    }

    // Create new user
    user := &models.User{
        ID:    utils.GenerateUUID(),
        Name:  req.Name,
        Email: req.Email,
    }

    // Set and hash password
    user.Password = req.Password
    if err := user.HashPassword(); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
        return
    }

    // Save user
    users[user.ID] = *user

    // Remove password from response
    user.Password = ""
    utils.RespondWithJSON(w, http.StatusCreated, user)
}

// Helper function to find user by email
func findUserByEmail(email string) *models.User {
    for _, user := range users {
        if user.Email == email {
            return &user
        }
    }
    return nil
}