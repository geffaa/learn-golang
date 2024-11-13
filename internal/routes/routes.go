package routes

import (
    "github.com/gorilla/mux"
    "go-rest-api/internal/handlers"
    "go-rest-api/internal/middleware"
)

func SetupRoutes(r *mux.Router) {
    // API versioning prefix
    api := r.PathPrefix("/api/v1").Subrouter()

    // Public routes
    api.HandleFunc("/auth/login", handlers.Login).Methods("POST")
    api.HandleFunc("/auth/register", handlers.Register).Methods("POST")

    // Protected routes - Users
    api.HandleFunc("/users", middleware.AuthMiddleware(handlers.GetUsers)).Methods("GET")
    api.HandleFunc("/users/{id}", middleware.AuthMiddleware(handlers.GetUser)).Methods("GET")
    api.HandleFunc("/users", middleware.AuthMiddleware(handlers.CreateUser)).Methods("POST")
    api.HandleFunc("/users/{id}", middleware.AuthMiddleware(handlers.UpdateUser)).Methods("PUT")
    api.HandleFunc("/users/{id}", middleware.AuthMiddleware(handlers.DeleteUser)).Methods("DELETE")

    // Protected routes - Products
    api.HandleFunc("/products", middleware.AuthMiddleware(handlers.GetProducts)).Methods("GET")
    api.HandleFunc("/products/{id}", middleware.AuthMiddleware(handlers.GetProduct)).Methods("GET")
    api.HandleFunc("/products", middleware.AuthMiddleware(handlers.CreateProduct)).Methods("POST")
    api.HandleFunc("/products/{id}", middleware.AuthMiddleware(handlers.UpdateProduct)).Methods("PUT")
    api.HandleFunc("/products/{id}", middleware.AuthMiddleware(handlers.DeleteProduct)).Methods("DELETE")

	// Categories
	api.HandleFunc("/categories", middleware.AuthMiddleware(handlers.GetCategories)).Methods("GET")
	api.HandleFunc("/categories", middleware.AuthMiddleware(handlers.CreateCategory)).Methods("POST")

	// Orders
	api.HandleFunc("/orders", middleware.AuthMiddleware(handlers.CreateOrder)).Methods("POST")
	api.HandleFunc("/orders/{id}/status", middleware.AuthMiddleware(handlers.UpdateOrderStatus)).Methods("PUT")

	// Reviews
	api.HandleFunc("/products/{productId}/reviews", handlers.GetProductReviews).Methods("GET")
	api.HandleFunc("/reviews", middleware.AuthMiddleware(handlers.CreateReview)).Methods("POST")
}