package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, "Something went wrong")
}
