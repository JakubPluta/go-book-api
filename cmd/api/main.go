package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JakubPluta/go-book-api/api/router"
	"github.com/JakubPluta/go-book-api/config"
)

//  @title          Books API
//  @version        1.0
//  @description    This is a Books RESTful API with a CRUD

//  @contact.name   Jakub Pluta
//  @contact.url    https://github.com/JakubPluta

// @host       localhost:8080
// @basePath   /v1
func main() {
	config := config.New()
	r := router.New()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		Handler:      r,
		ReadTimeout:  config.Server.TimeoutRead,
		WriteTimeout: config.Server.TimeoutWrite,
		IdleTimeout:  config.Server.TimeoutIdle,
	}
	log.Println("Listening on :", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed: ", err)
	}

}
