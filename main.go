package main

import (
    "log"
    "net/http"
    "os"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "go-rest-api/internal/routes"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize router
    r := mux.NewRouter()
    
    // Setup routes
    routes.SetupRoutes(r)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server is running on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}