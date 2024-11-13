package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "go-rest-api/internal/models"
    "go-rest-api/internal/utils"
)

// Untuk menyimpan data sementara (ganti dengan database di production)
var products = make(map[string]models.Product)

type CreateProductRequest struct {
    Name        string  `json:"name" validate:"required"`
    Price       float64 `json:"price" validate:"required,gt=0"`
    Description string  `json:"description" validate:"required"`
    Stock       int     `json:"stock" validate:"required,gte=0"`
}

type UpdateProductRequest struct {
    Name        string  `json:"name" validate:"required"`
    Price       float64 `json:"price" validate:"required,gt=0"`
    Description string  `json:"description" validate:"required"`
    Stock       int     `json:"stock" validate:"required,gte=0"`
}

// GetProducts handles GET /api/v1/products
func GetProducts(w http.ResponseWriter, r *http.Request) {
    productList := make([]models.Product, 0)
    for _, product := range products {
        productList = append(productList, product)
    }
    utils.RespondWithJSON(w, http.StatusOK, productList)
}

// GetProduct handles GET /api/v1/products/{id}
func GetProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    product, exists := products[id]
    if !exists {
        utils.RespondWithError(w, http.StatusNotFound, "Product not found")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, product)
}

// CreateProduct handles POST /api/v1/products
func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var req CreateProductRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Create new product
    product := models.Product{
        ID:          utils.GenerateUUID(), // Implement this in utils
        Name:        req.Name,
        Price:       req.Price,
        Description: req.Description,
        Stock:       req.Stock,
    }

    // Save product
    products[product.ID] = product

    utils.RespondWithJSON(w, http.StatusCreated, product)
}

// UpdateProduct handles PUT /api/v1/products/{id}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    if _, exists := products[id]; !exists {
        utils.RespondWithError(w, http.StatusNotFound, "Product not found")
        return
    }

    var req UpdateProductRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Update product
    product := models.Product{
        ID:          id,
        Name:        req.Name,
        Price:       req.Price,
        Description: req.Description,
        Stock:       req.Stock,
    }

    // Save updated product
    products[id] = product

    utils.RespondWithJSON(w, http.StatusOK, product)
}

// DeleteProduct handles DELETE /api/v1/products/{id}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    if _, exists := products[id]; !exists {
        utils.RespondWithError(w, http.StatusNotFound, "Product not found")
        return
    }

    delete(products, id)
    utils.RespondWithJSON(w, http.StatusNoContent, nil)
}