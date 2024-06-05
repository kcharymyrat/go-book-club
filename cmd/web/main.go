package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/kcharymyrat/go-book-club/internal/web/routes"
)

func main() {

	// get the router (mux)
	mux := routes.Routes()

	// Initializ server
	srv := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	fmt.Println("Server started")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
