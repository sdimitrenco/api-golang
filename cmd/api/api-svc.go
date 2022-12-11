package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var authData = auth{
	User:     "stas",
	Password: "1111",
}

type auth struct {
	User     string
	Password string
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1", basicAuth(health))
	mux.HandleFunc("/api/v1/protected", basicAuth(protected))

	server := http.Server{
		Addr:         ":8888",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.Printf("Starting server on: %s", server.Addr)

	error := server.ListenAndServe()

	if error != nil {
		fmt.Errorf("Error %v/n", error.Error())
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json, _ := json.Marshal(map[string]interface{}{"status": "ok", "service": "api"})
	w.Write(json)

}

func protected(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	(w).Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json, _ := json.Marshal(map[string]interface{}{"status": "ok", "service": "api"})
	w.Write(json)

}

func basicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if user != authData.User || password != authData.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		handler.ServeHTTP(w, r)

	}
}
