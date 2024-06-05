package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kcharymyrat/go-book-club/internal/model"
)

const tmplUrl = "./internal/web/templates/"

const (
	baseTmpl   = tmplUrl + "base.html"
	navTmpl    = tmplUrl + "nav.html"
	footerTmpl = tmplUrl + "footer.html"
)

var Books = []model.Book{
	{ID: 1, Title: "Book 1", Description: "Desc 1", Year: 2001},
	{ID: 1, Title: "Book 2", Description: "Desc 2", Year: 2001},
	{ID: 1, Title: "Book 3", Description: "Desc 3", Year: 2001},
}

func BookFormHandler(w http.ResponseWriter, r *http.Request) {
	bookFormTmpl := tmplUrl + "books/book-form.html"
	t, err := template.ParseFiles(
		baseTmpl, navTmpl, footerTmpl, bookFormTmpl,
	)
	if err != nil {
		log.Fatal("Could not parse template files. " + err.Error())
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalf("Could not execute template %v", t.Tree.Name)
	}
}

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	booksTmpl := tmplUrl + "books/books.html"
	t, err := template.ParseFiles(
		baseTmpl, navTmpl, footerTmpl, booksTmpl,
	)
	if err != nil {
		log.Fatal("Could not parse template files. " + err.Error())
	}

	err = t.Execute(w, Books)
	if err != nil {
		log.Fatalf("Could not execute template %v", t.Tree.Name)
	}
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {

	// Parse the form
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("Could not parse the form. %v", err.Error())
	}

	// Retrieve parsed values - validation step
	year, err := strconv.Atoi(r.PostForm["year"][0])
	if err != nil {
		log.Fatalf("Could not get the valid year. %v", err.Error())
	}

	book := model.Book{
		ID:          3,
		Title:       r.PostForm["title"][0],
		Description: r.PostForm["description"][0],
		Year:        year,
	}

	// Set headers
	w.Header().Set("Location", book.Title)
	w.Header().Set("Content-Type", "text/html")
	// w.WriteHeader(http.StatusPermanentRedirect) - it auto redirects to given "Location" header

	// Dispay the result
	successTmpl := tmplUrl + "books/book-success.html"
	t, err := template.ParseFiles(
		baseTmpl, navTmpl, footerTmpl, successTmpl,
	)
	if err != nil {
		log.Fatal("Could not parse template files. " + err.Error())
	}

	err = t.Execute(w, book)
	if err != nil {
		log.Fatalf("Could not execute template %v", t.Tree.Name)
	}

}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Get a book with %s", id)
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Update the book, id = %s", id)
}

func PatchBookHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Patch the book, id = %s", id)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Delete the book, id = %s", id)
}
