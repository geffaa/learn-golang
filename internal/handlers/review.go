package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "go-rest-api/internal/models"
    "go-rest-api/internal/utils"
)

var reviews = make(map[string]models.Review)

func CreateReview(w http.ResponseWriter, r *http.Request) {
    var review models.Review
    if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Validate rating
    if review.Rating < 1 || review.Rating > 5 {
        utils.RespondWithError(w, http.StatusBadRequest, "Rating must be between 1 and 5")
        return
    }

    // Verify product exists
    if _, exists := products[review.ProductID]; !exists {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
        return
    }

    review.ID = utils.GenerateUUID()
    review.CreatedAt = time.Now()
    review.UpdatedAt = time.Now()

    reviews[review.ID] = review

    utils.RespondWithJSON(w, http.StatusCreated, review)
}

func GetProductReviews(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    productID := params["productId"]

    productReviews := make([]models.Review, 0)
    for _, review := range reviews {
        if review.ProductID == productID {
            productReviews = append(productReviews, review)
        }
    }

    utils.RespondWithJSON(w, http.StatusOK, productReviews)
}