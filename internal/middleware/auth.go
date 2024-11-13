package middleware

import (
    "net/http"
    "strings"
    "go-rest-api/internal/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header is required")
            return
        }

        bearerToken := strings.Split(authHeader, " ")
        if len(bearerToken) != 2 {
            utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token format")
            return
        }

        token := bearerToken[1]
        if !utils.ValidateToken(token) {
            utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
            return
        }

        next.ServeHTTP(w, r)
    }
}