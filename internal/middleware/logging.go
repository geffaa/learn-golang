package middleware

import (
    "log"
    "net/http"
    "time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()
        next.ServeHTTP(w, r)
        log.Printf(
            "%s %s %s",
            r.Method,
            r.RequestURI,
            time.Since(startTime),
        )
    }
}