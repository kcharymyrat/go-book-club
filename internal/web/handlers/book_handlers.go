package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kcharymyrat/go-book-club/pkg/model"
)

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	pageData.Content = Books
	fmt.Printf("%+v\n", pageData)

	w.Header().Set("Content-Type", "text/html")

	booksTmplURL := tmplDirURL + "/books/books.html"
	t, err := template.ParseFiles(
		baseTmplURL, navbarTmplURL, footerTmplURL, booksTmplURL,
	)
	if err != nil {
		log.Fatalf("Could not parse template files. %s", err.Error())
	}

	err = t.ExecuteTemplate(w, "base", pageData)
	if err != nil {
		log.Printf("Could not execute template %v\n%s", t.Tree.Name, err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func AddBookGetHandler(w http.ResponseWriter, r *http.Request) {
	bookFormTmplURL := tmplDirURL + "/books/book-form.html"
	t, err := template.ParseFiles(
		baseTmplURL, navbarTmplURL, footerTmplURL, bookFormTmplURL,
	)
	if err != nil {
		log.Fatal("Could not parse template files. " + err.Error())
	}

	err = t.ExecuteTemplate(w, "base", pageData)
	if err != nil {
		log.Fatalf("Could not execute template %v", t.Tree.Name)
	}
}

func AddBookPostHandler(w http.ResponseWriter, r *http.Request) {

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
		ID:      3,
		Title:   r.PostForm["title"][0],
		Summary: r.PostForm["summary"][0],
		Year:    year,
	}

	// Set headers
	w.Header().Set("Location", book.Title)
	w.Header().Set("Content-Type", "text/html")
	// w.WriteHeader(http.StatusPermanentRedirect) - it auto redirects to given "Location" header

	// Dispay the result
	successTmplURL := tmplDirURL + "/books/book-success.html"
	t, err := template.ParseFiles(
		baseTmplURL, navbarTmplURL, footerTmplURL, successTmplURL,
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
	w.Header().Set("Content-Type", "text/html")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID format", http.StatusBadRequest)
		return
	}

	var book *model.Book
	for _, b := range Books {
		if b.ID == id {
			book = &b
			break
		}
	}

	if book == nil {
		log.Printf("Book with ID %v not found", id)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	pageData.Content = *book
	fmt.Printf("%+v\n", pageData)

	bookDetailTmplURL := tmplDirURL + "/books/book-detail.html"
	t, err := template.ParseFiles(
		baseTmplURL, navbarTmplURL, footerTmplURL, bookDetailTmplURL,
	)
	if err != nil {
		log.Printf("Could not parse template %v\n", t.Tree.Name)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "base", pageData)
	if err != nil {
		log.Printf("Could not execute template %v\n%s", t.Tree.Name, err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func UpdateBookGetHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Update the book, id = %s", id)
}

func UpdateBookPostHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Patch the book, id = %s", id)
}

func DeleteBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Delete the book, id = %s", id)
}
