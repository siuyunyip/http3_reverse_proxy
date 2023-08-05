package main

import "net/http"

func SetupHandler(root string, workers map[string]Worker) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(root)))

	w1 := workers["127.0.0.1:8001"]
	mux.HandleFunc("/1/", func(w http.ResponseWriter, r *http.Request) {
		w1.BuffReq(r, w)
	})

	w2 := workers["127.0.0.1:8002"]
	mux.HandleFunc("/2/", func(w http.ResponseWriter, r *http.Request) {
		w2.BuffReq(r, w)
	})

	w3 := workers["127.0.0.1:8003"]
	mux.HandleFunc("/3/", func(w http.ResponseWriter, r *http.Request) {
		w3.BuffReq(r, w)
	})

	w4 := workers["127.0.0.1:8004"]
	mux.HandleFunc("/4/", func(w http.ResponseWriter, r *http.Request) {
		w4.BuffReq(r, w)
	})

	return mux
}
