package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "go-rest-api/internal/models"
    "go-rest-api/internal/utils"
)

var categories = make(map[string]models.Category)

func GetCategories(w http.ResponseWriter, r *http.Request) {
    categoryList := make([]models.Category, 0)
    for _, category := range categories {
        categoryList = append(categoryList, category)
    }
    utils.RespondWithJSON(w, http.StatusOK, categoryList)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
    var category models.Category
    if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    category.ID = utils.GenerateUUID()
    categories[category.ID] = category

    utils.RespondWithJSON(w, http.StatusCreated, category)
}