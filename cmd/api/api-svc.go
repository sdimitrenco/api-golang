package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/1", health)
	mux.HandleFunc("/api/1/protected", health)

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

}
