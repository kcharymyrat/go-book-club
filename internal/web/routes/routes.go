package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	book_handlers "github.com/kcharymyrat/go-book-club/internal/web/handlers"
)

func Routes() http.Handler {
	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Books Handlers
	r.Get("/books/create-form", book_handlers.BookFormHandler)
	r.Get("/books", book_handlers.GetAllBooksHandler)
	r.Post("/books", book_handlers.CreateBookHandler)
	r.Get("/books/{id}", book_handlers.GetBookHandler)
	r.Put("/books/{id}", book_handlers.UpdateBookHandler)
	r.Patch("/books/{id}", book_handlers.PatchBookHandler)
	r.Delete("/books/{id}", book_handlers.DeleteBookHandler)

	return r
}
