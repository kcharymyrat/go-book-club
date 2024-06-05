package handlers

import (
	"time"

	"github.com/kcharymyrat/go-book-club/internal/web/utils"
	"github.com/kcharymyrat/go-book-club/pkg/model"
)

const booksBaseURL = "/books"

const tmplDirURL = utils.TmplDirURL

const (
	baseTmplURL   = utils.BaseTmplURL
	navbarTmplURL = utils.NavbarTmplURL
	footerTmplURL = utils.FooterTmplURL
)

var homeLink = utils.Link{
	Href:  "/",
	Label: "Home",
}

var booksLink = utils.Link{
	Href:  booksBaseURL,
	Label: "Book",
}

var bookAddFormLink = utils.Link{
	Href:  booksBaseURL + "/add-book",
	Label: "Add Book",
}

var pageData = utils.PageData{
	Title: "Books",
	Navbar: utils.NavbarData{
		Links:  []utils.Link{homeLink, booksLink, bookAddFormLink},
		Active: booksLink.Href,
	},
	Footer:  utils.FooterData{CopyrightText: "Book Club"},
	BaseURL: booksBaseURL,
}

var Books []*model.Book
var Authors []*model.Author

func init() {

	var book1 = model.Book{
		ID:      1,
		Title:   "Book 1",
		Summary: "Summary 1",
		Year:    2001,
		Authors: make([]*model.Author, 0),
	}

	var book2 = model.Book{
		ID:      2,
		Title:   "Book 2",
		Summary: "Summary 2",
		Year:    2002,
		Authors: make([]*model.Author, 0),
	}

	var book3 = model.Book{
		ID:      3,
		Title:   "Book 3",
		Summary: "Summary 3",
		Year:    2003,
		Authors: make([]*model.Author, 0),
	}

	var author1 = model.Author{
		ID:         1,
		FirstName:  "Author 1",
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
		Books:      make([]*model.Book, 0),
	}

	var author2 = model.Author{
		ID:         2,
		FirstName:  "Author 2",
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
		Books:      make([]*model.Book, 0),
	}

	book1.Authors = append(book1.Authors, &author1)
	book2.Authors = append(book2.Authors, &author2)
	book3.Authors = append(book3.Authors, &author1, &author2)

	author1.Books = append(author1.Books, &book1, &book3)
	author2.Books = append(author2.Books, &book2, &book3)

	Books = []*model.Book{&book1, &book2, &book3}
	Authors = []*model.Author{&author1, &author2}
}
