package main

import "net/http"

// standard way to write go handlers
func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something Went Wrong!!")
}