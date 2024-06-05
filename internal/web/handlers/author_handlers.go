package handlers

import "net/http"

func GetAllAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All Authors"))
}
