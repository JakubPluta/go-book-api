package router

import (
	"github.com/JakubPluta/go-book-api/api/resource/book"
	"github.com/JakubPluta/go-book-api/api/resource/health"
	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", health.Read)

	r.Route("/v1", func(r chi.Router) {
		booksApi := &book.API{}
		r.Get("/books", booksApi.List)
		r.Get("/books/{id}", booksApi.Read)
		r.Post("/books", booksApi.Create)
		r.Put("/books/{id}", booksApi.Update)
		r.Delete("/books/{id}", booksApi.Delete)
	})
	return r
}
