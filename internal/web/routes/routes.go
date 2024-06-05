package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/kcharymyrat/go-book-club/internal/web/handlers"
)

func Routes() http.Handler {
	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./internal/web/static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Books Handlers
	r.Get("/books", handlers.GetAllBooksHandler)

	r.Get("/books/add-book", handlers.AddBookGetHandler)
	r.Post("/books/add-book", handlers.AddBookPostHandler)

	r.Get("/books/{id}", handlers.GetBookHandler)

	r.Get("/books/update/{id}", handlers.UpdateBookGetHandler)
	r.Post("/books/update/{id}", handlers.UpdateBookPostHandler)

	r.Post("/books/delete/{id}", handlers.DeleteBookByIDHandler)

	return r
}
