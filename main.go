package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// disable cache
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	fmt.Fprintf(w, "Hello World!")

}

// mux
var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)

	fmt.Println("Listening on port 8000...")
	http.ListenAndServe("localhost:8000", router)

}
