package main

import (
	"log"
	"net/http"
	"time"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("incoming request")
		defer log.Printf("request handeled in %s", time.Since(start))
		next.ServeHTTP(w, r)
	})
}
