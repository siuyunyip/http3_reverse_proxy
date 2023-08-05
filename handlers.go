package main

import "net/http"

func SetupHandler(root string) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(root)))

	mux.HandleFunc("/1/", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("/2/", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("/3/", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("/4/", func(w http.ResponseWriter, r *http.Request) {

	})

	return mux
}