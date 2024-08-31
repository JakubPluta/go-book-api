package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/JakubPluta/go-book-api/config"
)

func main() {
	config := config.New()
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		Handler:      mux,
		ReadTimeout:  config.Server.TimeoutRead,
		WriteTimeout: config.Server.TimeoutWrite,
		IdleTimeout:  config.Server.TimeoutIdle,
	}
	log.Println("Listening on :", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed: ", err)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}
