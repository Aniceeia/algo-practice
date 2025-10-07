package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type contextKey string

const key contextKey = "userID"

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := "w83-hjd"
		ctx := context.WithValue(r.Context(), key, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value(key).(string)
	if !ok {
		http.Error(w, "cant identify you", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "hello, %v!", id)
}

func main() {
	http.HandleFunc("/", authMiddleware(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
