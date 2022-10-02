package logger

import (
	"log"
	"net/http"
)

func LogHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%-10s %-20s %s", r.Method, r.RemoteAddr, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%-10s %-20s %s", r.Method, r.RemoteAddr, r.URL.Path)
		next(w, r)
	}
}
