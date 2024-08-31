package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  5 * time.Second,
	}
	log.Println("Listening on :8080")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed: ", err)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}
