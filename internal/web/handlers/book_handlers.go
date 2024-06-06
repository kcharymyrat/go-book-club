package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kcharymyrat/go-book-club/internal/web/utils"
	"github.com/kcharymyrat/go-book-club/pkg/model"
)

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Println("path =", r.URL.Path)

	navbar := pageData.Navbar.(utils.NavbarData)
	navbar.Active = r.URL.Path

	pageData.Navbar = navbar
	pageData.Content = Books
	fmt.Printf("%+v\n", pageData)

	booksTmplURL := tmplDirURL + "/books/books.html"
	t, err := template.ParseFiles(
		baseTmplURL, navbarTmplURL, footerTmplURL, booksTmplURL,
	)
	if err != nil {
		log.Printf("Could not parse template files. %s", err.Error())
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

func AddBookGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	pageData.FormData = struct {
		Title   string
		Summary string
		Year    int
		Author  string
		Authors []*model.Author
	}{
		Authors: Authors,
	}

	fmt.Printf("%+v\n", pageData)

	bookFormTmplURL := tmplDirURL + "/books/book-form.html"
	t, err := template.ParseFiles(
		baseTmplURL, navbarTmplURL, footerTmplURL, bookFormTmplURL,
	)
	if err != nil {
		log.Printf("Could not parse template files. %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return

	}

	err = t.ExecuteTemplate(w, "base", pageData)
	if err != nil {
		log.Printf("Could not execute template %v", t.Tree.Name)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func AddBookPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Parse the form
	err := r.ParseForm()
	if err != nil {
		log.Printf("Could not parse the form. %v", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Retrieve parsed values - validation step
	title := r.PostFormValue("title")
	summary := r.PostFormValue("summary")
	yearStr := r.PostFormValue("year")
	author := r.PostFormValue("author")

	formData := struct {
		Title   string
		Summary string
		Year    int
		Author  string
		Authors []*model.Author
	}{
		Title:   title,
		Summary: summary,
		Author:  author,
		Authors: Authors,
	}

	formErrors := make(map[string]string)

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		formErrors["year"] = "Invalid year format"
	}

	if title == "" {
		formErrors["title"] = "Title is required"
	}

	if summary == "" {
		formErrors["summary"] = "Summary is required"
	}

	if len(formErrors) > 0 {
		pageData.FormData = formData
		pageData.FormErrors = formErrors

		bookFormTmplURL := tmplDirURL + "/books/book-form.html"
		t, err := template.ParseFiles(
			baseTmplURL, navbarTmplURL, footerTmplURL, bookFormTmplURL,
		)
		if err != nil {
			log.Printf("Could not parse template files. %s", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = t.ExecuteTemplate(w, "base", pageData)
		if err != nil {
			log.Printf("Could not execute template %v", t.Tree.Name)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		return
	}

	book := model.Book{
		ID:      3,
		Title:   title,
		Summary: summary,
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
			book = b
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
